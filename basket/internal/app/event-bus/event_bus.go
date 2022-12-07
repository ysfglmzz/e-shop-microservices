package eventbus

import "github.com/ysfglmzz/e-shop-microservices/basket/internal/app/event"

type IEventBus interface {
	Subscribe()
	PublishBasketVerifiedEvent(basketVerifiedEvent event.BasketVerified) error
}
