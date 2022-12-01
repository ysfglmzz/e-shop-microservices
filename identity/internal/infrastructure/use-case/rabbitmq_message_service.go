package usecase

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/event"
)

type RabbitMqMessageService struct {
	connection *amqp.Connection
}

func NewRabbitMqMessageService() *RabbitMqMessageService {
	mesgService := &RabbitMqMessageService{}
	mesgService.connect().exchangeDeclare()
	return mesgService
}

func (r *RabbitMqMessageService) connect() *RabbitMqMessageService {
	r.connection, _ = amqp.Dial("amqp://guest:guest@localhost:5672")
	return r
}

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
	err := ch.PublishWithContext(context.Background(), "UserCreated", "UserCreated", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})
	return err
}

func (r *RabbitMqMessageService) exchangeDeclare() *RabbitMqMessageService {
	ch, _ := r.connection.Channel()
	defer ch.Close()
	ch.ExchangeDeclare("UserCreated", amqp.ExchangeTopic, true, false, false, false, nil)
	return r
}
