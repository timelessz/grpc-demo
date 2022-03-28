package calculate

import (
	"context"
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
	// END POINT 中间件定义
	logeMW := func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// something TODO about log
			// ctx = ...
			// request = ...
			//logger.Log("...")
			return next(ctx, request)
		}
	}
	emws := map[string][]endpoint.Middleware{}
	emws["Foo"] = []endpoint.Middleware{logeMW}
	cal := MakeEndpoint(svc, emws)
	return EndPointServer{CalculateEndPoint: cal}
}
