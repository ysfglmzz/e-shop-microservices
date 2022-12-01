package dto

import "time"

type TokenResponse struct {
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expirationDate"`
}
