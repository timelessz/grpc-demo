package client

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
	"os"
)

func Run() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var client consul.Client
	{
		consulConfig := consulapi.DefaultConfig()
		consulConfig.Address = "192.168.2.168:8500"
		consulClient, err := consulapi.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)
	}
	ctx := context.Background()
	discovererEndpoint := MakeAuthDiscoverEndpoint(ctx, client, logger)
	r := MakeHttpHandler(discovererEndpoint)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
}
