package model

import "time"

type User struct {
	Id           int `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	FirstName    string
	LastName     string
	Username     string
	Email        string
	PhoneNumber  string
	PasswordHash string
	EmailVerify  bool
	UserRoles    []*UserRole
}
