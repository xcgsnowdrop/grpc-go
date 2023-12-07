package main

import (
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/examples/auth/client"
)

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const authServicePath = "/proto.AuthService/"

	return map[string]bool{
		// authServicePath + "Login":        true,
		authServicePath + "CreateLaptop": true,
		authServicePath + "UploadImage":  true,
		authServicePath + "RateLaptop":   true,
	}
}

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()
	log.Printf("dial server %s", *addr)

	conn, err := grpc.Dial(
		*addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("grpc.Dial(%q): %v", *addr, err)
	}
	defer conn.Close()

	authClient := client.NewAuthClient(conn, username, password)

	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	conn2, err := grpc.Dial(
		*addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial(%q): %v", *addr, err)
	}
	defer conn2.Close()

	// 等待通道的值，阻塞主进程
	done := make(chan bool)
	<-done
}
