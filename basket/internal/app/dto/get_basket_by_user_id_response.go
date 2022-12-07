package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetBasketByUserIdResponse struct {
	Id        primitive.ObjectID `json:"id"`
	UserId    int                `json:"userId"`
	Products  []*Product         `json:"products"`
	ItemCount int                `json:"itemCount"`
}
