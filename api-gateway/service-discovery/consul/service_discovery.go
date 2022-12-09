package consul

import (
	"fmt"
	"os"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
)

type consul struct {
	consulConfig config.Consul
	client       *consulApi.Client
}

func NewConsul(consulConfig config.Consul) *consul {
	config := &consulApi.Config{
		Address: fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port),
		Scheme:  "http",
	}
	c, err := consulApi.NewClient(config)
	if err != nil {
		os.Exit(1)
	}
	return &consul{consulConfig: consulConfig, client: c}
}

func (c *consul) GetServiceIp(serviceName string) string {
	agendService, _, err := c.client.Agent().Service(serviceName, &consulApi.QueryOptions{UseCache: true})
	if err != nil {
		os.Exit(1)
	}
	return fmt.Sprintf("http://%s:%v", agendService.Address, agendService.Port)
}
