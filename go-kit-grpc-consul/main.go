package main

import (
	"bkgrpc/endpoint"
	"bkgrpc/healthservice"
	"bkgrpc/proto"
	"bkgrpc/register"
	"bkgrpc/router"
	"bkgrpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main() {
	svc := services.ServicesA{}
	endpoints := endpoint.EndpointA{
		ConcatEndpoint: endpoint.MakeConcatEndpoint(svc),
		DiffEndpoint:   endpoint.MakeDiffEndpoint(svc),
		HealthEndpoint: endpoint.MakeHealthEndpoint(svc),
	}
	r := router.NewRouter(endpoints)
	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Println(err)
		return
	}
	grpcserver := grpc.NewServer()
	gproto.RegisterStringServicesServer(grpcserver, r)
	//register into consul
	client := register.NewRegister("127.0.0.1", 8085)
	//service check
	c := healthservice.Service{}
	grpc_health_v1.RegisterHealthServer(grpcserver, &c)
	client.Register("Service", "testname", "testid")
	grpcserver.Serve(lis)
}
