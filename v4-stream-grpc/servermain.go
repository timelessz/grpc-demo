package main

import (
	"context"
	"fmt"
	hellostream "go-kit-demo/v4-stream-grpc/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, args *hellostream.String) (*hellostream.String, error) {
	//TODO implement me
	reply := &hellostream.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (h *HelloService) Channel(stream hellostream.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println(args)
		reply := &hellostream.String{Value: "hello:" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	hS := new(HelloService)
	hellostream.RegisterHelloServiceServer(grpcServer, hS)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err.Error())
	}
	grpcServer.Serve(lis)
}
