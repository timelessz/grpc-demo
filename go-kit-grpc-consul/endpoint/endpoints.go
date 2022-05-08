package endpoint

import (
	"bkgrpc/proto"
	"bkgrpc/services"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndpointA struct {
	ConcatEndpoint endpoint.Endpoint
	DiffEndpoint   endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

func MakeConcatEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*gproto.StringRequest)
		response, err = svc.Concat(ctx, req)
		return
	}
}

func MakeDiffEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*gproto.StringRequest)
		response, err = svc.Diff(ctx, req)
		return
	}
}

func MakeHealthEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*gproto.HealthRequest)
		response, err = svc.HealtStatus(ctx, req)
		return
	}
}
