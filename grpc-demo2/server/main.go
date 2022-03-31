package main

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go-kit-demo/grpc-demo2/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"strings"
)

type server struct {
	orderMap map[string]*order.Order
}

func (s *server) GetOrder(ctx context.Context, req *wrappers.StringValue) (resp *order.Order, err error) {
	resp = &order.Order{}
	id := req.Value
	var exist bool
	if resp, exist = s.orderMap[id]; !exist {
		err = status.Error(codes.NotFound, "order not found id = "+id)
		return
	}
	return
}

func (s *server) AddServer(ctx context.Context, req *order.Order) (resp *wrappers.StringValue, err error) {
	resp = &wrappers.StringValue{}
	if s.orderMap == nil {
		s.orderMap = make(map[string]*order.Order)
	}
	v4, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "gen uuid err", err)
	}
	id := v4.String()
	req.Id = id
	s.orderMap[id] = req
	resp.Value = id
	return
}

//搜索订单
func (s *server) SearchOrder(searchKey *wrappers.StringValue, stream order.OrderManagement_SearchOrderServer) (err error) {
	for _, val := range s.orderMap {
		for _, item := range val.Items {
			if strings.Contains(item, searchKey.Value) {
				err = stream.Send(val)
				if err != nil {
					log.Println("stream send order err.", err)
					return
				}
				break
			}
		}
	}
	return
}

//更新订单
func (s *server) UpdateOrder(stream order.OrderManagement_UpdateOrderServer) (err error) {
	updatedIds := "updated order ids : "
	for {
		val, err := stream.Recv()
		if err == io.EOF {
			//完成读取订单流
			//向客户端发送消息
			return stream.SendAndClose(&wrappers.StringValue{Value: updatedIds})
		}
		s.orderMap[val.Id] = val
		log.Println("[server]update the order : ", val.Id)
		updatedIds += val.Id + ", "
	}
}
