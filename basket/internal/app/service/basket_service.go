package service

import (
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/model"
)

type IBasketService interface {
	GetBasketByUserId(userId int) (*dto.GetBasketByUserIdResponse, error)
	AddProductToBasket(addProductRequest dto.AddProductToBasketRequest) error
	CreateBasket(basket model.Basket) error
	VerifyBasket(userId int) error
}
