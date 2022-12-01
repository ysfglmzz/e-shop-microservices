package config

type AppConfig struct {
	System   SystemConfig   `yaml:"system" mapstructure:"system"`
	Mysql    MysqlConfig    `yaml:"mysql" mapstructure:"mysql"`
	RabbitMq RabbitMqConfig `yaml:"rabbitMq" mapstructure:"yaml"`
}

type SystemConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DbDriver   string `yaml:"dbDriver"`
	DbManager  string `yaml:"dbManager"`
	MessageBus string `yaml:"messageBus"`
}

type MysqlConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RabbitMqConfig struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Exchange   string `yaml:"exchange"`
	RoutingKey string `yaml:"routingKey"`
}