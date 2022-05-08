package client

import (
	"context"
	"go-kit-demo/v8-gokit-all/auth"
	v8 "go-kit-demo/v8-gokit-all/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	svc := v8.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := svc.CreateBook(ctx, &v8.Book{Id: 1, Name: "book1", Author: "author1", Price: 1.0})
	if err != nil {
		log.Fatalf("could not put: %v", err)
	}
	log.Println(r.Id)
	login := auth.NewAuthClient(conn)
	r1, err := login.Login(ctx, &auth.LoginRequest{Username: "admin", Password: "admin"})
	if err != nil {
		log.Fatalf("could not put: %v", err)
	}
	c := grpc_health_v1.NewHealthClient(conn)
	cr, err := c.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: "book"})
	if err != nil {
		log.Fatalf("could not put: %v", err)
	}
	log.Println(cr.Status)
	cra, err := c.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: "auth"})
	if err != nil {
		log.Fatalf("could not put: %v", err)
	}
	log.Println(cra.Status)
	println(r1.Token)
}
