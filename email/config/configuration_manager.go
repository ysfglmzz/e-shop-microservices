package config

import (
	"github.com/spf13/viper"
	"github.com/ysfglmzz/e-shop-microservices/email/pkg/constants"
)

func init() {
	viper.AddConfigPath(constants.ConfigPath)
	viper.SetConfigType(constants.ConfigType)
}

type configurationManager struct {
	appConfig AppConfig
}

func NewConfigurationManager() *configurationManager {
	return &configurationManager{
		appConfig: readApplicationConfig(),
	}
}

func (c *configurationManager) GetSystemConfig() SystemConfig {
	return c.appConfig.System
}

func (c *configurationManager) GetRabbitMqConfig() RabbitMqConfig {
	return c.appConfig.RabbitMq
}

func (c *configurationManager) GetAppConfig() AppConfig {
	return c.appConfig
}

func readApplicationConfig() AppConfig {
	viper.SetConfigName(constants.AppConfigName)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	return appConfig
}
