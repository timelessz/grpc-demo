package main

import (
	"context"
	"fmt"
	"go-kit-demo/v6-grpc-jwt/auth"
	"go-kit-demo/v6-grpc-jwt/tool"
	"google.golang.org/grpc"
	"time"
)

const (
	addr = "localhost:8888"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
	defer conn.Close()
	a := auth.NewAuthClient(conn)
	resp, err := a.Login(ctx, &auth.LoginRequest{
		Username: "timeless",
		Password: "timeless",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Token)
	aToken := new(tool.AuthToken)
	aToken.Token = resp.Token
	conn, _ = grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithPerRPCCredentials(aToken))
	defer conn.Close()
	a = auth.NewAuthClient(conn)
	userLoginInfo, err := a.GetUserInfo(ctx, &auth.UserIdRequest{Id: 0})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userLoginInfo.Username)
}
