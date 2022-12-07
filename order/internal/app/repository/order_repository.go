package repository

import (
	"context"

	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/model"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, order model.Order) error
	SetStausOrderCompleted(ctx context.Context, id string) (model.Order, error)
	GetOrderByUserId(ctx context.Context, userId int) (model.Order, error)
}
