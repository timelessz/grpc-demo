package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			return
		default:
		}
		fmt.Println("worker...")
		time.Sleep(time.Second * 1)
	}
}

//
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go worker(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	fmt.Println("over ...")
}
