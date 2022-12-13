package consul

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/ysfglmzz/e-shop-microservices/basket/config"
)

type consul struct {
	appCofig config.AppConfig
}

func NewConsul(appCofig config.AppConfig) *consul {
	return &consul{appCofig: appCofig}
}

func (c *consul) Register() {
	config := consulapi.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", c.appCofig.Consul.Host, c.appCofig.Consul.Port)
	consul, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}

	serviceID := "basket"

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceID,
		Port:    c.appCofig.System.Port,
		Address: c.appCofig.System.Host,
	}

	if err = consul.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}
}
