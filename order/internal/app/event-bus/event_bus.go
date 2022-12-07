package eventbus

import "github.com/ysfglmzz/e-shop-microservices/order/internal/app/event"

type IEventBus interface {
	Subscribe()
	PublishOrderCompletedEvent(orderCompleted event.OrderCompleted) error
}
