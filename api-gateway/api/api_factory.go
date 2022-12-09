package api

import (
	ginserver "github.com/ysfglmzz/e-shop-microservices/api-gateway/api/gin-server"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
	servicediscovery "github.com/ysfglmzz/e-shop-microservices/api-gateway/service-discovery"
)

type apiFactory struct {
	api                     IApi
	routesConfig            config.RoutesConfig
	serviceDiscoveryFactory servicediscovery.ServiceDiscoveryFactory
}

func NewApiFactory(routesConfig config.RoutesConfig, serviceDiscoveryFactory servicediscovery.ServiceDiscoveryFactory) *apiFactory {
	return &apiFactory{routesConfig: routesConfig, serviceDiscoveryFactory: serviceDiscoveryFactory}
}

func (a *apiFactory) GetApi() IApi {
	switch a.routesConfig.Server {
	case "gin":
		return ginserver.NewGinServer(a.routesConfig, a.serviceDiscoveryFactory.GetServiceDiscovery())
	}
	return nil
}
