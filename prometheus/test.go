package main

import (
	"fmt"
	"net/http"
	"time"
)

func Test() {
	for {
		time.Sleep(50 * time.Millisecond)
		send()
	}
	time.Sleep(24 * time.Hour)
}

func send() {
	resp, err := http.Get("http://192.168.2.55:8888/howie")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
}

func main() {
	Test()
}
