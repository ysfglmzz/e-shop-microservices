package factories

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ysfglmzz/e-shop-microservices/identity/config"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/infrastructure/db/mysql"
	"github.com/ysfglmzz/e-shop-microservices/identity/pkg/constants"

	"gorm.io/gorm"
)

type ConnectionFactory struct {
	cfg                config.AppConfig
	rabbitMqConnection *amqp.Connection
	gormDb             *gorm.DB
}

func NewConnectionFactory(cfg config.AppConfig) *ConnectionFactory {
	return &ConnectionFactory{cfg: cfg}
}

func (c *ConnectionFactory) GetGormDb() *gorm.DB {
	return c.gormDb
}

func (c *ConnectionFactory) GetRabbitMqConnection() *amqp.Connection {
	return c.rabbitMqConnection
}

func (c *ConnectionFactory) DbConnect() *ConnectionFactory {
	switch c.cfg.System.DbManager {
	case "gorm":
		c.connectGorm()
	}
	return c
}
func (c *ConnectionFactory) MessageBusConnect() *ConnectionFactory {
	switch c.cfg.System.MessageBus {
	case "rabbitMq":
		c.connectRabbitMq()
	}
	return c
}

func (c *ConnectionFactory) connectRabbitMq() {
	rabbitMqConfig := c.cfg.RabbitMq
	connectionString := fmt.Sprintf(constants.RabbitMqConnectionFormat, rabbitMqConfig.User, rabbitMqConfig.Password, rabbitMqConfig.Host, rabbitMqConfig.Port)
	c.rabbitMqConnection, _ = amqp.Dial(connectionString)

}

func (c *ConnectionFactory) connectGorm() {
	switch c.cfg.System.DbDriver {
	case "mysql":
		c.gormDb = mysql.GormConnect(c.cfg.Mysql)
	}
	c.gormDb.AutoMigrate(
		model.User{},
		model.Role{},
		model.UserRole{},
		model.TokenDetail{},
	)
	if c.cfg.System.InitDb {
		c.gormDb.Create([]model.Role{
			{Id: 1, Name: "admin"},
			{Id: 2, Name: "customer"},
		})
	}
}
