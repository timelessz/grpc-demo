package main

import (
	"context"
	"fmt"
	hellostream "go-kit-demo/v4-stream-grpc/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
	}
	client := hellostream.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hellostream.String{Value: "baidu.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
	stream, errs := client.Channel(context.Background())
	if errs != nil {
		log.Fatal(errs)
	}
	go func() {
		for {
			if err := stream.Send(&hellostream.String{Value: "赵兴壮"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
