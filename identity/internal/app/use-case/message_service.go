package usecase

import (
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/event"
)

type IMessageService interface {
	PublishUserCreatedEvent(event event.UserCreated) error
}
