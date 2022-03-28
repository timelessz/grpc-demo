package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "go-kit-demo/v3-grpc-grpc-gateway/book"
	"go-kit-demo/v3-grpc-grpc-gateway/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
)

const (
	port    = ":8088"
	hashkey = "booklist"
)

type server struct {
}

func (s *server) AddBookInfo(ctx context.Context, bi *pb.BookInfo) (*pb.BookInfoParams, error) {
	book_id := bi.BookId
	book_name := bi.BookName
	key := strconv.FormatInt(int64(book_id), 10)
	exists := redis.REDIS.HExists(ctx, hashkey, key)
	if exists.Val() {
		// 存在的情况
		return nil, status.Errorf(codes.AlreadyExists, "该书已经存在")
	}
	bookinfo := pb.BookInfo{
		BookId:   book_id,
		BookName: book_name,
	}
	fmt.Println(bookinfo)
	jbookinfo, err := json.Marshal(bookinfo)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "序列化数据失败")
	}
	intcmd := redis.REDIS.HSetNX(ctx, hashkey, key, jbookinfo)
	if !intcmd.Val() {
		//添加失败
		return nil, status.Errorf(codes.Unavailable, "书籍添加失败")
	}
	return &pb.BookInfoParams{
		BookId: book_id,
	}, status.Errorf(codes.OK, "数据添加成功")
}

func (s *server) GetBookInfo(ctx context.Context, bid *pb.BookInfoParams) (*pb.BookInfo, error) {
	book_id := strconv.FormatInt(int64(bid.GetBookId()), 10)
	exists := redis.REDIS.HExists(ctx, hashkey, book_id)
	if !exists.Val() {
		// 存在的情况
		return nil, status.Errorf(codes.AlreadyExists, "该书不存在")
	}
	value := redis.REDIS.HGet(ctx, hashkey, book_id)
	bookinfo := new(pb.BookInfo)
	bytes, _ := value.Bytes()
	json.Unmarshal(bytes, &bookinfo)
	return bookinfo, status.Errorf(codes.OK, "")
}

func (s *server) GetBookList(ctx context.Context, bbl *pb.BookListParams) (*pb.BookList, error) {
	kvs, _ := redis.REDIS.HGetAll(ctx, hashkey).Result()
	var bl []*pb.BookInfo
	for _, v := range kvs {
		bookinfo := new(pb.BookInfo)
		json.Unmarshal([]byte(v), &bookinfo)
		bl = append(bl, bookinfo)
	}
	return &pb.BookList{
		BookList: bl,
	}, status.Errorf(codes.OK, "")
}

func main() {
	redis.InitRedis()
	bookinfo := pb.BookInfo{
		BookId:   1,
		BookName: "cmd info",
	}
	jbookinfo, err := json.Marshal(bookinfo)
	var ctx = context.Background()
	intcmd := redis.REDIS.HSet(ctx, hashkey, "demo", jbookinfo)
	fmt.Println(intcmd.Val())
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
