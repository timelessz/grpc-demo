package healthservice

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type AuthHealthCheckService struct{}

func (h AuthHealthCheckService) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h AuthHealthCheckService) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}
