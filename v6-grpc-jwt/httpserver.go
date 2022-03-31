package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-kit-demo/v6-grpc-jwt/auth"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := "localhost:8888"
	err := auth.RegisterAuthHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux)
}
