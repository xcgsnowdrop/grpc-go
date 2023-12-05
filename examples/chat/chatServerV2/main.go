// Package main implements a server for Chat service.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/chat/chat_v2"
)

type Room struct {
	id         string
	msgHistory []*pb.ChatMessage
	msgQueue   chan *pb.ChatMessage
	streamMap  map[string]pb.Chat_SendMessageServer
	mu         sync.Mutex
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
		id:        "world",
		msgQueue:  make(chan *pb.ChatMessage, 100),
		streamMap: make(map[string]pb.Chat_SendMessageServer),
	}
	room2 := &Room{
		id:        "guild",
		msgQueue:  make(chan *pb.ChatMessage, 100),
		streamMap: make(map[string]pb.Chat_SendMessageServer),
	}

	if _, ok := s.rooms[room1.id]; !ok {
		s.rooms[room1.id] = room1
	}
	if _, ok := s.rooms[room2.id]; !ok {
		s.rooms[room2.id] = room2
	}
}

// 每个客户端启动都会调用SendMessage以获得客户端stream对象，这里也会获得服务器stream对象，该stream对象相当于TCP中的conn
func (s *server) SendMessage(stream pb.Chat_SendMessageServer) error {

	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	log.Printf("%v", stream)

	// receive messages - init a go routine
	go s.receiveFromStream(stream, clientUniqueCode, errch)

	// send messages - init a go routine
	go s.sendToStream(stream, clientUniqueCode, errch)

	return <-errch
}

func (s *server) broadcastRoomMessage(senderStream pb.Chat_SendMessageServer, room *Room) {
	for {
		select {
		case msg := <-room.msgQueue:
			log.Printf("broadcast for room: %s senderStream: %v message[%v]", room.id, senderStream, msg)

			for _, stream := range room.streamMap {
				if stream != senderStream {
					log.Printf("send msg to loop stream: %v", stream)
					if err := stream.Send(msg); err != nil {
						log.Printf("Error broadcast message in room %s: %v", room.id, err)
					}
				}
			}
		}
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

// receive messages
// 聊天服务器侧，从客户端接收消息，即从grpc.ServerStream中接收来自客户端的消息
func (s *server) receiveFromStream(stream pb.Chat_SendMessageServer, clientUniqueCode_ int, errch_ chan error) {
	// 循环接收来自客户端的消息
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errch_ <- err
		} else {
			username := msg.GetUsername()
			roomid := msg.Roomid

			// Add client to the map
			s.mu.Lock()

			if _, ok := s.rooms[roomid]; ok {
				s.rooms[roomid].mu.Lock()
				s.rooms[roomid].msgQueue <- msg
				s.rooms[roomid].msgHistory = append(s.rooms[roomid].msgHistory, msg)
				s.rooms[roomid].streamMap[username] = stream
				s.rooms[roomid].mu.Unlock()
			} else {
				log.Printf("不会进入该分支,因为已经提前初始化过rooms")
				s.rooms[roomid] = &Room{
					id:        roomid,
					msgQueue:  make(chan *pb.ChatMessage, 100),
					streamMap: make(map[string]pb.Chat_SendMessageServer),
				}
				s.rooms[roomid].mu.Lock()
				s.rooms[roomid].msgQueue <- msg
				s.rooms[roomid].msgHistory = append(s.rooms[roomid].msgHistory, msg)
				s.rooms[roomid].streamMap[username] = stream
				s.rooms[roomid].mu.Unlock()
				// s.broadcastUserJoinMessage(username, s.rooms[roomid])
			}
			s.mu.Unlock()

			// 打印消息队列中的最新一条消息
			log.Printf("msg from stream %v is %v", stream, msg)
		}
	}
}

// send message
// 聊天服务器侧，向客户端发送消息，即向grpc.ServerStream发送消息
func (s *server) sendToStream(stream pb.Chat_SendMessageServer, clientUniqueCode_ int, errch_ chan error) {
	for _, room := range s.rooms {
		go s.broadcastRoomMessage(stream, room)
	}
}

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
