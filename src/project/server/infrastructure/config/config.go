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
	Target string
	Export bool
	Type   string
	Name   string
	Host   string
	Port   int
}
type Queue struct {
	Token string
	Host  string
	Port  int
}
type Cache struct {
	Password string
	Host     string
	Port     int
}
type Engine struct {
	Username string
	Password string
	Database string
	Host     string
	Port     int
}
type Storage struct {
	Engine
	Cache
	Queue
}
type Config struct {
	file    string
	LogPath string
	Consul
	Service
	Internal map[string]Service
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
	if err := viper.Unmarshal(config); err != nil {
		return config, err
	}
	if config.Service.Host == "" {
		config.Service.Host = "127.0.0.1"
	}
	return config, nil
}
