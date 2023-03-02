package calculate

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type EndPointServer struct {
	CalculateEndPoint endpoint.Endpoint
}

func MakeEndpoint(s Service, mdw map[string][]endpoint.Middleware) endpoint.Endpoint {
	endpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		in := request.(CalculateIn)
		ack := s.Calculate(ctx, in)
		return ack, nil
	}
	for _, m := range mdw["Foo"] {
		endpoint = m(endpoint)
	}
	return endpoint
}

func NewEndPointServer(svc Service) EndPointServer {
	//var logger log.Logger
	// ENDPOINT 中间件定义
	logeMW := func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// something TODO about log
			// endpoint 中间件， 执行服务某项操作之前执行操作
			// ctx = ...
			// request = ...
			//logger.Log("...")
			data, ok := request.(CalculateIn)
			// 计算类型不为  add 时，返回错误
			if data.Type != "add" && ok {
				return nil, errors.New("type is not add")
			}
			return next(ctx, request)
		}
	}
	emws := map[string][]endpoint.Middleware{}
	emws["Foo"] = []endpoint.Middleware{logeMW}
	cal := MakeEndpoint(svc, emws)
	return EndPointServer{CalculateEndPoint: cal}
}

/*
type Error struct {
	Msg error
}

func (e Error) Failed() error {
	return e.Msg
}*/
