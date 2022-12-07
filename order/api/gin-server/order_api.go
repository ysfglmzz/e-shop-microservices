package ginserver

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	service "github.com/ysfglmzz/e-shop-microservices/order/internal/app/service"
)

type orderApi struct {
	logger       *logrus.Logger
	orderService service.IOrderService
}

func NewOrderApi(orderService service.IOrderService, logger *logrus.Logger) *orderApi {
	return &orderApi{orderService: orderService, logger: logger}
}

// @Tags Order Api
// @Security ApiKeyAuth
// @Summary Get Order By Id
// @Param userId path string true "UserID"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /orders/{userId} [get]
func (i *orderApi) GetOrderByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		i.logger.WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orderResponse, err := i.orderService.GetOrderByUserId(userId)
	if err != nil {
		i.logger.WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	i.logger.WithFields(logrus.Fields{"body": orderResponse}).Info()
	c.JSON(http.StatusOK, gin.H{"data": orderResponse, "message": "success"})
}

// @Tags Order Api
// @Security ApiKeyAuth
// @Summary Complete Order
// @Param id path string true "Order Id"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /orders/{id}/complete [put]
func (i *orderApi) SetStatusOrderCompleted(c *gin.Context) {
	id := c.Param("id")
	err := i.orderService.OrderCompleted(id)
	if err != nil {
		i.logger.WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	i.logger.Info()
	c.JSON(http.StatusOK, "success")
}
