package main

import (
	"github.com/ysfglmzz/e-shop-microservices/order/api"
	"github.com/ysfglmzz/e-shop-microservices/order/config"
	_ "github.com/ysfglmzz/e-shop-microservices/order/docs"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/factories"
)

// @title Order Service Api
// @version 1.0
// @description Order Service documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	configManager := config.NewConfigurationManager()
	appConfig := configManager.GetAppConfig()
	systemConfig := configManager.GetSystemConfig()
	queuesConfig := configManager.GetQueuesConfig()
	connectionFactory := factories.NewConnectionFactory(appConfig).ConnectDb().EventBusConnect()
	repositoryFactory := factories.NewRepositoryFactory(appConfig, *connectionFactory)
	eventBusFactory := factories.NewEventBusFactory(systemConfig, queuesConfig, *connectionFactory, *repositoryFactory)
	serviceFactory := factories.NewServiceFactory(*repositoryFactory, *eventBusFactory)
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory)
	eventBusFactory.GetEventBus().Subscribe()
	apiFactory.GetApi().Start()
}
