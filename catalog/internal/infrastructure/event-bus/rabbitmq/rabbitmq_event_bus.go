package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/catalog/config"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/event"
	"github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/service"
)

type rabbitMqEventHandler func(delivery amqp.Delivery)

type rabbitMqEventBus struct {
	cfg             config.QueuesConfig
	connection      *amqp.Connection
	productService  service.IProductService
	queueHandlerMap map[config.QueueConfig]rabbitMqEventHandler
}

func NewRabbitMqEventBus(
	cfg config.QueuesConfig,
	connection *amqp.Connection,
	productService service.IProductService,
) *rabbitMqEventBus {
	return &rabbitMqEventBus{
		cfg:             cfg,
		connection:      connection,
		productService:  productService,
		queueHandlerMap: make(map[config.QueueConfig]rabbitMqEventHandler),
	}
}

func (r *rabbitMqEventBus) Subscribe() {
	r.setHandlers().
		generateConsumers()
}

func (r *rabbitMqEventBus) orderCompletedEventHandler(delivery amqp.Delivery) {
	var orderCompletedEvent event.OrderCompletedEvent
	json.Unmarshal(delivery.Body, &orderCompletedEvent)
	if err := r.productService.ReduceProductsQuantities(orderCompletedEvent); err != nil {
		return
	}
	delivery.Ack(false)
}

func (r *rabbitMqEventBus) setHandlers() *rabbitMqEventBus {
	r.queueHandlerMap[r.cfg.Order.OrderCompleted] = r.orderCompletedEventHandler
	return r
}

func (r *rabbitMqEventBus) exchangeDeclare(ch *amqp.Channel, queueConfig config.QueueConfig) *rabbitMqEventBus {
	ch.ExchangeDeclare(queueConfig.Exchange, queueConfig.ExchangeType, true, false, false, false, nil)
	return r
}

func (r *rabbitMqEventBus) queueDeclare(ch *amqp.Channel, queueConfig config.QueueConfig) *rabbitMqEventBus {
	ch.QueueDeclare(queueConfig.Queue, true, false, false, false, nil)
	return r
}

func (r *rabbitMqEventBus) queueBindingDeclare(ch *amqp.Channel, queueConfig config.QueueConfig) *rabbitMqEventBus {
	ch.QueueBind(queueConfig.Queue, queueConfig.RoutingKey, queueConfig.Exchange, false, nil)
	return r
}

func (r *rabbitMqEventBus) generateConsumers() {
	for queueConfig, handler := range r.queueHandlerMap {
		ch, _ := r.connection.Channel()
		r.exchangeDeclare(ch, queueConfig)
		r.queueDeclare(ch, queueConfig)
		r.queueBindingDeclare(ch, queueConfig)
		delivery, _ := ch.Consume(queueConfig.Queue, queueConfig.Queue, false, false, false, false, nil)
		go eventHandler(delivery, handler)
	}
}

func eventHandler(delivery <-chan amqp.Delivery, handler rabbitMqEventHandler) {
	for d := range delivery {
		handler(d)
	}
}
