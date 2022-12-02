package service

import (
	"github.com/jinzhu/copier"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/event"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/repository"
)

type productService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) *productService {
	return &productService{productRepository: productRepository}
}

func (p *productService) CreateProduct(createProductDto dto.CreateProductDTO) error {
	var productModel model.Product
	if err := copier.Copy(&productModel, createProductDto); err != nil {
		return err
	}
	return p.productRepository.AddProduct(&productModel)
}

func (p *productService) GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error) {
	return p.productRepository.GetProducts(productFilter)
}

func (p *productService) ReduceProductsQuantities(orderCompletedEvent event.OrderCompletedEvent) error {
	var idList []int

	for _, productInfo := range orderCompletedEvent.Products {
		idList = append(idList, productInfo.Id)
	}

	products, err := p.productRepository.GetProductsByIdList(idList...)
	if err != nil {
		return err
	}

	for _, product := range products {
		for _, productInfo := range orderCompletedEvent.Products {
			if product.Id == productInfo.Id {
				product.Quantity -= productInfo.Quantity
			}
		}
	}

	return p.productRepository.UpdateProducts(products...)
}
