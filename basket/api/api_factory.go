package api

import (
	ginserver "github.com/ysfglmzz/e-shop-microservices/basket/api/gin-server"
	"github.com/ysfglmzz/e-shop-microservices/basket/config"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/factories"
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
