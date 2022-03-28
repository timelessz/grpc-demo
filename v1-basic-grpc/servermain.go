package main

import (
	"context"
	"github.com/gofrs/uuid"
	"go-kit-demo/basic-grpc/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type server struct {
	productMap map[string]*product.Product
}

func main() {
	port := "0.0.0.0:50011"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}
	s := grpc.NewServer()
	product.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

func (s *server) AddProduct(ctx context.Context, in *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "生成uuid 失败", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}
	s.productMap[in.Id] = in
	resp.Value = in.Id
	return resp, err
}

func (s *server) GetProduct(ctx context.Context, in *product.ProductId) (resp *product.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}
	return s.productMap[in.Value], nil
}
