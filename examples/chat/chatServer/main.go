// Package main implements a server for Chat service.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/chat/chat"
)

type messageUnit struct {
	ClientName        string // 客户端名
	MessageBody       string // 消息体
	MessageUniqueCode int    // 消息唯一标识码
	ClientUniqueCode  int    // 客户端唯一标识码
}

type messageHandle struct {
	MQue []messageUnit // 消息队列
	mu   sync.Mutex    // 互斥锁
}

var (
	port                = flag.Int("port", 50051, "The server port")
	messageHandleObject = messageHandle{}
)

// server is used to implement chat.ChatServer.
type server struct {
	pb.UnimplementedChatServer
}

func (s *server) Say(stream pb.Chat_SayServer) error {

	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	// receive messages - init a go routine
	go receiveFromStream(stream, clientUniqueCode, errch)

	// send messages - init a go routine
	go sendToStream(stream, clientUniqueCode, errch)

	return <-errch
}

// receive messages
// 聊天服务器侧，从客户端接收消息，即从grpc.ServerStream中接收来自客户端的消息
func receiveFromStream(stream pb.Chat_SayServer, clientUniqueCode_ int, errch_ chan error) {

	// 循环接收来自客户端的消息，保存到消息队列MQue中
	for {
		mssg, err := stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errch_ <- err
		} else {

			messageHandleObject.mu.Lock()

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageUnit{
				ClientName:        mssg.Name,
				MessageBody:       mssg.Body,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode:  clientUniqueCode_,
			})

			// 打印消息队列中的最新一条消息
			log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])

			messageHandleObject.mu.Unlock()

		}
	}
}

// send message
// 聊天服务器侧，向客户端发送消息，即向grpc.ServerStream发送消息
func sendToStream(stream pb.Chat_SayServer, clientUniqueCode_ int, errch_ chan error) {

	// 循环
	for {

		//loop through messages in MQue
		// 安装先进先出顺序，循环发送消息队列MQue中接收到的所有来自客户端的消息
		for {

			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			senderUniqueCode := messageHandleObject.MQue[0].ClientUniqueCode // 客户端唯一标识码
			senderName4Client := messageHandleObject.MQue[0].ClientName      // 客户端名
			message4Client := messageHandleObject.MQue[0].MessageBody        // 消息体

			messageHandleObject.mu.Unlock()

			//send message to designated client (do not send to the same client)
			if senderUniqueCode != clientUniqueCode_ {

				err := stream.Send(&pb.FromServer{Name: senderName4Client, Body: message4Client})

				if err != nil {
					errch_ <- err
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:] // delete the message at index 0 after sending to receiver
				} else {
					messageHandleObject.MQue = []messageUnit{}
				}

				messageHandleObject.mu.Unlock()

			}

		}

		time.Sleep(100 * time.Millisecond)
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
	pb.RegisterChatServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
