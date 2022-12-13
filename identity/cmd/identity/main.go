package main

import (
	_ "github.com/ysfglmzz/e-shop-microservices/identity/docs"

	"github.com/ysfglmzz/e-shop-microservices/identity/api"
	"github.com/ysfglmzz/e-shop-microservices/identity/config"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/factories"
	"github.com/ysfglmzz/e-shop-microservices/identity/registration"
)

// @title Identity Service Api
// @version 1.0
// @description Identity Service documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	configurationManager := config.NewConfigurationManager()
	systemConfig := configurationManager.GetSystemConfig()
	appConfig := configurationManager.GetAppConfig()
	connectionFactory := factories.NewConnectionFactory(appConfig).DbConnect().MessageBusConnect()
	repositoryFactory := factories.NewRepositoryFactory(appConfig, *connectionFactory)
	serviceFactory := factories.NewServiceFactory(*connectionFactory, *repositoryFactory, appConfig)
	registrationFactory := registration.NewRegistrationFactory(appConfig)
	registrationFactory.GetRegistrationService().Register()
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory).GetApi()
	apiFactory.Start()
}
