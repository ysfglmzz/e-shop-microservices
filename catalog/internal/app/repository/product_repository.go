package repository

import (
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/model"
)

type IProductRepository interface {
	AddProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProductById(id int) error
	GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error)
	// GetProductById(id int) (*model.Product, error)
	// GetProductsByCategoryId(categoryId int) ([]*model.Product, error)
	// GetProductsByBrandId(brandId int) ([]*model.Product, error)
	// GetProductsBySubCategoryId(subCategoryId int) ([]*model.Product, error)
}
