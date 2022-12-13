package config

type AppConfig struct {
	System   SystemConfig   `yaml:"system" mapstructure:"system"`
	Mysql    MysqlConfig    `yaml:"mysql" mapstructure:"mysql"`
	RabbitMq RabbitMqConfig `yaml:"rabbitMq" mapstructure:"rabbitMq"`
	Consul   ConsulConfig   `yaml:"consul" mapstructure:"consul"`
}

type SystemConfig struct {
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	DbDriver            string `yaml:"dbDriver"`
	DbManager           string `yaml:"dbManager"`
	MessageBus          string `yaml:"messageBus"`
	Server              string `yaml:"server"`
	InitDb              bool   `yaml:"initDb"`
	TokenExpirationTime int    `yaml:"tokenExpirationTime"`
	TokenSecretKey      string `yaml:"tokenSecretKey"`
	ServiceDiscovery    string `yaml:"serviceDiscovery"`
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

type ConsulConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}
