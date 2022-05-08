package transport

import (
	gproto "bkgrpc/proto"
	"context"
)

func DecodString(ctx context.Context, req interface{}) (request interface{}, err error) {
	request = req.(*gproto.StringRequest)
	return
}

func DecodeHealth(ctx context.Context, req interface{}) (request interface{}, err error) {
	request = req.(*gproto.HealthRequest)
	return
}
func HealthEncode(ctx context.Context, rep interface{}) (response interface{}, err error) {
	response = rep.(*gproto.HealthResponse)
	return
}
func EncodString(ctx context.Context, rep interface{}) (response interface{}, err error) {
	response = rep.(*gproto.StringResponse)
	return
}
