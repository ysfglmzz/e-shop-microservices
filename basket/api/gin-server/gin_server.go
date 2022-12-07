package ginserver

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ysfglmzz/e-shop-microservices/basket/config"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/factories"
)

type GinServer struct {
	cfg            config.SystemConfig
	logger         *logrus.Logger
	router         *gin.Engine
	serviceFactory factories.ServiceFactory
}

func NewGinServer(serviceFactory factories.ServiceFactory, cfg config.SystemConfig) *GinServer {
	return &GinServer{serviceFactory: serviceFactory, cfg: cfg}
}

func (g *GinServer) Start() {
	g.create().
		generateSwagger().
		generateLogger().
		generateBasketGroup().
		listen()
}

func (g *GinServer) create() *GinServer {
	g.router = gin.Default()
	return g
}
func (g *GinServer) generateLogger() *GinServer {
	g.logger = logrus.New()
	g.logger.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	g.logger.SetOutput(file)
	return g

}
func (g *GinServer) generateSwagger() *GinServer {
	g.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return g
}

func (g *GinServer) generateBasketGroup() *GinServer {
	basketApi := NewBasketApi(g.serviceFactory.GetBasketService(), g.logger)
	routerGroup := g.router.Group("baskets")
	routerGroup.GET("/:userId", basketApi.GetBasketByUserID)
	routerGroup.POST("/addProduct", basketApi.AddProductToBasket)
	routerGroup.PUT("/:userId/verify", basketApi.VerifyBasketByUserId)
	return g
}

func (g *GinServer) listen() {
	address := fmt.Sprintf("%s:%d", g.cfg.Host, g.cfg.Port)
	g.router.Run(address)
}