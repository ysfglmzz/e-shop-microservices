package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenDetail struct {
	Id             int `gorm:"primarykey"`
	UUID           uuid.UUID
	ExpirationDate time.Time
	Token          string
	UserId         int
	User           User
}
