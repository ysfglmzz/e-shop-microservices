package ginserver

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ysfglmzz/e-shop-microservices/order/api/gin-server/middleware"
	"github.com/ysfglmzz/e-shop-microservices/order/config"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/factories"
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
		generateOrderGroup().
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

func (g *GinServer) generateOrderGroup() *GinServer {
	productApi := NewOrderApi(g.serviceFactory.GetOrderService(), g.logger)
	routerGroup := g.router.Group("orders")
	routerGroup.GET("/:userId", middleware.Authorization("customer"), productApi.GetOrderByUserId)
	routerGroup.PUT("/:id/complete", productApi.SetStatusOrderCompleted)
	return g
}

func (g *GinServer) listen() {
	address := fmt.Sprintf("%s:%d", g.cfg.Host, g.cfg.Port)
	g.router.Run(address)
}
