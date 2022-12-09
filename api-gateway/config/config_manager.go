package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./resource")
	viper.SetConfigType("yaml")
}

type configManager struct {
	routesConfig RoutesConfig
}

func NewConfigManager() *configManager {
	return &configManager{routesConfig: readQueuesConfig()}
}

func (c *configManager) GetRoutesQunfig() RoutesConfig {
	return c.routesConfig
}

func readQueuesConfig() RoutesConfig {
	viper.SetConfigName("routes")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	var routesConfig RoutesConfig
	if err := viper.Unmarshal(&routesConfig); err != nil {
		panic(err.Error())
	}
	return routesConfig
}
