package config

type RoutesConfig struct {
	Server              string    `yaml:"server"`
	Host                string    `yaml:"host"`
	Port                int       `yaml:"port"`
	ServiceDiscovery    string    `yaml:"serviceDiscovery"`
	UseServiceDiscovery bool      `yaml:"useServiceDiscovery"`
	TokenSecretKey      string    `yaml:"tokenSecretKey"`
	Redis               Redis     `yaml:"redis" mapstructure:"redis"`
	Consul              Consul    `yaml:"consul" mapstructure:"consul"`
	Services            []Service `yaml:"services" mapstructure:"services"`
}

type Service struct {
	Address string  `yaml:"address"`
	Name    string  `yaml:"name"`
	Routes  []Route `yaml:"routes" mapstructure:"routes"`
}

type Route struct {
	Path       string `yaml:"path"`
	Method     string `yaml:"method"`
	Middleware bool   `yaml:"middleware"`
}

type Consul struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
