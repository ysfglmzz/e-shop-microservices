package eventbus

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jinzhu/copier"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ysfglmzz/e-shop-microservices/order/config"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/event"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/order/internal/app/repository"
)

type rabbitMqEventHandler func(delivery amqp.Delivery)

type rabbitMqEventBus struct {
	cfg             config.QueuesConfig
	connection      *amqp.Connection
	orderRepository repository.IOrderRepository
	queueHandlerMap map[config.QueueConfig]rabbitMqEventHandler
}

func NewRabbitMqEventBus(
	cfg config.QueuesConfig,
	connection *amqp.Connection,
	orderRepository repository.IOrderRepository,
) *rabbitMqEventBus {
	return &rabbitMqEventBus{
		cfg:             cfg,
		connection:      connection,
		orderRepository: orderRepository,
		queueHandlerMap: make(map[config.QueueConfig]rabbitMqEventHandler),
	}
}

func (r *rabbitMqEventBus) Subscribe() {
	r.setHandlers().
		basketQueueDeclareAndBind().
		generateConsumers()
}

func (r *rabbitMqEventBus) PublishOrderCompletedEvent(orderCompletedEvent event.OrderCompleted) error {
	queueConfig := r.cfg.Order.OrderCompleted
	body, _ := json.Marshal(&orderCompletedEvent)
	ch, _ := r.connection.Channel()
	defer ch.Close()

	err := ch.PublishWithContext(context.Background(), queueConfig.Exchange, queueConfig.RoutingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	return err
}

func (r *rabbitMqEventBus) basketQueueDeclareAndBind() *rabbitMqEventBus {
	ch, _ := r.connection.Channel()
	defer ch.Close()
	return r.queueDeclare(ch, r.cfg.Basket.BasketVerified).
		queueBindingDeclare(ch, r.cfg.Basket.BasketVerified)
}

func (r *rabbitMqEventBus) BasketVerifiedEventHandler(delivery amqp.Delivery) {
	var basketVerifiedEvent event.BasketVerified
	if err := json.Unmarshal(delivery.Body, &basketVerifiedEvent); err != nil {
		return
	}
	var order model.Order
	if err := copier.Copy(&order, basketVerifiedEvent); err != nil {
		return
	}
	totalAmount := 0
	for _, product := range basketVerifiedEvent.Products {
		totalAmount += product.UnitPrice * product.Quantity
	}
	order.Id = primitive.NewObjectID()
	order.Date = time.Now()
	order.Status = "waiting"
	order.TotalAmount = totalAmount
	if err := r.orderRepository.CreateOrder(context.Background(), order); err != nil {
		return
	}
	delivery.Ack(false)
}

func (r *rabbitMqEventBus) setHandlers() *rabbitMqEventBus {
	r.queueHandlerMap[r.cfg.Basket.BasketVerified] = r.BasketVerifiedEventHandler
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

func (r *rabbitMqEventBus) exchangeDeclare(ch *amqp.Channel, queueConfig config.QueueConfig) *rabbitMqEventBus {
	ch.ExchangeDeclare(queueConfig.Exchange, queueConfig.ExchangeType, true, false, false, false, nil)
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
