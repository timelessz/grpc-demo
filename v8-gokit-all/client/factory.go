package client

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"go-kit-demo/v8-gokit-all/auth"
	"google.golang.org/grpc"
	"io"
)

func authFactory(_ context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {

		//println("authFactory")
		//println(instance)
		//conn, err := grpc.Dial(instance, grpc.WithInsecure())
		//var (
		//	enc grpctransport.EncodeRequestFunc
		//	dec grpctransport.DecodeResponseFunc
		//)
		//enc, dec = encodeAuthRequest, decodeAuthResponse
		//return grpctransport.NewClient(conn, serviceName, method, enc, dec, auth.LoginResponse{}).Endpoint(), conn, err

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			conn, err := grpc.Dial(instance, grpc.WithInsecure())
			if err != nil {
				return nil, err
			}
			defer conn.Close()
			c := auth.NewAuthClient(conn)
			switch request.(type) {
			case *auth.LoginRequest:
				return c.Login(ctx, request.(*auth.LoginRequest))
			case *auth.UserIdRequest:
				return c.GetUserInfo(ctx, request.(*auth.UserIdRequest))
			default:
				return nil, nil
			}
		}, nil, nil
	}
}

//  处理请求
//func encodeAuthRequest(_ context.Context, request interface{}) (interface{}, error) {
//	//println("factory encodeAuthRequest")
//	req := request.(auth.LoginRequest)
//	return &req, nil
//}
//  处理响应
//func decodeAuthResponse(_ context.Context, r interface{}) (interface{}, error) {
//	println("decodeAuthResponse")
//	println(reflect.TypeOf(r))
//	return r, nil
//}
