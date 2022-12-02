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
	GetProductsByIdList(idList ...int) ([]*model.Product, error)
	UpdateProducts(products ...*model.Product) error
	// GetProductsByCategoryId(categoryId int) ([]*model.Product, error)
	// GetProductsByBrandId(brandId int) ([]*model.Product, error)
	// GetProductsBySubCategoryId(subCategoryId int) ([]*model.Product, error)
}
