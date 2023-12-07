package main

import (
	"flag"
	"fmt"
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
		authServicePath + "CreateUser":  true,
		authServicePath + "UploadImage": true,
		authServicePath + "RateLaptop":  true,
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

	// 注意：这里需要用conn2重新创建authClient2，调用CreateUser()才会触发客户端拦截器，因为客户端拦截器设置在conn2上
	authClient2 := client.NewAuthClient(conn2, username, password)
	success, err := authClient2.CreateUser("xcg", "123456", "user")
	fmt.Printf("create user success: %v, err: %v", success, err)

	// 等待通道的值，阻塞主进程
	done := make(chan bool)
	<-done
}
