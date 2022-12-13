package ginserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/ysfglmzz/e-shop-microservices/api-gateway/config"
	serviceDiscovery "github.com/ysfglmzz/e-shop-microservices/api-gateway/service-discovery"
)

type ginServer struct {
	redisClient      *redis.Client
	routesConfig     config.RoutesConfig
	serviceDiscovery serviceDiscovery.IServiceDiscovery
	router           *gin.Engine
}

func NewGinServer(
	routesConfig config.RoutesConfig,
	serviceDiscovery serviceDiscovery.IServiceDiscovery,
	redisClient *redis.Client,
) *ginServer {
	return &ginServer{routesConfig: routesConfig, serviceDiscovery: serviceDiscovery, redisClient: redisClient}
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
						location := "http://" + g.serviceDiscovery.GetServiceIp(serviceName) + ctx.Request.RequestURI
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
	// Check token status and get uuid
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	tokenString := strings.Split(auth, " ")[1]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(g.routesConfig.TokenSecretKey), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	uuid := claims["uuid"].(string)

	// Check token from redis
	_, err = g.redisClient.Get(context.Background(), uuid).Result()
	if err != nil {
		// Check token from identity service
		if err = g.checkTokenFromIdentityService(ctx, uuid, tokenString); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
	}

	ctx.Next()
}

func (g *ginServer) checkTokenFromIdentityService(ctx *gin.Context, uuid string, tokenString string) error {
	tokenControlAddr := "http://localhost:5001"

	if g.routesConfig.UseServiceDiscovery {
		tokenControlAddr = g.serviceDiscovery.GetServiceIp("identity")
	}

	req := http.Request{
		Method: http.MethodGet,
		URL:    &url.URL{Scheme: "http", Host: tokenControlAddr, Path: "/auth/" + uuid},
		Header: ctx.Request.Header,
	}

	res, err := http.DefaultClient.Do(&req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("token is not valid")
	}
	go func() {
		g.redisClient.Set(context.Background(), uuid, tokenString, 0)
	}()
	return nil
}

func (g *ginServer) run() {
	address := fmt.Sprintf("%s:%d", g.routesConfig.Host, g.routesConfig.Port)
	g.router.Run(address)
}
