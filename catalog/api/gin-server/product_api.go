package ginserver

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/dto"
	service "github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/service"
)

type ProductApi struct {
	logger         *logrus.Logger
	productService service.IProductService
}

func NewProductApi(identityService service.IProductService, logger *logrus.Logger) *ProductApi {
	return &ProductApi{productService: identityService, logger: logger}
}

// @Tags Product Api
// @Security ApiKeyAuth
// @Summary Crate Product
// @Param data body dto.CreateProductDTO true "CreateProductDTO"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /products [post]
func (i *ProductApi) CreateProduct(c *gin.Context) {
	var createProductDto dto.CreateProductDTO
	if err := c.ShouldBindJSON(&createProductDto); err != nil {
		i.logger.WithFields(logrus.Fields{"body": c.Request.Body}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := i.productService.CreateProduct(createProductDto); err != nil {
		i.logger.WithFields(logrus.Fields{"body": createProductDto}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	i.logger.WithFields(logrus.Fields{"body": createProductDto}).Info()
	c.JSON(http.StatusCreated, "Created successfully")
}

// @Tags Product Api
// @Security ApiKeyAuth
// @Summary Get Product
// @Param data query dto.ProductFilter true "ProductFilter"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /products [get]
func (i *ProductApi) GetProducts(c *gin.Context) {
	var productFilter dto.ProductFilter
	if err := c.ShouldBindQuery(&productFilter); err != nil {
		i.logger.WithFields(logrus.Fields{"body": c.Request.Body}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	productResponseList, err := i.productService.GetProducts(productFilter)
	if err != nil {
		i.logger.WithFields(logrus.Fields{"body": productFilter}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	i.logger.WithFields(logrus.Fields{"body": productResponseList}).Info()
	c.JSON(http.StatusOK, map[string]any{"data": productResponseList, "message": "Get products successfuly"})
}
