package main

import (
	"context"
	"fmt"
	pb "go-kit-demo/v3-grpc-grpc-gateway/book"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

const (
	addr = "localhost:8088"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewBookServiceClient(conn)
	log.Printf("echo request: wang\n")
	r, err := c.AddBookInfo(ctx, &pb.BookInfo{
		BookId:   6,
		BookName: "go 微服务实战 与 回退",
	})
	if err != nil {
		log.Fatalf("could not echo: %v\n", err)
	}
	log.Printf("Echo reply: %s\n", r.GetBookId())
	booklist, err := c.GetBookList(ctx, &pb.BookListParams{
		Page:  0,
		Limit: 0,
	})
	if err != nil {
		log.Fatalf("could not echo: %v\n", err)
	}
	for _, v := range booklist.BookList {
		fmt.Println(strconv.FormatInt(int64(v.BookId), 32) + v.BookName)
	}
	bkinfo, err := c.GetBookInfo(ctx, &pb.BookInfoParams{BookId: 1})
	fmt.Println(strconv.FormatInt(int64(bkinfo.BookId), 32) + bkinfo.BookName)
}
