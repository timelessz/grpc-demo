package router

import (
	"bkgrpc/endpoint"
	. "bkgrpc/proto"
	"bkgrpc/services"
	"bkgrpc/transport"
	"context"
	"github.com/go-kit/kit/transport/grpc"
)

type ServicesW interface {
	Concat(context.Context, *StringRequest) (*StringResponse, error)
	Diff(context.Context, *StringRequest) (*StringResponse, error)
	HealtStatus(context.Context, *HealthRequest) (*HealthResponse, error)
}
type ServicesWA struct {
	concat      grpc.Handler
	diff        grpc.Handler
	healtStatus grpc.Handler
}

func (s ServicesWA) Concat(ctx context.Context, request *StringRequest) (*StringResponse, error) {
	_, rep, err := s.concat.ServeGRPC(ctx, request)
	return rep.(*StringResponse), err
}

func (s ServicesWA) Diff(ctx context.Context, request *StringRequest) (*StringResponse, error) {
	_, rep, err := s.diff.ServeGRPC(ctx, request)
	return rep.(*StringResponse), err
}

func (s ServicesWA) HealtStatus(ctx context.Context, request *HealthRequest) (*HealthResponse, error) {
	_, rep, err := s.healtStatus.ServeGRPC(ctx, request)
	return rep.(*HealthResponse), err
}

func NewRouter(endpoints endpoint.EndpointA) services.Service {
	return &ServicesWA{
		healtStatus: grpc.NewServer(endpoints.HealthEndpoint, transport.DecodeHealth, transport.HealthEncode),
		diff:        grpc.NewServer(endpoints.DiffEndpoint, transport.DecodString, transport.EncodString),
		concat:      grpc.NewServer(endpoints.ConcatEndpoint, transport.DecodString, transport.EncodString),
	}

}
