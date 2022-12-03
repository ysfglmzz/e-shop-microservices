package factory

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/email/config"
	eventbus "github.com/ysfglmzz/e-shop-microservices/email/internal/event-bus"
	"github.com/ysfglmzz/e-shop-microservices/email/pkg/constants"
)

type eventBusFactory struct {
	cfg config.AppConfig
}

func NewEventBusFactory(cfg config.AppConfig) *eventBusFactory {
	return &eventBusFactory{cfg: cfg}
}

func (e *eventBusFactory) GetEventBus() eventbus.IEventBus {
	switch e.cfg.System.MessageBus {
	case "rabbitMq":
		return eventbus.NewRabbitMqEventBus(e.cfg.RabbitMq, e.cfg.System, e.rabbitMqConnect())
	}
	return nil
}

func (e *eventBusFactory) rabbitMqConnect() *amqp.Connection {
	rabbitMqConfig := e.cfg.RabbitMq
	connectionString := fmt.Sprintf(constants.RabbitMqConnectionFormat, rabbitMqConfig.User, rabbitMqConfig.Password, rabbitMqConfig.Host, rabbitMqConfig.Port)
	conn, _ := amqp.Dial(connectionString)
	return conn
}
