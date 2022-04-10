package consul

import (
	"github.com/go-kit/kit/sd/consul"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"sync"
)

type ConsulClient struct {
	consulHost  string //consul的地址
	consulPort  int    //consul的端口
	Client      consul.Client
	Config      *consulapi.Config //kit封装的consul配置
	mutex       sync.Mutex
	instanceMap sync.Map
	Log         *logrus.Logger
}

type ServiceInfo struct {
	ServiceID              string
	ServiceName            string
	ServiceAddr            string
	ServicePort            int
	HealthCheckServiceName string
}

func NewKitDiscoverClient(consulHost string, consulPort int) (*ConsulClient, error) {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &ConsulClient{
		consulHost: consulHost,
		consulPort: consulPort,
		Client:     client,
		Config:     consulConfig,
	}, nil
}

// 生成 consul client
func NewConsulClient(consulHost string, consulPort int, log *logrus.Logger) (*ConsulClient, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &ConsulClient{
		consulHost: consulHost,
		consulPort: consulPort,
		Client:     client,
		Config:     config,
		Log:        log,
	}, nil
}

// service register in consul
func (c ConsulClient) Register(set *ServiceInfo) bool {
	err := c.Client.Register(&consulapi.AgentServiceRegistration{
		ID:      set.ServiceID,
		Name:    set.ServiceName,
		Address: set.ServiceAddr,
		Port:    set.ServicePort,
		Check: &consulapi.AgentServiceCheck{
			GRPC:     set.ServiceAddr + ":" + strconv.Itoa(set.ServicePort) + "/" + set.HealthCheckServiceName,
			Interval: "10s",
			Timeout:  "5s",
		},
	})
	if err != nil {
		return false
	}
	return true
}

// service deregister in consul
func (c *ConsulClient) DeRegister(instanceId string) bool {
	serviceRegistration := &consulapi.AgentServiceRegistration{
		ID: instanceId,
	}
	err := c.Client.Deregister(serviceRegistration)
	if err != nil {
		c.Log.Println("Deregister Service Error!")
		return false
	}
	c.Log.Println("Deregister Service Success!")
	return true
}

// discover service from consul
func (c *ConsulClient) DiscoveryServices(serviceName string, logger *log.Logger) []interface{} {
	//判断服务是否已缓存
	instanceList, ok := c.instanceMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	//加锁后在判断一次，服务是否已缓存
	instanceList, ok = c.instanceMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	//响应服务变更通知，更新服务map
	go func() {
		params := make(map[string]interface{})
		params["type"] = "service"
		params["service"] = serviceName
		plan, _ := watch.Parse(params)
		plan.Handler = func(u uint64, i interface{}) {
			if i == nil {
				return
			}
			v, ok := i.([]*consulapi.ServiceEntry)
			if !ok {
				return
			}
			if len(v) == 0 {
				c.instanceMap.Store(serviceName, []interface{}{})
			}
			var healthServices []interface{}
			for _, service := range v {
				if service.Checks.AggregatedStatus() == consulapi.HealthPassing {
					healthServices = append(healthServices, service)
				}
			}
			c.instanceMap.Store(serviceName, healthServices)
		}
		defer plan.Stop()
		plan.Run(c.Config.Address)
	}()
	//调用go-kit库向consul获取服务
	entries, _, err := c.Client.Service(serviceName, "", false, nil)
	if err != nil {
		c.instanceMap.Store(serviceName, []interface{}{})
		c.Log.Println("Discover Service Error")
		return nil
	}
	instances := make([]interface{}, 0, len(entries))
	for _, instance := range entries {
		instances = append(instances, instance)
	}
	c.instanceMap.Store(serviceName, instances)
	return instances
}
