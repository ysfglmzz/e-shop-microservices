package ginserver

import (
	"fmt"
	"net/http"
	"net/url"

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
				handlers := []gin.HandlerFunc{}
				if route.Middleware {
					handlers = append(handlers, g.TokenControlMiddleware)
				}
				switch g.routesConfig.UseServiceDiscovery {
				case true:
					handlers = append(handlers, func(ctx *gin.Context) {
						location := g.serviceDiscovery.GetServiceIp(serviceName) + ctx.Request.RequestURI
						ctx.Redirect(http.StatusTemporaryRedirect, location)
					})
				case false:
					handlers = append(handlers, func(ctx *gin.Context) {
						location := address + ctx.Request.RequestURI
						ctx.Redirect(http.StatusTemporaryRedirect, location)
					})
				}
				g.router.Handle(route.Method, route.Path, handlers...)

			}()
		}
	}
	return g
}

func (g *ginServer) TokenControlMiddleware(ctx *gin.Context) {
	tokenControlAddr := "http://localhost:5001"
	if g.routesConfig.UseServiceDiscovery {
		tokenControlAddr = g.serviceDiscovery.GetServiceIp("identity")
	}
	fmt.Println(tokenControlAddr)
	req := http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "http", Host: "localhost:5001", Path: "/auth/checkToken"},
		Header: make(http.Header),
	}

	req.Header.Set("Authorization", ctx.GetHeader("Authorization"))
	fmt.Println(ctx.Request.Header.Get("Authorization"))

	res, err := http.DefaultClient.Do(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if res.StatusCode != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	ctx.Next()
}

func (g *ginServer) run() {
	address := fmt.Sprintf("%s:%d", g.routesConfig.Host, g.routesConfig.Port)
	g.router.Run(address)
}
