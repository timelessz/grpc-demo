package client

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"time"
)

func MakeAuthDiscoverEndpoint(ctx context.Context, client consul.Client, logger log.Logger) endpoint.Endpoint {
	serviceName := "authServer"
	tags := []string{}
	passingOnly := true
	duration := 500 * time.Millisecond
	instancer := consul.NewInstancer(client, logger, serviceName, tags, passingOnly)
	factory := authFactory(ctx)
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, duration, balancer)
	return retry
}
