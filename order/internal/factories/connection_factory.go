package factories

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/order/config"
	"github.com/ysfglmzz/e-shop-microservices/order/pkg/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionFactory struct {
	cfg                config.AppConfig
	rabbitMqConnection *amqp.Connection
	mongoDb            *mongo.Database
}

func NewConnectionFactory(cfg config.AppConfig) *ConnectionFactory {
	return &ConnectionFactory{cfg: cfg}
}

func (c *ConnectionFactory) GetMongoDb() *mongo.Database {
	return c.mongoDb
}

func (c *ConnectionFactory) GetRabbitMqConnection() *amqp.Connection {
	return c.rabbitMqConnection
}

func (c *ConnectionFactory) EventBusConnect() *ConnectionFactory {
	switch c.cfg.System.MessageBus {
	case "rabbitMq":
		c.connectRabbitMq()
	}
	return c
}

func (c *ConnectionFactory) ConnectDb() *ConnectionFactory {
	switch c.cfg.System.DbDriver {
	case "mongo":
		c.mongoDbConnect()
	}
	return c
}

func (c *ConnectionFactory) connectRabbitMq() {
	rabbitMqConfig := c.cfg.RabbitMq
	connectionString := fmt.Sprintf(constants.RabbitMqConnectionFormat, rabbitMqConfig.User, rabbitMqConfig.Password, rabbitMqConfig.Host, rabbitMqConfig.Port)
	c.rabbitMqConnection, _ = amqp.Dial(connectionString)
}

func (c *ConnectionFactory) mongoDbConnect() {
	connectionString := fmt.Sprintf(constants.MongoDbConnectionFormat, c.cfg.Mongo.Host, c.cfg.Mongo.Port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err.Error())
	}
	c.mongoDb = client.Database(c.cfg.Mongo.Database)
}
