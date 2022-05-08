package register

import (
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

type ConsulServer struct {
	Adress string
	Port   int
	Client consul.Client
}

func NewRegister(adress string, port int) *ConsulServer {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	apiserver, err := api.NewClient(config)
	if err != nil {
		log.Println(err)
		return nil
	}
	client := consul.NewClient(apiserver)
	return &ConsulServer{Client: client, Adress: adress, Port: port}
}

func (c ConsulServer) Register(router, name, id string) bool {
	err := c.Client.Register(&api.AgentServiceRegistration{
		Port:    c.Port,
		ID:      id,
		Address: c.Adress,
		Name:    name,
		Check: &api.AgentServiceCheck{
			GRPC:     c.Adress + ":" + strconv.Itoa(c.Port) + "/" + router,
			Interval: "5s",
			Timeout:  "5s",
		},
	})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
