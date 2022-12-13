package registration

import (
	"github.com/ysfglmzz/e-shop-microservices/catalog/config"
	"github.com/ysfglmzz/e-shop-microservices/catalog/registration/consul"
)

type registrationFactory struct {
	appConfig config.AppConfig
}

func NewRegistrationFactory(appConfig config.AppConfig) *registrationFactory {
	return &registrationFactory{appConfig: appConfig}
}

func (r *registrationFactory) GetRegistrationService() IServiceRegistration {
	switch r.appConfig.System.ServiceDiscovery {
	case "consul":
		return consul.NewConsul(r.appConfig)
	}
	return nil
}
