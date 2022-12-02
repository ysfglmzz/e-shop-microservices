package config

type AppConfig struct {
	System SystemConfig `yaml:"system" mapstructure:"system"`
	Mysql  MysqlConfig  `yaml:"mysql" mapstructure:"mysql"`
}

type SystemConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DbDriver  string `yaml:"dbDriver"`
	DbManager string `yaml:"dbManager"`
	Server    string `yaml:"server"`
	InitDb    bool   `yaml:"initDb"`
}

type MysqlConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
