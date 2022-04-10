package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sirupsen/logrus"
	"go-kit-demo/v8-gokit-all/auth"
	v8 "go-kit-demo/v8-gokit-all/book"
	"go-kit-demo/v8-gokit-all/endpoint"
	"google.golang.org/grpc/metadata"
)

type GrpcServer struct {
	Book   grpctransport.Handler
	Author grpctransport.Handler
	Log    *logrus.Logger
}

func NewGRPCServer(endpoints endpoint.EndPointServer, log *logrus.Logger) *GrpcServer {
	options := []grpctransport.ServerOption{
		//grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		//	grpc_ctxtags.StreamServerInterceptor(),
		//	grpc_recovery.StreamServerInterceptor(),
		//	grpc_prometheus.StreamServerInterceptor,
		//)),
		//grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		//	grpc_ctxtags.UnaryServerInterceptor(),
		//	grpc_recovery.UnaryServerInterceptor(),
		//	grpc_prometheus.UnaryServerInterceptor,
		//)),
		grpctransport.ServerBefore(func(context.Context, metadata.MD) context.Context {
			return context.WithValue(context.Background(), "traceID", "123456")
		}),
		//grpctransport.ServerErrorHandler(),
	}
	return &GrpcServer{
		Book: grpctransport.NewServer(
			endpoints.BookEndPoint,
			decodeGRPCBookRequest,
			encodeGRPCBookResponse,
			options...,
		),
		Author: grpctransport.NewServer(
			endpoints.AuthEndPoint,
			decodeGRPCAuthorRequest,
			encodeGRPCAuthorResponse,
			options...,
		),
		Log: log,
	}
}

// decodeGRPCBookRequest  book解码
func decodeGRPCBookRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	switch grpcReq.(type) {
	case *v8.BookQueryParams:
		return grpcReq.(*v8.BookQueryParams), nil
	case *v8.OneBookQueryParams:
		//return grpcReq, nil
		return grpcReq.(*v8.OneBookQueryParams), nil
	case *v8.Book:
		println("decodeGRPCBookRequest")
		return grpcReq.(*v8.Book), nil
	default:
		return nil, nil
	}
}

// GRPC Book 返回结果编码
func encodeGRPCBookResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	switch grpcResp.(type) {
	case *v8.BookList:
		return grpcResp.(*v8.BookList), nil
	case *v8.Book:
		println("encodeGRPCBookResponse")
		bk := grpcResp.(*v8.Book)
		return bk, nil
	default:
		return nil, nil
	}
}

// decodeGRPCAuthorRequest  author解码
func decodeGRPCAuthorRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	switch grpcReq.(type) {
	case *auth.LoginRequest:
		return grpcReq.(*auth.LoginRequest), nil
	case *auth.User:
		//return grpcReq, nil
		println("decodeGRPCAuthorRequest")
		return grpcReq.(*auth.User), nil
	default:
		return nil, nil
	}
}

// GRPC Author 返回结果编码
func encodeGRPCAuthorResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	switch grpcResp.(type) {
	case *auth.LoginResponse:
		return grpcResp.(*auth.LoginResponse), nil
	case *auth.User:
		println("encodeGRPCAuthorResponse")
		bk := grpcResp.(*auth.User)
		return bk, nil
	default:
		return nil, nil
	}
}

func (g GrpcServer) GetBookList(ctx context.Context, params *v8.BookQueryParams) (*v8.BookList, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, params)
	blist := resp.(*v8.BookList)
	return blist, err
}

func (g GrpcServer) GetOneBook(ctx context.Context, params *v8.OneBookQueryParams) (*v8.Book, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, params)
	b := resp.(*v8.Book)
	return b, err
}

func (g GrpcServer) CreateBook(ctx context.Context, book *v8.Book) (*v8.Book, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, book)
	if err != nil {
		println("create book transport 后", err.Error())
	}
	println("grpcserver")
	b := resp.(*v8.Book)
	return b, err
}

func (g GrpcServer) Login(ctx context.Context, params *auth.LoginRequest) (*auth.LoginResponse, error) {
	_, resp, err := g.Author.ServeGRPC(ctx, params)
	if err != nil {
		println("login", err.Error())
	}
	println("grpcserver")
	b := resp.(*auth.LoginResponse)
	return b, err
}

func (g GrpcServer) GetUserInfo(ctx context.Context, uir *auth.UserIdRequest) (*auth.User, error) {
	_, resp, err := g.Author.ServeGRPC(ctx, uir)
	if err != nil {
		println("getuserinfo", err.Error())
	}
	println("grpcserver")
	b := resp.(*auth.User)
	return b, err
}
