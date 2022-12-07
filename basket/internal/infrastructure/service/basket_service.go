package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/event"
	eventbus "github.com/ysfglmzz/e-shop-microservices/basket/internal/app/event-bus"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/repository"
)

type basketService struct {
	eventBus         eventbus.IEventBus
	basketRepository repository.IBasketRepository
}

func NewBasketService(basketRepository repository.IBasketRepository, eventBus eventbus.IEventBus) *basketService {
	return &basketService{basketRepository: basketRepository, eventBus: eventBus}
}

func (b *basketService) GetBasketByUserId(userId int) (*dto.GetBasketByUserIdResponse, error) {
	basket, err := b.basketRepository.GetBasketByUserId(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	var basketResponse dto.GetBasketByUserIdResponse
	if err = copier.Copy(&basketResponse, basket); err != nil {
		return nil, err
	}
	return &basketResponse, nil
}

func (b *basketService) AddProductToBasket(addProductRequest dto.AddProductToBasketRequest) error {
	return b.basketRepository.AddProductToBasket(
		context.Background(),
		addProductRequest.BaketId,
		model.Product(*addProductRequest.Product),
	)
}

func (b *basketService) CreateBasket(basket model.Basket) error {
	return b.basketRepository.CreateBasket(context.Background(), basket)
}

func (b *basketService) VerifyBasket(userId int) error {
	basket, err := b.GetBasketByUserId(userId)
	if err != nil {
		return err
	}
	var basketVerifiedEvent event.BasketVerified
	if err := copier.Copy(&basketVerifiedEvent, basket); err != nil {
		return err
	}
	return b.eventBus.PublishBasketVerifiedEvent(basketVerifiedEvent)
}
