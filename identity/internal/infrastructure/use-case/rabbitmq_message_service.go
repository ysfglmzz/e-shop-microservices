package usecase

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/identity/config"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/event"
)

type RabbitMqMessageService struct {
	cfg        config.RabbitMqConfig
	connection *amqp.Connection
}

func NewRabbitMqMessageService(cfg config.RabbitMqConfig, connection *amqp.Connection) *RabbitMqMessageService {
	mesgService := &RabbitMqMessageService{cfg: cfg, connection: connection}
	mesgService.exchangeDeclare()
	return mesgService
}

// func (r *RabbitMqMessageService) connect() *RabbitMqMessageService {
// 	connectionString := fmt.Sprintf(constants.RabbitMqConnectionFormat, r.cfg.User, r.cfg.Password, r.cfg.Host, r.cfg.Port)
// 	r.connection, _ = amqp.Dial(connectionString)
// 	return r
// }

// func (r *RabbitMqMessageService) queueDeclare() *RabbitMqMessageService {
// 	ch, _ := r.connection.Channel()
// 	defer ch.Close()

// 	ch.QueueDeclare("UserCreated", true, false, false, false, nil)
// 	return r
// }

func (r *RabbitMqMessageService) PublishUserCreatedEvent(event event.UserCreated) error {
	ch, _ := r.connection.Channel()
	defer ch.Close()
	data, _ := json.Marshal(&event)
	err := ch.PublishWithContext(context.Background(), r.cfg.Exchange, r.cfg.RoutingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})
	return err
}

func (r *RabbitMqMessageService) exchangeDeclare() *RabbitMqMessageService {
	ch, _ := r.connection.Channel()
	defer ch.Close()
	ch.ExchangeDeclare(r.cfg.Exchange, amqp.ExchangeTopic, true, false, false, false, nil)
	return r
}
