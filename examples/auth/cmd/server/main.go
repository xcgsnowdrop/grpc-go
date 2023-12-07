package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/auth/proto"
	"google.golang.org/grpc/examples/auth/service"
	"google.golang.org/grpc/status"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

var (
	port = flag.Int("port", 50051, "the port to serve on")

	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
)

func accessibleRoles() map[string][]string {
	const authServicePath = "/proto.AuthService/"

	return map[string][]string{
		// authServicePath + "Login":        {"admin"},
		authServicePath + "CreateUser":  {"admin"},
		authServicePath + "UploadImage": {"admin"},
		authServicePath + "RateLaptop":  {"admin", "user"},
	}
}

func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Listening on local port %q: %v", *port, err)
	}

	userStore := service.NewInMemoryUserStore()
	err2 := seedUsers(userStore)
	if err2 != nil {
		log.Fatal("cannot seed users: ", err2)
	}

	jwtManager := service.NewJWTManager(secretKey, tokenDuration)

	authServer := service.NewAuthServer(userStore, jwtManager)

	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serving Echo service on local port: %v", err)
	}
}
