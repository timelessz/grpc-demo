package main

import (
	pb "bkgrpc/proto"
	//	"bkgrpc/services"
	"context"
	//	"flag"
	"fmt"
	"google.golang.org/grpc"
	//	"time"
)

func main() {
	ctx := context.Background()
	con, _ := grpc.Dial("127.0.0.1:80", grpc.WithInsecure())
	fu := pb.NewStringServicesClient(con)
	rep, _ := fu.Diff(ctx, &pb.StringRequest{A: "a", B: "b"})
	fmt.Println(rep)
}
