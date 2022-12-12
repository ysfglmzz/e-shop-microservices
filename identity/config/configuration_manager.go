package config

import (
	"github.com/spf13/viper"
	"github.com/ysfglmzz/e-shop-microservices/identity/pkg/constants"
)

func init() {
	viper.AddConfigPath(constants.ConfigPath)
	viper.SetConfigType(constants.ConfigType)
}

type ConfigurationManager struct {
	appConfig AppConfig
}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{appConfig: readApplicationConfig()}
}

func (c *ConfigurationManager) GetMysqlConfig() MysqlConfig {
	return c.appConfig.Mysql
}

func (c *ConfigurationManager) GetSystemConfig() SystemConfig {
	return c.appConfig.System
}

func (c *ConfigurationManager) GetRabbitMqConfig() RabbitMqConfig {
	return c.appConfig.RabbitMq
}

func (c *ConfigurationManager) GetAppConfig() AppConfig {
	return c.appConfig
}

func readApplicationConfig() AppConfig {
	viper.SetConfigName(constants.ConfigName)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	constants.TokenSecretKey = appConfig.System.TokenSecretKey
	return appConfig
}
