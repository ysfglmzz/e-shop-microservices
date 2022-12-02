package factories

import (
	"github.com/ysfglmzz/e-shop-microservices/catalog/config"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/infrastructure/db/mysql"

	"gorm.io/gorm"
)

type ConnectionFactory struct {
	cfg    config.AppConfig
	gormDb *gorm.DB
}

func NewConnectionFactory(cfg config.AppConfig) *ConnectionFactory {
	return &ConnectionFactory{cfg: cfg}
}

func (c *ConnectionFactory) GetGormDb() *gorm.DB {
	return c.gormDb
}

func (c *ConnectionFactory) DbConnect() *ConnectionFactory {
	switch c.cfg.System.DbManager {
	case "gorm":
		c.connectGorm()
	}
	return c
}

func (c *ConnectionFactory) connectGorm() {
	switch c.cfg.System.DbDriver {
	case "mysql":
		c.gormDb = mysql.GormConnect(c.cfg.Mysql)
	}
	c.gormDb.AutoMigrate(
		model.Product{},
		model.Category{},
		model.SubCategory{},
		model.Brand{},
	)
	if c.cfg.System.InitDb {
		c.gormDb.Create([]model.Category{
			{Id: 1, Name: "Electronic", SubCategories: []*model.SubCategory{
				{Id: 1, Name: "Computer"},
				{Id: 2, Name: "Phone"},
				{Id: 3, Name: "Television"},
				{Id: 4, Name: "Headphone"},
			}},
			{Id: 2, Name: "Home", SubCategories: []*model.SubCategory{
				{Id: 5, Name: "Carpet"},
				{Id: 6, Name: "Mirror"},
				{Id: 7, Name: "Duvet Cover"},
				{Id: 8, Name: "Sofa"},
			}},
			{Id: 3, Name: "Sport", SubCategories: []*model.SubCategory{
				{Id: 9, Name: "Track suit"},
				{Id: 10, Name: "Spikes"},
				{Id: 11, Name: "Dumbell set"},
				{Id: 12, Name: "Fitness Equipment"},
			}},
			{Id: 4, Name: "Clothing", SubCategories: []*model.SubCategory{
				{Id: 13, Name: "Dress"},
				{Id: 14, Name: "Shoe"},
				{Id: 15, Name: "Suit"},
				{Id: 16, Name: "Skirt"},
			}},
		})
		c.gormDb.Create([]model.Brand{
			{Id: 1, Name: "HP"},
			{Id: 2, Name: "Samsung"},
			{Id: 3, Name: "Apple"},
			{Id: 4, Name: "Vestel"},
			{Id: 5, Name: "Xiaomi"},
			{Id: 6, Name: "Beko"},
			{Id: 7, Name: "Arcelik"},
			{Id: 8, Name: "Nike"},
			{Id: 9, Name: "Adidas"},
			{Id: 10, Name: "Mavi"},
			{Id: 11, Name: "Kigili"},
		})
		c.gormDb.Create([]model.Product{
			{Id: 1, Name: "Hp Pavilion Notebook", BrandId: 1, CategoryId: 1, SubCategoryId: 1, Quantity: 120, UnitPrice: 12000},
			{Id: 2, Name: "iPhone 11", BrandId: 3, CategoryId: 1, SubCategoryId: 2, Quantity: 350, UnitPrice: 20000},
			{Id: 3, Name: "Galaxy A32", BrandId: 2, CategoryId: 1, SubCategoryId: 2, Quantity: 451, UnitPrice: 12000},
			{Id: 4, Name: "Mi True Headphone", BrandId: 5, CategoryId: 1, SubCategoryId: 4, Quantity: 200, UnitPrice: 400},
		})
	}
}
