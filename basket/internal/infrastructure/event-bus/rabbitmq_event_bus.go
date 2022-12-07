package eventbus

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ysfglmzz/e-shop-microservices/basket/config"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/event"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/repository"
)

type rabbitMqEventHandler func(delivery amqp.Delivery)

type rabbitMqEventBus struct {
	cfg              config.QueuesConfig
	connection       *amqp.Connection
	basketRepository repository.IBasketRepository
	queueHandlerMap  map[config.QueueConfig]rabbitMqEventHandler
}

func NewRabbitMqEventBus(
	cfg config.QueuesConfig,
	connection *amqp.Connection,
	basketRepository repository.IBasketRepository,
) *rabbitMqEventBus {
	return &rabbitMqEventBus{
		cfg:              cfg,
		connection:       connection,
		basketRepository: basketRepository,
		queueHandlerMap:  make(map[config.QueueConfig]rabbitMqEventHandler),
	}
}

func (r *rabbitMqEventBus) Subscribe() {
	r.setHandlers().
		basketQueueDeclareAndBind().
		generateConsumers()
}

func (r *rabbitMqEventBus) PublishBasketVerifiedEvent(basketVerifiedEvent event.BasketVerified) error {
	queueConfig := r.cfg.Basket.BasketVerified
	body, _ := json.Marshal(&basketVerifiedEvent)
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

func (r *rabbitMqEventBus) userCreatedEvent(delivery amqp.Delivery) {
	var userCreatedEvent event.UserCreated
	json.Unmarshal(delivery.Body, &userCreatedEvent)
	if err := r.basketRepository.CreateBasket(context.Background(), model.Basket{
		Id:        primitive.NewObjectID(),
		UserId:    userCreatedEvent.UserId,
		ItemCount: 0,
		Products:  []*model.Product{},
	}); err != nil {
		return
	}
	delivery.Ack(false)
}

func (r *rabbitMqEventBus) orderCompletedEvent(delivery amqp.Delivery) {
	var orderCompletedEvent event.OrderCompleted
	json.Unmarshal(delivery.Body, &orderCompletedEvent)
	if err := r.basketRepository.EmptyBasket(context.Background(), orderCompletedEvent.UserId); err != nil {
		return
	}
	delivery.Ack(false)
}

func (r *rabbitMqEventBus) setHandlers() *rabbitMqEventBus {
	r.queueHandlerMap[r.cfg.User.UserCreated] = r.userCreatedEvent
	r.queueHandlerMap[r.cfg.Order.OrderCompleted] = r.orderCompletedEvent
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
