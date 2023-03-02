package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- 3
	}()

	select {
	case i := <-ch1:
		fmt.Printf("从ch1读取了数据%d", i)
	case j := <-ch2:
		fmt.Printf("从ch2读取了数据%d", j)
	case <-time.After(time.Second * 2):
		fmt.Println("阻塞两秒后，超时了")
	}
}
