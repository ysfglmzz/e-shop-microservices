package gormrepository

import (
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/model"
	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *gormProductRepository {
	return &gormProductRepository{db: db}
}

func (g *gormProductRepository) AddProduct(product *model.Product) error {
	return g.db.Create(&product).Error
}

func (g *gormProductRepository) UpdateProduct(product *model.Product) error {
	return g.db.Save(&product).Error
}

func (g *gormProductRepository) DeleteProductById(id int) error {
	return g.db.Delete(&model.Product{}, "id", id).Error
}

func (g *gormProductRepository) GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error) {
	var productResponseList []*dto.ProductResponse

	query := g.db.Debug().Model(&model.Product{})
	query = query.Offset(*productFilter.Offset)
	query = query.Limit(*productFilter.Limit)

	if productFilter.BrandId != nil {
		query = query.Joins("INNER JOIN brands AS b ON b.id = products.brand_id AND b.id = ?", productFilter.BrandId)
	} else {
		query = query.Joins("INNER JOIN brands AS b ON b.id = products.brand_id")
	}

	if productFilter.CategoryId != nil {
		query = query.Joins("INNER JOIN categories AS c ON c.id = products.category_id AND c.id = ?", productFilter.CategoryId)
	} else {
		query = query.Joins("INNER JOIN categories AS c ON c.id = products.category_id")
	}

	if productFilter.SubCategoryId != nil {
		query = query.Joins("INNER JOIN sub_categories AS sc ON sc.id = products.sub_category_id AND sc.id = ?", productFilter.SubCategoryId)
	} else {
		query = query.Joins("INNER JOIN sub_categories AS sc ON sc.id = products.sub_category_id")
	}

	err := query.Select("products.*",
		"c.id as c_id", "c.name as c_name",
		"sc.id as sc_id", "sc.name as sc_name",
		"b.id as b_id", "b.name as b_name").Scan(&productResponseList).Error

	if err != nil {
		return nil, err
	}

	return productResponseList, nil
}

func (g *gormProductRepository) GetProductsByIdList(idList ...int) ([]*model.Product, error) {
	var products []*model.Product

	if err := g.db.Where(idList).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (g *gormProductRepository) UpdateProducts(products ...*model.Product) error {
	return g.db.Save(&products).Error
}
