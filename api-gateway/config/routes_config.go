package config

type RoutesConfig struct {
	Server   string    `yaml:"server"`
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	Services []Service `yaml:"services" mapstructure:"services"`
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
