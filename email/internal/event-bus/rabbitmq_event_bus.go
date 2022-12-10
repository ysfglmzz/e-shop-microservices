package eventbus

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/email/config"
	"github.com/ysfglmzz/e-shop-microservices/email/internal"
	"github.com/ysfglmzz/e-shop-microservices/email/internal/event"
)

type rabbitMqEventHandler func(delivery amqp.Delivery)

type rabbitMqEventBus struct {
	rabbitMqConfig config.RabbitMqConfig
	systemConfig   config.SystemConfig
	connection     *amqp.Connection
}

func NewRabbitMqEventBus(cfg config.RabbitMqConfig, systemConfig config.SystemConfig, connection *amqp.Connection) *rabbitMqEventBus {
	return &rabbitMqEventBus{rabbitMqConfig: cfg, systemConfig: systemConfig, connection: connection}
}

func (r *rabbitMqEventBus) Subscribe() {
	var forever chan struct{}

	ch, _ := r.connection.Channel()
	r.queueDeclare(ch, r.rabbitMqConfig)
	r.queueBindingDeclare(ch, r.rabbitMqConfig)
	delivery, _ := ch.Consume(r.rabbitMqConfig.Queue, r.rabbitMqConfig.Queue, false, false, false, false, nil)
	go eventHandler(delivery, r.userCreatedEventHandler)

	<-forever
}

func eventHandler(delivery <-chan amqp.Delivery, handler rabbitMqEventHandler) {
	for d := range delivery {
		handler(d)
	}
}

func (r *rabbitMqEventBus) userCreatedEventHandler(delivery amqp.Delivery) {
	var userCreatedEvent event.UserCreatedEvent
	json.Unmarshal(delivery.Body, &userCreatedEvent)

	if err := internal.SendEmail(r.systemConfig, userCreatedEvent); err != nil {
		return
	}
	delivery.Ack(false)

}

func (r *rabbitMqEventBus) queueDeclare(ch *amqp.Channel, rabbitMqConfig config.RabbitMqConfig) *rabbitMqEventBus {
	ch.QueueDeclare(rabbitMqConfig.Queue, true, false, false, false, nil)
	return r
}

func (r *rabbitMqEventBus) queueBindingDeclare(ch *amqp.Channel, rabbitMqConfig config.RabbitMqConfig) *rabbitMqEventBus {
	ch.QueueBind(rabbitMqConfig.Queue, rabbitMqConfig.RoutingKey, rabbitMqConfig.Exchange, false, nil)
	return r
}
