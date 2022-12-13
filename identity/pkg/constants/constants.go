package constants

const (
	ConfigPath               = "./resource"
	ConfigType               = "yaml"
	ConfigName               = "application"
	MysqlConnectionFormat    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	RabbitMqConnectionFormat = "amqp://%s:%s@%s:%d/"
)
