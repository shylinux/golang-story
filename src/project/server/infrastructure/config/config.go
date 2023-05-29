package config

import (
	"flag"

	"github.com/spf13/viper"
)

type Consul struct {
	Addr     string
	Interval string
}
type Service struct {
	Name string
	Port int
}
type Queue struct {
	Token string
	Host  string
	Port  string
}
type Cache struct {
	Password string
	Host     string
	Port     string
}
type Engine struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}
type Storage struct {
	Engine
	Cache
	Queue
}
type Config struct {
	file string
	Consul
	Service
	Storage
}

var config = &Config{}

func init() {
	flag.StringVar(&config.file, "config.file", "./config/service.yaml", "")
	flag.IntVar(&config.Service.Port, "service.port", 0, "")
}

func New() (*Config, error) {
	flag.Parse()
	defer flag.Parse()
	viper.SetConfigFile(config.file)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	return config, viper.Unmarshal(config)
}
