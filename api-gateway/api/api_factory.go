package api

import (
	ginserver "github.com/ysfglmzz/e-shop-microservices/api-gateway/api/gin-server"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
)

type apiFactory struct {
	routesConfig config.RoutesConfig
	api          IApi
}

func NewApiFactory(routesConfig config.RoutesConfig) *apiFactory {
	return &apiFactory{routesConfig: routesConfig}
}

func (a *apiFactory) GetApi() IApi {
	switch a.routesConfig.Server {
	case "gin":
		return ginserver.NewGinServer(a.routesConfig)
	}
	return nil
}
