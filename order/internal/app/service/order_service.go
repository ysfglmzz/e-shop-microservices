package service

import "github.com/ysfglmzz/e-shop-microservices/order/internal/app/dto"

type IOrderService interface {
	OrderCompleted(id string) error
	GetOrderByUserId(userId int) (*dto.GetOrderByUserIdResponse, error)
}
