package service

import (
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/event"
)

type IProductService interface {
	CreateProduct(createProductDto dto.CreateProductDTO) error
	GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error)
	ReduceProductsQuantities(orderCompletedEvent event.OrderCompletedEvent) error
}
