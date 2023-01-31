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
			return
		default:
		}
		fmt.Println("worker...")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("over ...")
}
