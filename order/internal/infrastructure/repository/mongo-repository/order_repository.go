package mongorepository

import (
	"context"

	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(collection *mongo.Collection) *mongoOrderRepository {
	return &mongoOrderRepository{collection: collection}
}

func (m *mongoOrderRepository) CreateOrder(ctx context.Context, order model.Order) error {
	_, err := m.collection.InsertOne(ctx, order)
	return err
}

func (m *mongoOrderRepository) SetStausOrderCompleted(ctx context.Context, id string) (model.Order, error) {
	orderIdByteArray, _ := primitive.ObjectIDFromHex(id)
	var order model.Order
	updateResult := m.collection.FindOneAndUpdate(ctx, bson.M{"_id": orderIdByteArray}, bson.M{"$set": bson.M{"status": "completed"}})
	err := updateResult.Decode(&order)
	return order, err
}

func (m *mongoOrderRepository) GetOrderByUserId(ctx context.Context, userId int) (model.Order, error) {
	var order model.Order
	result := m.collection.FindOne(ctx, bson.M{"userId": userId})
	err := result.Decode(&order)
	return order, err
}
