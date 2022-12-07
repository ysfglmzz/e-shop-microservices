package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      int                `bson:"userId"`
	Date        time.Time          `bson:"date"`
	Status      string             `bson:"status"`
	TotalAmount int                `bson:"totalAmount"`
	Products    []*Product         `bson:"products"`
}

type Product struct {
	Id        int    `bson:"id"`
	Name      string `bson:"name"`
	Quantity  int    `bson:"quantity"`
	UnitPrice int    `bson:"unitPrice"`
}
