package registration

import (
	"github.com/ysfglmzz/e-shop-microservices/basket/config"
	"github.com/ysfglmzz/e-shop-microservices/basket/registration/consul"
)

type registrationFactory struct {
	systemConfig config.SystemConfig
}

func NewRegistrationFactory(systemConfig config.SystemConfig) *registrationFactory {
	return &registrationFactory{systemConfig: systemConfig}
}

func (r *registrationFactory) GetRegistrationService() IServiceRegistration {
	switch r.systemConfig.ServiceDiscovery {
	case "consul":
		return consul.NewConsul(r.systemConfig)
	}
	return nil
}
