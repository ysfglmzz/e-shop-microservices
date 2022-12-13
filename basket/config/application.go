package config

type AppConfig struct {
	System   SystemConfig   `yaml:"system" mapstructure:"system"`
	Mongo    MongoConfig    `yaml:"mongo" mapstructure:"mongo"`
	RabbitMq RabbitMqConfig `yaml:"rabbitMq" mapstructure:"rabbitMq"`
	Consul   ConsulConfig   `yaml:"consul" mapstructure:"consul"`
}

type QueuesConfig struct {
	Order  OrderQueueConfig  `yaml:"order" mapstructure:"order"`
	User   UserQueueConfig   `yaml:"user" mapstructure:"user"`
	Basket BasketQueueConfig `yaml:"basket" mapstructure:"basket"`
}

type OrderQueueConfig struct {
	OrderCompleted QueueConfig `yaml:"orderCompleted" mapstructure:"orderCompleted"`
}

type UserQueueConfig struct {
	UserCreated QueueConfig `yaml:"userCreated" mapstructure:"userCreated"`
}

type BasketQueueConfig struct {
	BasketVerified QueueConfig `yaml:"basketVerified" mapstructure:"basketVerified"`
}

type QueueConfig struct {
	Exchange     string `yaml:"exchange"`
	ExchangeType string `yaml:"exchangeType"`
	RoutingKey   string `yaml:"routingKey"`
	Queue        string `yaml:"queue"`
}

type SystemConfig struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	DbDriver         string `yaml:"dbDriver"`
	DbManager        string `yaml:"dbManager"`
	Server           string `yaml:"server"`
	InitDb           bool   `yaml:"initDb"`
	MessageBus       string `yaml:"messageBus"`
	ServiceDiscovery string `yaml:"serviceDiscovery"`
	TokenSecretKey   string `yaml:"tokenSecretKey"`
}

type MongoConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

type RabbitMqConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ConsulConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}
