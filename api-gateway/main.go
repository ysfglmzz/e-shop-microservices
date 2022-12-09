package main

import (
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/api"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
	servicediscovery "github.com/ysfglmzz/e-shop-microservices/api-gateway/service-discovery"
)

func main() {
	configManager := config.NewConfigManager()
	routesConfig := configManager.GetRoutesQunfig()
	if routesConfig.UseServiceDiscovery {

	}
	serviceDiscoveryFactory := servicediscovery.NewServiceDiscoveryFactory(routesConfig)

	apiFactory := api.NewApiFactory(routesConfig, *serviceDiscoveryFactory)
	apiFactory.GetApi().Start()
}
