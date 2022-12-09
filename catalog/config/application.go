package config

type AppConfig struct {
	System   SystemConfig   `yaml:"system" mapstructure:"system"`
	Mysql    MysqlConfig    `yaml:"mysql" mapstructure:"mysql"`
	RabbitMq RabbitMqConfig `yaml:"rabbitMq" mapstructure:"rabbitMq"`
}

type QueuesConfig struct {
	Order OrderQueueConfig `yaml:"order" mapstructure:"order"`
}

type OrderQueueConfig struct {
	OrderCompleted QueueConfig `yaml:"orderCompleted" mapstructure:"orderCompleted"`
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
}

type MysqlConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RabbitMqConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
