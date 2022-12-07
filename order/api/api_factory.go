package api

import (
	ginserver "github.com/ysfglmzz/e-shop-microservices/order/api/gin-server"
	"github.com/ysfglmzz/e-shop-microservices/order/config"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/factories"
)

type ApiFactory struct {
	cfg            config.SystemConfig
	serviceFactory factories.ServiceFactory
}

func NewApiFactory(cfg config.SystemConfig, serviceFactory factories.ServiceFactory) *ApiFactory {
	return &ApiFactory{cfg: cfg, serviceFactory: serviceFactory}
}

func (a *ApiFactory) GetApi() IApi {
	switch a.cfg.Server {
	case "gin":
		return ginserver.NewGinServer(a.serviceFactory, a.cfg)
	}
	return nil
}
