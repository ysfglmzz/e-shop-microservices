package model

type Brand struct {
	Id   int    `gorm:"primarykey"`
	Name string `gorm:"unique"`
}
