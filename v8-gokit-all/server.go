package main

import (
	"context"
	v2runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"go-kit-demo/v8-gokit-all/auth"
	"go-kit-demo/v8-gokit-all/book"
	"go-kit-demo/v8-gokit-all/consul"
	"go-kit-demo/v8-gokit-all/endpoint"
	"go-kit-demo/v8-gokit-all/service"
	"go-kit-demo/v8-gokit-all/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// 监听 grpc 和 http
func main() {
	wg := sync.WaitGroup{}
	log := initLogger()
	defer func() {
		log.Info("deregister service")
		DeRegisterService(log)
	}()
	// http 服务
	httpServer(&wg, log)
	bookService := service.NewBookService(log)
	authorService := service.NewLoginService(log)
	endpointserver := endpoint.NewEndpoint(bookService, authorService)
	grpcserver := transport.NewGRPCServer(endpointserver, log)
	grpcListen, err := net.Listen("tcp", ":8881")
	if err != nil {
		log.Fatal(err)
	}
	baseServer := grpc.NewServer()
	book.RegisterBookServiceServer(baseServer, grpcserver)
	auth.RegisterAuthServer(baseServer, grpcserver)
	reflection.Register(baseServer)
	log.Info("grpc server start")
	RegisterService(baseServer, log)
	if err := baseServer.Serve(grpcListen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	wg.Wait()
}

// dereister service
func DeRegisterService(log *logrus.Logger) {
	client, err := consul.NewConsulClient("127.0.0.1", 8500, log)
	if err != nil {
		log.Fatal(err)
	}
	client.DeRegister("auth")
	client.DeRegister("book")
}

// 注册服务到 consul
func RegisterService(baseServer *grpc.Server, log *logrus.Logger) {
	go func() {
		// 注册服务到 consul
		log.Info("register service")
		registerServiceToConsul(log)
	}()
}

func registerServiceToConsul(log *logrus.Logger) {
	client, err := consul.NewConsulClient("127.0.0.1", 8500, log)
	if err != nil {
		logrus.Errorf("consul client error: %v", err)
		return
	}
	bookService := &consul.ServiceInfo{
		ServiceID:       "book",
		ServiceName:     "bookServer",
		ServiceAddr:     "192.168.2.168",
		ServicePort:     8881,
		HealthCheckPort: 8081,
		HealthCheckPath: "/v8/book/health",
	}
	if client.Register(bookService) {
		log.Info("Consul register book service success")
	} else {
		log.Error("Consul register book service failed")
	}
	authService := &consul.ServiceInfo{
		ServiceID:       "auth",
		ServiceName:     "authServer",
		ServiceAddr:     "192.168.2.168",
		ServicePort:     8881,
		HealthCheckPort: 8082,
		HealthCheckPath: "/v8/auth/health",
	}
	if client.Register(authService) {
		log.Info("Consul register auth service success")
	} else {
		log.Error("Consul register auth service failed")
	}
}

// grpc-gateway implements
func httpServer(wg *sync.WaitGroup, log *logrus.Logger) {
	go func() {
		wg.Add(1)
		// httpGRPCGateway
		println("http grpc gateway start")
		errs := bookHttpListenAndServe(log)
		if errs != nil {
			log.Fatal(errs)
		}
	}()
	go func() {
		wg.Add(1)
		// grpc
		println("grpc start")
		errs := authHttpListenAndServe(log)
		if errs != nil {
			log.Fatal(errs)
		}
	}()
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	//You can set any 'io. Writer' type such as file as log output
	/*
		file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			logger.Info("Failed to log to file, using default stderr")
		}
	*/
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	return logger
}

// httpGRPCGateway
func bookHttpListenAndServe(log *logrus.Logger) error {
	println("httpGRPCGateway")
	mux := v2runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	endpoint := "192.168.2.168:8881"
	err := book.RegisterBookServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	return http.ListenAndServe(":8081", mux)
}

func authHttpListenAndServe(log *logrus.Logger) error {
	v2authmux := v2runtime.NewServeMux()
	println("authHttpGRPCGateway")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	endpoint := "192.168.2.168:8881"
	err1 := auth.RegisterAuthHandlerFromEndpoint(ctx, v2authmux, endpoint, opts)
	if err1 != nil {
		log.Fatal(err1)
	}
	return http.ListenAndServe(":8082", v2authmux)
}
