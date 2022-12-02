package service

import "github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/dto"

type IProductService interface {
	CreateProduct(createProductDto dto.CreateProductDTO) error
	GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error)
}
