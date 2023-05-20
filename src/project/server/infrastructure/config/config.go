package config

import (
	"flag"

	"github.com/spf13/viper"
)

type Service struct {
	Addr string
}
type Storage struct {
	Cache
	Engine
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
type Config struct {
	file string
	Service
	Storage
}

var config = &Config{}

func init() {
	flag.StringVar(&config.file, "config.file", "./config/service.yaml", "")
	flag.StringVar(&config.Service.Addr, "service.addr", "", "host:port")
}

func NewConfig() (*Config, error) {
	flag.Parse()
	defer flag.Parse()
	viper.SetConfigFile(config.file)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	return config, viper.Unmarshal(config)
}
