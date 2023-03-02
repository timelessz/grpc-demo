package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync/atomic"
	"time"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	time.Sleep(3 * time.Second)
	fmt.Println("get data from db demo")
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	var g singleflight.Group
	const n = 10
	waited := int32(n)
	done := make(chan struct{})
	key := "https://weibo.com/1227368500/H3GIgngon"
	for i := 0; i < n; i++ {
		go func(j int) {
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}
		}(i)
	}

	select {
	case <-done:
	case <-time.After(20 * time.Second):
		fmt.Println("Do hangs")
	}
}
