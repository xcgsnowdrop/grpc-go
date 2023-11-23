// Binary client is an example client.
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/chat/chat"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

// clienthandle封装pb.Chat_SayServer和客户端名
type clienthandle struct {
	stream     pb.Chat_SayClient
	clientName string
}

func (ch *clienthandle) clientConfig() {
	fmt.Printf("Your Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	ch.clientName = strings.Trim(name, "\r\n")

}

// send message
func (ch *clienthandle) sendMessage() {
	// 循环从Stdin读取消息并发送到聊天服务器
	for {
		// fmt.Printf("Your Message: ")
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &pb.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}

// receive message
func (ch *clienthandle) receiveMessage() {

	//循环从聊天服务器接收消息
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		//print message to console
		fmt.Printf("Saying by %s : %s \n", mssg.Name, mssg.Body)
	}
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatClient(conn)

	// 调用客户端的ChatService获取到一个封装了grpc.ClientStream的stream对象
	stream, err := c.Say(context.Background())
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}

	// implement communication with gRPC server
	ch := clienthandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	//blocker
	bl := make(chan bool)
	<-bl

}
