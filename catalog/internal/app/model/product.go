package model

type Product struct {
	Id            int    `gorm:"primarykey"`
	Name          string `gorm:"unique"`
	BrandId       int
	Brand         Brand
	CategoryId    int
	Category      Category
	SubCategoryId int
	SubCategory   SubCategory
	Quantity      int
	UnitPrice     int
}
