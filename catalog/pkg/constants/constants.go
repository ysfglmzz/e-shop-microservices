package constants

const (
	ConfigPath               = "./resource"
	ConfigType               = "yaml"
	AppConfigName            = "application"
	QueueConfigName          = "queue"
	MysqlConnectionFormat    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	RabbitMqConnectionFormat = "amqp://%s:%s@%s:%d/"
)
