package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

const (
	consulAddress = "192.168.2.130:30098"
	localIp       = "192.168.2.136"
	localPort     = 8888
)

func consulRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}
	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "1"
	registration.Name = "calculate"
	registration.Port = localPort
	registration.Tags = []string{"calculate"}
	registration.Address = localIp
	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/health", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check
	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

func main() {
	consulRegister()
}
