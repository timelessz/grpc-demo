package services

import (
	. "bkgrpc/proto"
	"context"
	"errors"
)

type Service interface {
	Concat(context.Context, *StringRequest) (*StringResponse, error)
	Diff(context.Context, *StringRequest) (*StringResponse, error)
	HealtStatus(context.Context, *HealthRequest) (*HealthResponse, error)
}

type ServicesA struct{}

func (s ServicesA) Concat(_ context.Context, request *StringRequest) (*StringResponse, error) {
	if len(request.A+request.B) > 10 {
		return nil, errors.New("too long strings")
	}
	return &StringResponse{Msg: request.B + request.A}, nil
}

func (s ServicesA) Diff(_ context.Context, request *StringRequest) (*StringResponse, error) {
	if len(request.A+request.B) > 10 {
		return nil, errors.New("too long strings")
	}
	if request.A == request.B {
		return &StringResponse{Msg: "two strings is same"}, nil
	} else {
		return &StringResponse{Msg: "two strings is not same"}, nil
	}
}

func (s ServicesA) HealtStatus(_ context.Context, request *HealthRequest) (*HealthResponse, error) {
	return &HealthResponse{Status: "the status is health"}, nil
}
