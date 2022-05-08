package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sirupsen/logrus"
	"go-kit-demo/v8-gokit-all/auth"
	"go-kit-demo/v8-gokit-all/book"
	"go-kit-demo/v8-gokit-all/endpoint"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	Book   grpctransport.Handler
	Author grpctransport.Handler
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
	}
}

// decodeGRPCBookRequest  book解码
func decodeGRPCBookRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	switch grpcReq.(type) {
	case *book.BookQueryParams:
		return grpcReq.(*book.BookQueryParams), nil
	case *book.OneBookQueryParams:
		//return grpcReq, nil
		return grpcReq.(*book.OneBookQueryParams), nil
	case *book.Book:
		println("decodeGRPCBookRequest")
		return grpcReq.(*book.Book), nil
	case *emptypb.Empty:
		println("decodeGRPCBookResponse")
		bk := grpcReq.(*emptypb.Empty)
		return bk, nil
	default:
		return nil, nil
	}
}

// GRPC Book 返回结果编码
func encodeGRPCBookResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	switch grpcResp.(type) {
	case *book.BookList:
		return grpcResp.(*book.BookList), nil
	case *book.Book:
		println("encodeGRPCBookResponse")
		bk := grpcResp.(*book.Book)
		return bk, nil
	case *book.HealthResponse:
		return grpcResp.(*book.HealthResponse), nil
	default:
		return nil, nil
	}
}

// decodeGRPCAuthorRequest  author解码
func decodeGRPCAuthorRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	switch grpcReq.(type) {
	case *auth.LoginRequest:
		return grpcReq.(*auth.LoginRequest), nil
	case *auth.UserIdRequest:
		//return grpcReq, nil
		println("decodeGRPCAuthorRequest")
		return grpcReq.(*auth.UserIdRequest), nil
	case *emptypb.Empty:
		println("decodeGRPCAuthorRequest")
		bk := grpcReq.(*emptypb.Empty)
		return bk, nil
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
	case *emptypb.Empty:
		println("encodeGRPCAuthorResponse")
		bk := grpcResp.(*emptypb.Empty)
		return bk, nil
	case *auth.HealthResponse:
		return grpcResp.(*auth.HealthResponse), nil
	default:
		return nil, nil
	}
}

func (g GrpcServer) GetBookList(ctx context.Context, params *book.BookQueryParams) (*book.BookList, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, params)
	blist := resp.(*book.BookList)
	return blist, err
}

func (g GrpcServer) GetOneBook(ctx context.Context, params *book.OneBookQueryParams) (*book.Book, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, params)
	b := resp.(*book.Book)
	return b, err
}

func (g GrpcServer) CreateBook(ctx context.Context, b *book.Book) (*book.Book, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, b)
	if err != nil {
		println("create book transport 后", err.Error())
	}
	println("grpcserver")
	bk := resp.(*book.Book)
	return bk, err
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

func (g GrpcServer) AuthHealthCheck(ctx context.Context, empty *emptypb.Empty) (*auth.HealthResponse, error) {
	_, resp, err := g.Author.ServeGRPC(ctx, empty)
	if err != nil {
		println("getuserinfo", err.Error())
	}
	println("grpcserver")
	println(resp)
	b := resp.(*auth.HealthResponse)
	println(b)
	return b, err
}

func (g GrpcServer) BookHealthCheck(ctx context.Context, empty *emptypb.Empty) (*book.HealthResponse, error) {
	_, resp, err := g.Book.ServeGRPC(ctx, empty)
	if err != nil {
		println("getuserinfo", err.Error())
	}
	println("grpcserver")
	b := resp.(*book.HealthResponse)
	return b, err
}
