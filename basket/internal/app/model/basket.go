package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Basket struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    int                `bson:"userId"`
	Products  []*Product         `bson:"products"`
	ItemCount int                `bson:"itemCount"`
}

type Product struct {
	Id        int    `bson:"id"`
	Name      string `bson:"name"`
	Quantity  int    `bson:"quantity"`
	UnitPrice int    `bson:"unitPrice"`
}
