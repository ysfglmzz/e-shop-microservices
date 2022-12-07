package mongorepository

import (
	"context"

	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoBasketRepository struct {
	collection *mongo.Collection
}

func NewMongoBasketRepository(collection *mongo.Collection) *mongoBasketRepository {
	return &mongoBasketRepository{collection: collection}
}

func (m *mongoBasketRepository) CreateBasket(ctx context.Context, basket model.Basket) error {
	_, err := m.collection.InsertOne(ctx, basket)
	return err
}

func (m *mongoBasketRepository) AddProductToBasket(ctx context.Context, basketId string, product model.Product) error {
	basketIdByteArray, _ := primitive.ObjectIDFromHex(basketId)
	addProductBson := bson.M{
		"$push": bson.M{"products": product},
		"$inc":  bson.M{"itemCount": product.Quantity},
	}
	_, err := m.collection.UpdateByID(ctx, basketIdByteArray, addProductBson)
	return err
}

func (m *mongoBasketRepository) GetBasketByUserId(ctx context.Context, userId int) (*model.Basket, error) {
	var basket *model.Basket
	result := m.collection.FindOne(ctx, bson.M{"userId": userId})
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&basket); err != nil {
		return nil, err
	}
	return basket, nil
}

func (m *mongoBasketRepository) EmptyBasket(ctx context.Context, userId int) error {
	filter := bson.M{"userId": userId}
	emptyBasketBson := bson.M{
		"$set": bson.M{
			"products":  []*model.Product{},
			"itemCount": 0,
		},
	}
	_, err := m.collection.UpdateOne(ctx, filter, emptyBasketBson)
	return err
}
