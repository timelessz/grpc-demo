package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-demo/v8-gokit-all/auth"
	v8 "go-kit-demo/v8-gokit-all/book"
	"go-kit-demo/v8-gokit-all/service"
)

type EndPointServer struct {
	BookEndPoint endpoint.Endpoint
	AuthEndPoint endpoint.Endpoint
}

// 构建 BookEndPoint
func makeBookEndpoint(service service.BookService, mws map[string][]endpoint.Middleware) endpoint.Endpoint {
	endpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		switch request.(type) {
		case *v8.BookQueryParams:
			bq := request.(*v8.BookQueryParams)
			return service.GetBookList(ctx, bq)
		case *v8.OneBookQueryParams:
			obq := request.(*v8.OneBookQueryParams)
			return service.GetOneBook(ctx, obq)
		case *v8.Book:
			b := request.(*v8.Book)
			println("endpoint", b.Name)
			return service.CreateBook(ctx, b)
		default:
			println("错误")
			println(request)
			return nil, nil
		}
	}
	for _, m := range mws["loginMW"] {
		endpoint = m(endpoint)
	}
	return endpoint
}

// 构建授权登录 EndPoint
func makeAuthEndpoint(service service.AuthService, mws map[string][]endpoint.Middleware) endpoint.Endpoint {
	endpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		switch request.(type) {
		case *auth.LoginRequest:
			println("makeAuthEndpoint")
			bq := request.(*auth.LoginRequest)
			return service.Login(ctx, bq)
		case *auth.UserIdRequest:
			println("makeAuthEndpoint")
			obq := request.(*auth.UserIdRequest)
			return service.GetUserInfo(ctx, obq)
		default:
			return nil, nil
		}
	}
	for _, m := range mws["loginMW"] {
		endpoint = m(endpoint)
	}
	return endpoint
}

// 构建 	EndPointServer
func NewEndpoint(bookService service.BookService, authorService service.AuthService) EndPointServer {
	loginMW := func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// endpoint 中间件
			println("endpoint 中间件")
			return next(ctx, request)
		}
	}
	mws := map[string][]endpoint.Middleware{
		"loginMW": {
			loginMW,
		},
	}
	return EndPointServer{BookEndPoint: makeBookEndpoint(bookService, mws), AuthEndPoint: makeAuthEndpoint(authorService, mws)}
}
