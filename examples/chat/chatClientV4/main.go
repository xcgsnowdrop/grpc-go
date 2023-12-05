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
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	pb "google.golang.org/grpc/examples/chat/chat_v4"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/metadata"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")
var token = ""

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
func logger(format string, a ...any) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}

type Client struct {
	username string
	roomid   string
	stream   pb.Chat_SendMessageClient
}

func initClient() *Client {
	fmt.Printf("Your Name: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	username = strings.Trim(username, "\r\n")

	fmt.Printf("Join Roomid: ")

	roomid, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	roomid = strings.Trim(roomid, "\r\n")

	return &Client{
		username: username,
		roomid:   roomid,
	}
}

func initClient1() *Client {
	return &Client{
		username: "xcg1",
		roomid:   "world",
	}
}

// 注意，这里的形参不能为JoinRoom(chatClient *pb.ChatClient)否则会报错: chatClient.JoinRoom undefined (type *chat_v3.ChatClient is pointer to interface, not interface)
func (c *Client) JoinRoom(chatClient pb.ChatClient) error {
	// Connect to server.
	req := pb.JoinRoomRequest{
		Username: c.username,
		Roomid:   c.roomid,
	}
	resp, err := chatClient.JoinRoom(context.Background(), &req)
	if err != nil {
		return err
	}

	token = resp.Token
	// Initialize stream with token.
	header := metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", resp.Token), "roomid": resp.Roomid, "uid": resp.Uid})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	stream, err := chatClient.SendMessage(ctx)
	if err != nil {
		return err
	}

	c.stream = stream

	return nil
}

// send message
func (c *Client) sendMessage() {
	// 循环从Stdin读取消息并发送到聊天服务器
	for {
		// fmt.Printf("Your Message: ")
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &pb.ChatMessage{
			Username: c.username,
			Message:  clientMessage,
		}

		err = c.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}

// receive message 循环从聊天服务器接收消息
func (c *Client) receiveMessage() {
	for {
		resp, err := c.stream.Recv()
		if err != nil {
			log.Printf("can not receive, error: %v", err)
			return
		}

		//print message to console
		fmt.Printf("Saying by %s : %s \n", resp.Username, resp.Message)
	}
}

// wrappedStream  wraps around the embedded grpc.ClientStream, and intercepts the RecvMsg and
// SendMsg method call.
type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m any) error {
	logger("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m any) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

// streamInterceptor is an example stream interceptor.
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.TokenSource{
			TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		}))
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

func main() {
	flag.Parse()

	// Create tls based credential.
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStreamInterceptor(streamInterceptor))
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithStreamInterceptor(streamInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	chatClient := pb.NewChatClient(conn)

	client := initClient()
	if err := client.JoinRoom(chatClient); err != nil {
		log.Printf("JoinRoom error: %v", err)
	}

	// Handle stream messages.
	go client.receiveMessage()
	go client.sendMessage()

	//blocker
	bl := make(chan bool)
	<-bl

}
