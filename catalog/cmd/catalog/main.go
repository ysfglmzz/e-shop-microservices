package main

import (
	_ "github.com/ysfglmzz/e-shop-microservices/catalog/docs"

	"github.com/ysfglmzz/e-shop-microservices/catalog/api"
	"github.com/ysfglmzz/e-shop-microservices/catalog/config"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/factories"
)

// @title Catalog Service Api
// @version 1.0
// @description Catalog Service documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	configurationManager := config.NewConfigurationManager()
	systemConfig := configurationManager.GetSystemConfig()
	appConfig := configurationManager.GetAppConfig()
	connectionFactory := factories.NewConnectionFactory(appConfig).DbConnect()
	repositoryFactory := factories.NewRepositoryFactory(appConfig, *connectionFactory)
	serviceFactory := factories.NewServiceFactory(*repositoryFactory)
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory).GetApi()
	apiFactory.Start()
}
