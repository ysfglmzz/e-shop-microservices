package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetOrderByUserIdResponse struct {
	Id          primitive.ObjectID `json:"id"`
	UserId      int                `json:"userId"`
	Date        time.Time          `json:"date"`
	Status      string             `json:"status"`
	TotalAmount int                `json:"totalAmount"`
	Products    []*Product         `json:"products"`
}

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice string `json:"unitPrice"`
}
