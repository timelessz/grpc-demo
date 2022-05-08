package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-kit-demo/v8-gokit-all/book"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BookService interface {
	GetBookList(ctx context.Context, bqp *book.BookQueryParams) (*book.BookList, error)
	GetOneBook(ctx context.Context, obp *book.OneBookQueryParams) (*book.Book, error)
	CreateBook(ctx context.Context, bi *book.Book) (*book.Book, error)
	BookHealthCheck(context.Context, *emptypb.Empty) (*book.HealthResponse, error)
}
type NewMiddlewareService func(server BookService) BookService

type BookServer struct {
	logger *logrus.Logger
}

// BookServer is a concrete implementation of BookService
func NewBookService(logger *logrus.Logger) BookService {
	server := BookServer{
		logger: logger,
	}
	return NewBookServiceLogMiddleware(logger)(server)
}

func (b BookServer) GetBookList(ctx context.Context, bqp *book.BookQueryParams) (*book.BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}

func (b BookServer) GetOneBook(ctx context.Context, obqp *book.OneBookQueryParams) (*book.Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBook not implemented")
}

func (b BookServer) CreateBook(ctx context.Context, book *book.Book) (*book.Book, error) {
	println("service CreateBook")
	fmt.Println(book.Id)
	return book, status.Errorf(codes.OK, "返回值正确")
}

func (b BookServer) BookHealthCheck(ctx context.Context, empty *emptypb.Empty) (*book.HealthResponse, error) {
	//TODO implement me
	return &book.HealthResponse{Status: "OK"}, nil
}
