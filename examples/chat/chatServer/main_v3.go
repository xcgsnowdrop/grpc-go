// Package main implements a server for Chat service.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/chat/chat_v3"
	"google.golang.org/grpc/metadata"
)

type Token uuid.UUID

type Client struct {
	username     string
	token        uuid.UUID
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
	port = flag.Int("port", 50051, "The server port")
)

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
	token := uuid.New()
	client := &Client{
		username: req.Username,
		token:    token,
		roomid:   req.GetRoomid(),
	}
	room := &Room{
		id:       req.GetRoomid(),
		msgQueue: make(chan *pb.ChatMessage, 100),
	}
	room, err := s.getRoomByid(req.Roomid)
	if err != nil {
		return nil, err
	}
	room.clients[token] = client
	s.rooms[req.Roomid] = room
	return &pb.JoinRoomReply{
		Token:  token.String(),
		Roomid: req.Roomid,
	}, nil
}

func (s *server) getClientFromContext(ctx context.Context) (*Client, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	tokenRaw := headers["authorization"]
	roomidRaw := headers["roomid"]
	if len(tokenRaw) == 0 {
		return nil, errors.New("no token provided")
	}
	if len(roomidRaw) == 0 {
		return nil, errors.New("no roomid provided")
	}
	token, err := uuid.Parse(tokenRaw[0])
	if err != nil {
		return nil, errors.New("cannot parse token")
	}
	roomid := roomidRaw[0]
	s.mu.Lock()
	currentRoom, ok := s.rooms[roomid]
	if !ok {
		return nil, errors.New(fmt.Sprintf("room %s not found.", roomid))
	}
	currentRoom.mu.RLock()
	currentClient, ok := currentRoom.clients[token]
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
	delete(room.clients, currentClient.token)
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
	// 新建一个*grpc.Server对象
	s := grpc.NewServer()
	// 向s注册服务
	var server *server = &server{}
	server.initRoom()
	pb.RegisterChatServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
