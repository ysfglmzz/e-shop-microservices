package ginserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
)

type ginServer struct {
	routesConfig config.RoutesConfig
	router       *gin.Engine
}

func NewGinServer(routesConfig config.RoutesConfig) *ginServer {
	return &ginServer{routesConfig: routesConfig}
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
				address := service.Address
				g.router.Handle(route.Method, route.Path, func(ctx *gin.Context) {
					ctx.Redirect(http.StatusTemporaryRedirect, address+ctx.Request.RequestURI)
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
