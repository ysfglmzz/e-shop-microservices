package main

import (
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/api"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
)

func main() {
	configManager := config.NewConfigManager()
	routesConfig := configManager.GetRoutesQunfig()
	apiFactory := api.NewApiFactory(routesConfig)
	apiFactory.GetApi().Start()
}
