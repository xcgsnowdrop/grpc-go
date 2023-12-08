package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/auth/proto"
	"google.golang.org/grpc/examples/auth/service"
	"google.golang.org/grpc/examples/data"
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

// 生成初始用户
func seedUsers(userStore service.UserStore) error {
	err := userStore.CreateUser("admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return userStore.CreateUser("user1", "secret", "user")
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

// 等价于loadTLSCredentials()
func loadTLSCredentials2() (credentials.TransportCredentials, error) {
	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}
	return creds, err
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

	// Create tls based credential.
	tlsCredentials, err := loadTLSCredentials2()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serving Echo service on local port: %v", err)
	}
}
