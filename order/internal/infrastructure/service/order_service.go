package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/event"
	eventbus "github.com/ysfglmzz/e-shop-microservices/order/internal/app/event-bus"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/repository"
)

type orderService struct {
	orderRepository repository.IOrderRepository
	eventBus        eventbus.IEventBus
}

func NewOrderService(orderRepository repository.IOrderRepository, eventBus eventbus.IEventBus) *orderService {
	return &orderService{orderRepository: orderRepository, eventBus: eventBus}
}

func (o *orderService) OrderCompleted(id string) error {
	order, err := o.orderRepository.SetStausOrderCompleted(context.Background(), id)
	if err != nil {
		return err
	}
	orderCompletedEvent := event.OrderCompleted{UserId: order.UserId}
	for _, product := range order.Products {
		orderCompletedEvent.Products = append(orderCompletedEvent.Products, &event.ProductInfo{
			Id:       product.Id,
			Quantity: product.Quantity,
		})
	}
	return o.eventBus.PublishOrderCompletedEvent(orderCompletedEvent)
}

func (o *orderService) GetOrderByUserId(userId int) (*dto.GetOrderByUserIdResponse, error) {
	var orderResponse dto.GetOrderByUserIdResponse
	order, err := o.orderRepository.GetOrderByUserId(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	if err = copier.Copy(&orderResponse, order); err != nil {
		return nil, err
	}
	return &orderResponse, nil
}
