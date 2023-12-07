package service

import (
	"context"

	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/auth/proto"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userStore  UserStore
	jwtManager *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	authServer := new(AuthServer)
	authServer.userStore = userStore
	authServer.jwtManager = jwtManager
	return authServer
	// return &AuthServer{userStore, jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}

func (server *AuthServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	if user != nil {
		return &pb.CreateUserResponse{Success: false, ErrorMessage: "user already exist"}, nil
	}

	err = server.userStore.CreateUser(req.Username, req.Password, req.Role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}

	res := &pb.CreateUserResponse{Success: true}
	return res, nil
}
