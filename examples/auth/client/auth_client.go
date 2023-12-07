package client

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/auth/proto"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

func (client *AuthClient) CreateUser(username string, password string, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	req := &pb.CreateUserRequest{
		Username: username,
		Password: password,
		Role:     role,
	}

	res, err := client.service.CreateUser(ctx, req)
	if err != nil {
		return false, err
	}
	if res.ErrorMessage != "" {
		return false, errors.New(res.ErrorMessage)
	}

	return res.Success, nil
}
