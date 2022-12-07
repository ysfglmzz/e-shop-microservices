package ginserver

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/service"
)

type basketApi struct {
	logger        *logrus.Logger
	basketService service.IBasketService
}

func NewBasketApi(basketService service.IBasketService, logger *logrus.Logger) *basketApi {
	return &basketApi{basketService: basketService, logger: logger}
}

// @Tags BasketApi
// @Security ApiKeyAuth
// @Summary Get Basket
// @Param userId path string true "User Id"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /baskets/{userId} [get]
func (i *basketApi) GetBasketByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		i.logger.WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	basket, err := i.basketService.GetBasketByUserId(userId)
	if err != nil {
		i.logger.WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	i.logger.WithFields(logrus.Fields{"body": basket}).Info()
	c.JSON(http.StatusCreated, gin.H{"data": basket, "message": "success"})
}

// @Tags BasketApi
// @Security ApiKeyAuth
// @Summary Add Product To Basket
// @Param data body dto.AddProductToBasketRequest true "AddProductToBasketRequest"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /baskets/addProduct [post]
func (i *basketApi) AddProductToBasket(c *gin.Context) {
	var addProductToBasketRequest dto.AddProductToBasketRequest
	if err := c.ShouldBindJSON(&addProductToBasketRequest); err != nil {
		i.logger.WithFields(logrus.Fields{"body": c.Request.Body}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := i.basketService.AddProductToBasket(addProductToBasketRequest)
	if err != nil {
		i.logger.WithFields(logrus.Fields{"body": addProductToBasketRequest}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	i.logger.WithFields(logrus.Fields{"body": addProductToBasketRequest}).Info()
	c.JSON(http.StatusOK, "Login successfully")
}
