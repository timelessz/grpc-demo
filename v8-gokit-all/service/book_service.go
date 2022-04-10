package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	v8 "go-kit-demo/v8-gokit-all/book"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService interface {
	GetBookList(ctx context.Context, bqp *v8.BookQueryParams) (*v8.BookList, error)
	GetOneBook(ctx context.Context, obp *v8.OneBookQueryParams) (*v8.Book, error)
	CreateBook(ctx context.Context, bi *v8.Book) (*v8.Book, error)
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

func (b BookServer) GetBookList(ctx context.Context, bqp *v8.BookQueryParams) (*v8.BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}

func (b BookServer) GetOneBook(ctx context.Context, obqp *v8.OneBookQueryParams) (*v8.Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBook not implemented")
}

func (b BookServer) CreateBook(ctx context.Context, book *v8.Book) (*v8.Book, error) {
	println("service CreateBook")
	fmt.Println(book.Id)
	return book, status.Errorf(codes.OK, "返回值正确")
}
