package api

import (
	"fmt"

	redis "github.com/go-redis/redis/v8"

	ginserver "github.com/ysfglmzz/e-shop-microservices/api-gateway/api/gin-server"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
	servicediscovery "github.com/ysfglmzz/e-shop-microservices/api-gateway/service-discovery"
)

type apiFactory struct {
	routesConfig            config.RoutesConfig
	serviceDiscoveryFactory servicediscovery.ServiceDiscoveryFactory
}

func NewApiFactory(routesConfig config.RoutesConfig, serviceDiscoveryFactory servicediscovery.ServiceDiscoveryFactory) *apiFactory {
	return &apiFactory{routesConfig: routesConfig, serviceDiscoveryFactory: serviceDiscoveryFactory}
}

func (a *apiFactory) GetApi() IApi {
	switch a.routesConfig.Server {
	case "gin":
		return ginserver.NewGinServer(a.routesConfig, a.serviceDiscoveryFactory.GetServiceDiscovery(), a.GetRedisClient())
	}
	return nil
}

func (a *apiFactory) GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", a.routesConfig.Redis.Host, a.routesConfig.Redis.Port),
		Password: "",
		DB:       0,
	})
}
