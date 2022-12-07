package factories

import (
	"github.com/ysfglmzz/e-shop-microservices/basket/config"
	eventbus "github.com/ysfglmzz/e-shop-microservices/basket/internal/app/event-bus"
	rabbitMqEventBus "github.com/ysfglmzz/e-shop-microservices/basket/internal/infrastructure/event-bus"
)

type EventBusFactory struct {
	systemConfig      config.SystemConfig
	queuesConfig      config.QueuesConfig
	connectionFactory ConnectionFactory
	repositoryFactory RepositoryFactory
}

func NewEventBusFactory(
	systemConfig config.SystemConfig,
	queuesConfig config.QueuesConfig,
	connectionFactory ConnectionFactory,
	repositoryFactory RepositoryFactory,
) *EventBusFactory {
	return &EventBusFactory{
		systemConfig:      systemConfig,
		queuesConfig:      queuesConfig,
		connectionFactory: connectionFactory,
		repositoryFactory: repositoryFactory,
	}
}

func (e *EventBusFactory) GetEventBus() eventbus.IEventBus {
	switch e.systemConfig.MessageBus {
	case "rabbitMq":
		return rabbitMqEventBus.NewRabbitMqEventBus(e.queuesConfig, e.connectionFactory.GetRabbitMqConnection(), e.repositoryFactory.GetBasketRepository())
	}
	return nil
}
