package main

import (
	"github.com/ysfglmzz/e-shop-microservices/email/config"
	"github.com/ysfglmzz/e-shop-microservices/email/internal/factory"
)

func main() {
	configManager := config.NewConfigurationManager()
	eventBusFactory := factory.NewEventBusFactory(configManager.GetAppConfig())
	eventBusFactory.GetEventBus().Subscribe()
}
