package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-kit-demo/v8-gokit-all/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService interface {
	Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error)
	GetUserInfo(context.Context, *auth.UserIdRequest) (*auth.User, error)
	AuthHealthCheck(context.Context, *emptypb.Empty) (*auth.HealthResponse, error)
}

type NewLoginMiddlewareService func(server AuthService) AuthService

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type AuthServer struct {
	logger *logrus.Logger
}

func (a AuthServer) AuthHealthCheck(ctx context.Context, empty *emptypb.Empty) (*auth.HealthResponse, error) {
	return &auth.HealthResponse{Status: "OK"}, nil
}

func (a AuthServer) Login(ctx context.Context, lr *auth.LoginRequest) (*auth.LoginResponse, error) {
	println("AuthServer.Login")
	return &auth.LoginResponse{Token: "demo"}, status.Errorf(codes.OK, "")
}

func (a AuthServer) GetUserInfo(ctx context.Context, uir *auth.UserIdRequest) (*auth.User, error) {
	return &auth.User{
		Id:           1,
		Email:        "834916321@qq.com",
		Username:     "zhaoxingzhuang",
		PasswordHash: "demodemodmeo",
		IsAdmin:      false,
	}, status.Errorf(codes.OK, "ok")
}

//
func NewLoginService(logger *logrus.Logger) AuthService {
	println("NewLoginService")
	authServer := AuthServer{logger: logger}
	return NewAuthServiceLogMiddleware(logger)(authServer)
}
