package config

type AppConfig struct {
	System   SystemConfig   `yaml:"system" mapstructure:"system"`
	RabbitMq RabbitMqConfig `yaml:"rabbitMq" mapstructure:"rabbitMq"`
}

type SystemConfig struct {
	EmailHost         string `yaml:"emailHost"`
	EmailPort         int    `yaml:"emailPort"`
	FromEmail         string `yaml:"fromEmail"`
	FromEmailPassword string `yaml:"fromEmailPassword"`
	MessageBus        string `yaml:"messageBus"`
}

type MysqlConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RabbitMqConfig struct {
	Port       int    `yaml:"port"`
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Exchange   string `yaml:"exchange"`
	RoutingKey string `yaml:"routingKey"`
	Queue      string `yaml:"queue"`
}
