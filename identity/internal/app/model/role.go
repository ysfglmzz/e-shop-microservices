package model

type Role struct {
	Id   int `gorm:"primarykey"`
	Name string
}
