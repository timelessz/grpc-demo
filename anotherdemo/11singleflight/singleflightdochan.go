package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	// DoChan returns a channel that will receive the results of the function
	// when they are available. If multiple callers call DoChan with the same key,
	// the function will be called only once, and all callers will receive the
	// same results.
	const n = 20
	key := "get"
	// 初始化 ch
	ch := make(<-chan singleflight.Result, 1)
	var g singleflight.Group
	for i := 0; i < n; i++ {
		go func(j int) {
			ch = getCache(key, &g)
			select {
			case d := <-ch:
				fmt.Println(d.Val)
				fmt.Println(d.Shared)
				fmt.Println(d.Err)
				fmt.Println("DoChan done")
			case <-time.After(40 * time.Second):
				fmt.Println("Do hangs")
			}
		}(i)
	}
	select {}
}

// getCache 从缓存中获取数据
func getCache(key string, g *singleflight.Group) <-chan singleflight.Result {
	ch := g.DoChan(key, func() (interface{}, error) {
		// 读取数据库字段
		time.Sleep(3 * time.Second)
		fmt.Println("get data from data")
		return "get data demo", nil
	})
	return ch
}
