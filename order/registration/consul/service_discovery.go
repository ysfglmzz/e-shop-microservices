package consul

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/ysfglmzz/e-shop-microservices/order/config"
)

type consul struct {
	systemConfig config.SystemConfig
}

func NewConsul(systemConfig config.SystemConfig) *consul {
	return &consul{systemConfig: systemConfig}
}

func (c *consul) Register() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}

	serviceID := "order"

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceID,
		Port:    c.systemConfig.Port,
		Address: c.systemConfig.Host,
	}

	if err = consul.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}
}
