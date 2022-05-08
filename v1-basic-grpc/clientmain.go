package main

import (
	"context"
	"fmt"
	"go-kit-demo/v1-basic-grpc/product"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50011"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()
	client := product.NewProductInfoClient(conn)
	ctx := context.Background()
	id := AddProduct(ctx, client)
	fmt.Println(id)
	GetProduct(ctx, client, id)
}

func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	aMac := &product.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success, id = ", productId.Value)
	return productId.Value
}

func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	productId := product.ProductId{Value: id}
	products, err := client.GetProduct(ctx, &productId)
	if err != nil {
		log.Println("GET PRODUCT ERR", err)
	}
	log.Printf("get prodcut success : %+v\n", products)
}
