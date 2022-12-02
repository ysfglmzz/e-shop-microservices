package model

type Category struct {
	Id            int    `gorm:"primarykey"`
	Name          string `gorm:"unique"`
	SubCategories []*SubCategory
}

type SubCategory struct {
	Id         int `gorm:"primarykey"`
	CategoryId int
	Category   Category
	Name       string `gorm:"unique"`
}
