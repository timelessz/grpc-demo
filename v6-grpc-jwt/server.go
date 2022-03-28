package main

import (
	"context"
	"fmt"
	"go-kit-demo/v6-grpc-jwt/auth"
	"go-kit-demo/v6-grpc-jwt/tool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	port = ":8888"
)

type server struct{}

// 登录 并获取 token 首先登录
func (*server) Login(ctx context.Context, lr *auth.LoginRequest) (*auth.LoginResponse, error) {
	userName, password := lr.Username, lr.Password
	if !(userName == "timeless" && password == "timeless") {
		return nil, status.Errorf(codes.PermissionDenied, "账号或密码错误")
	}
	//生成token
	token, err := tool.CreateJwtToken(1, "timeless")
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "token generate failed!")
	}
	fmt.Println(token)
	return &auth.LoginResponse{Token: token}, status.Errorf(codes.OK, "")
}

// 根据id 从用户信息中获取
func (s *server) GetUserInfo(ctx context.Context, UIR *auth.UserIdRequest) (*auth.User, error) {
	//parseJwtToken()
	token, err := getTokenFromContext(ctx)
	println(token)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token 验证异常")
	}
	mapClaim, err := tool.ParseJwtToken(token)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token 解析异常")
	}
	return &auth.User{
		Id:           mapClaim.Id,
		Email:        "demo@qq.com",
		Username:     "timeless",
		PasswordHash: "dsadadsadas",
		IsAdmin:      false,
	}, status.Errorf(codes.OK, "")
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, status := metadata.FromIncomingContext(ctx)
	if !status {
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}
	// md 的类型是 type MD map[string][]string
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}
	fmt.Println(token)
	// 因此，token 是一个字符串数组，我们只用了 token[0]
	return token[0], nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	auth.RegisterAuthServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
