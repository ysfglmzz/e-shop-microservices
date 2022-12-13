package consul

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/ysfglmzz/e-shop-microservices/identity/config"
)

type consul struct {
	appConfig config.AppConfig
}

func NewConsul(appConfig config.AppConfig) *consul {
	return &consul{appConfig: appConfig}
}

func (c *consul) Register() {
	config := consulapi.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", c.appConfig.Consul.Host, c.appConfig.Consul.Port)
	consul, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}

	serviceID := "identity"

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceID,
		Port:    c.appConfig.System.Port,
		Address: c.appConfig.System.Host,
	}

	if err = consul.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}
}
