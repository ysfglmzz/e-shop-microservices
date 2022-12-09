package ginserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
	serviceDiscovery "github.com/ysfglmzz/e-shop-microservices/api-gateway/service-discovery"
)

type ginServer struct {
	routesConfig     config.RoutesConfig
	serviceDiscovery serviceDiscovery.IServiceDiscovery
	router           *gin.Engine
}

func NewGinServer(
	routesConfig config.RoutesConfig,
	serviceDiscovery serviceDiscovery.IServiceDiscovery,
) *ginServer {
	return &ginServer{routesConfig: routesConfig, serviceDiscovery: serviceDiscovery}
}

func (g *ginServer) Start() {
	g.create().
		generateRouters().
		run()
}

func (g *ginServer) create() *ginServer {
	g.router = gin.Default()
	g.router.SetTrustedProxies([]string{"localhost"})
	return g
}

func (g *ginServer) generateRouters() *ginServer {
	for _, service := range g.routesConfig.Services {
		for _, route := range service.Routes {
			func() {
				serviceName := service.Name
				address := service.Address
				if g.routesConfig.UseServiceDiscovery {
					g.router.Handle(route.Method, route.Path, func(ctx *gin.Context) {
						location := g.serviceDiscovery.GetServiceIp(serviceName) + ctx.Request.RequestURI
						ctx.Redirect(http.StatusTemporaryRedirect, location)
					})
					return
				}

				g.router.Handle(route.Method, route.Path, func(ctx *gin.Context) {
					location := address + ctx.Request.RequestURI
					ctx.Redirect(http.StatusTemporaryRedirect, location)
				})

			}()
		}
	}
	return g
}

func (g *ginServer) run() {
	address := fmt.Sprintf("%s:%d", g.routesConfig.Host, g.routesConfig.Port)
	g.router.Run(address)
}
