package main

import (
	"fmt"
	"github.com/go-kit/log"
	calculate2 "go-kit-demo/v2-basic-http-go-kit/calculate"
	"net/http"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	svc := calculate2.NewService()
	s := calculate2.NewCalculateMiddleware(logger)(svc)
	e := calculate2.NewEndPointServer(s)
	r := calculate2.NewHttpHandler(e)
	fmt.Println("server run 0.0.0.0:8888")
	http.ListenAndServe("0.0.0.0:8008", r)
}
