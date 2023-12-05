// Package main implements a server for Chat service.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/chat/chat_v4"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Token uuid.UUID

type Client struct {
	uid          uuid.UUID
	username     string
	token        string
	streamServer pb.Chat_SendMessageServer
	done         chan error
	roomid       string
}

type Room struct {
	id         string
	msgHistory []*pb.ChatMessage
	msgQueue   chan *pb.ChatMessage
	clients    map[uuid.UUID]*Client
	mu         sync.RWMutex
}

var (
	port               = flag.Int("port", 50051, "The server port")
	secretKey          = []byte("your-secret-key") // 请替换为实际的密钥
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
func logger(format string, a ...any) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}

// server is used to implement chat.ChatServer.
type server struct {
	pb.UnimplementedChatServer
	mu    sync.Mutex
	rooms map[string]*Room
}

func (s *server) initRoom() {
	s.rooms = make(map[string]*Room)
	var room1 *Room = &Room{
		id:       "world",
		msgQueue: make(chan *pb.ChatMessage, 100),
		clients:  make(map[uuid.UUID]*Client),
	}
	room2 := &Room{
		id:       "guild",
		msgQueue: make(chan *pb.ChatMessage, 100),
		clients:  make(map[uuid.UUID]*Client),
	}

	if _, ok := s.rooms[room1.id]; !ok {
		s.rooms[room1.id] = room1
	}
	if _, ok := s.rooms[room2.id]; !ok {
		s.rooms[room2.id] = room2
	}
}

func (s *server) getRoomByid(roomid string) (*Room, error) {
	if room, ok := s.rooms[roomid]; ok {
		return room, nil
	}
	return nil, errors.New(fmt.Sprintf("room %s not found.", roomid))
}

func (s *server) JoinRoom(ctx context.Context, req *pb.JoinRoomRequest) (*pb.JoinRoomReply, error) {
	// TODO: authorize user

	uid := uuid.New()

	// 生成token

	// 设置过期时间为7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":      expirationTime.Unix(),
			"iat":      time.Now().Unix(),
			"iss":      "my-auth-server",
			"sub":      "john",
			"roomid":   req.Roomid,
			"username": req.Username,
			"uid":      uid,
		})
	// token值类似："eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1hdXRoLXNlcnZlciIsInJvb21pZCI6IndvcmxkIiwic3ViIjoiam9obiIsInVzZXJuYW1lIjoieGNnMSJ9.soXMNtYD4kaUk4s8l-Gr5suczqieP9Z7UCiwlNZaXoo"
	// token经过base64 decode后类似：{"alg":"HS256","typ":"JWT"}{"iss":"my-auth-server","roomid":"world","sub":"john","username":"xcg1"}6F<˜Ψ?{P(Z^
	token, err := t.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	client := &Client{
		uid:      uid,
		username: req.Username,
		token:    token,
		roomid:   req.GetRoomid(),
	}
	room, err := s.getRoomByid(req.Roomid)
	if err != nil {
		return nil, err
	}
	room.clients[uid] = client
	s.rooms[req.Roomid] = room
	return &pb.JoinRoomReply{
		Token:  token,
		Roomid: req.Roomid,
		Uid:    uid.String(),
	}, nil
}

func (s *server) getClientFromContext(ctx context.Context) (*Client, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	roomidRaw := md["roomid"]
	uidRaw := md["uid"]

	if len(roomidRaw) == 0 {
		return nil, errors.New("no roomid provided")
	}

	roomid := roomidRaw[0]
	stringUid := uidRaw[0] // 类似于："68eb1228-72ef-4b97-8016-055683e61bd3"

	// 将字符串解析为 UUID 类型
	uid, err := uuid.Parse(stringUid)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return nil, err
	}

	s.mu.Lock()
	currentRoom, ok := s.rooms[roomid]
	if !ok {
		return nil, errors.New(fmt.Sprintf("room %s not found.", roomid))
	}
	currentRoom.mu.RLock()
	currentClient, ok := currentRoom.clients[uid]
	currentRoom.mu.RUnlock()
	s.mu.Unlock()
	if !ok {
		return nil, errors.New("token not recognized")
	}
	return currentClient, nil
}

// 每个客户端启动都会调用SendMessage以获得客户端stream对象，这里也会获得服务器stream对象，该stream对象相当于TCP中的conn
// 客户端调用func (c *chatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error)初始化stream时，可以先不发送消息，仅仅通过context发送meta数据【包括token和roomid】
func (s *server) SendMessage(stream pb.Chat_SendMessageServer) error {
	ctx := stream.Context()
	currentClient, err := s.getClientFromContext(ctx)
	if err != nil {
		return err
	}
	if currentClient.streamServer != nil {
		return errors.New("stream already active")
	}
	currentClient.streamServer = stream

	// Wait for stream requests
	go func() {
		for {
			req, err := stream.Recv()
			if err != nil {
				log.Printf("receive error %v", err)
				currentClient.done <- errors.New("failed to receive request")
				return
			}
			log.Printf("got message %+v", req)
			s.handleChatMessage(req, currentClient)
		}
	}()

	// Wait for stream to be done
	var doneError error
	select {
	case <-ctx.Done():
		doneError = ctx.Err()
	case doneError = <-currentClient.done:
	}
	log.Printf(`stream done with error "%v"`, doneError)

	log.Printf("%s - removing client", currentClient.token)

	s.removeClient(currentClient)

	return doneError
}

func (s *server) handleChatMessage(req *pb.ChatMessage, currentClient *Client) {
	room, ok := s.rooms[currentClient.roomid]
	if !ok {
		currentClient.done <- errors.New(fmt.Sprintf("room %s not found", currentClient.roomid))
		return
	}
	s.broadcast(req, room)
}

func (s *server) removeClient(currentClient *Client) {
	room, ok := s.rooms[currentClient.roomid]
	if !ok {
		currentClient.done <- errors.New(fmt.Sprintf("room %s not found", currentClient.roomid))
		return
	}
	room.mu.Lock()
	delete(room.clients, currentClient.uid)
	room.mu.Unlock()
}

func (s *server) broadcast(msg *pb.ChatMessage, room *Room) {
	room.mu.Lock()
	room.msgHistory = append(room.msgHistory, msg)
	room.mu.Unlock()
	for token, currentClient := range room.clients {
		if currentClient.streamServer == nil {
			continue
		}
		if err := currentClient.streamServer.Send(msg); err != nil {
			log.Printf("%s - broadcast error %v", token, err)
			currentClient.done <- errors.New("failed to broadcast message")
			continue
		}
		log.Printf("%s - broadcasted %+v", msg, token)
	}
}

// func (s *server) broadcastUserJoinMessage(username string, room *Room) {
// 	joinMessage := &pb.ChatMessage{
// 		Username: "Server",
// 		Message:  fmt.Sprintf("%s join the room", username),
// 		Roomid:   room.id,
// 	}
// 	room.mu.Lock()
// 	room.msgQueue <- joinMessage
// 	room.msgHistory = append(room.msgHistory, joinMessage)
// 	room.mu.Unlock()
// 	s.broadcastRoomMessage(nil, room)
// }

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	// 新建一个*grpc.Server对象
	s := grpc.NewServer(grpc.Creds(creds), grpc.StreamInterceptor(authStreamInterceptor))
	// 向s注册服务
	var server *server = &server{}
	server.initRoom()
	pb.RegisterChatServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	tokenString := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		// 验证签名方法是否为HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥
		return secretKey, nil
	})

	// 处理解析过程中可能的错误
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Fatal("Token signature is invalid")
			return false
		}
		log.Fatal("Error parsing token: %v\n", err)
		return false
	}

	// 验证token是否有效
	if !token.Valid {
		log.Fatal("Token is invalid")
		return false
	}

	// 获取token中的声明信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error extracting claims")
		return false
	}

	// 验证有效期
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		log.Fatal("Token has expired")
		return false
	}

	return true
}

// wrappedStream wraps around the embedded grpc.ServerStream, and intercepts the RecvMsg and
// SendMsg method call.
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m any) error {
	logger("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m any) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// authStreamInterceptor looks up the authorization header from the incoming RPC context,
// retrieves the username from it and creates a new context with the username before invoking
// the provided handler.
func authStreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return errInvalidToken
	}

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		logger("RPC failed with error: %v", err)
	}
	return err
}
