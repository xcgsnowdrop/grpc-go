package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/auth/client"
	"google.golang.org/grpc/examples/data"
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

// 使用该方法需要修改Subject Alternative Name (SAN)，将localhost:50051添加到SAN
// 通过$ openssl x509 -in server_cert.pem -noout -text可以看到server的证书server_cert.pem中可以看到Subject Alternative Name: DNS:*.test.example.com
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(data.Path("x509/ca_cert.pem"))
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

// 等价于loadTLSCredentials()
func loadTLSCredentials2() (credentials.TransportCredentials, error) {
	// Create tls based credential.
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	return creds, err
}

func main() {
	flag.Parse()
	log.Printf("dial server %s", *addr)

	// Create tls based credential.
	tlsCredentials, err := loadTLSCredentials2()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(
		*addr,
		grpc.WithTransportCredentials(tlsCredentials),
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
		grpc.WithTransportCredentials(tlsCredentials),
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
