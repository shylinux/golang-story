package config

import (
	"flag"
	"os"
	"path"

	"github.com/spf13/viper"
)

type Log struct {
	Path    string
	MaxSize int
	MaxAge  int
	Stdout  bool
}
type Consul struct {
	Addr     string
	Interval string
}
type Service struct {
	Export bool
	Name   string
	Main   string
	Type   string
	Host   string
	Port   int
}
type Queue struct {
	Name  string
	Token string
	Host  string
	Port  int
}
type Cache struct {
	Name     string
	Password string
	Host     string
	Port     int
}
type Storage struct {
	Name     string
	Username string
	Password string
	Database string
	Host     string
	Port     int
}
type Engine struct {
	Storage
	Cache
	Queue
}
type Config struct {
	file string
	Log
	Consul
	Service
	Internal map[string]Service
	Engine
}

var config = &Config{}

func init() {
	flag.StringVar(&config.file, "config.file", "./config/service.yaml", "")
	flag.StringVar(&config.Log.Path, "log.path", "./log/service.log", "")
	flag.StringVar(&config.Consul.Addr, "consul.addr", "127.0.0.1:8500", "")
	flag.StringVar(&config.Service.Name, "service.name", path.Base(os.Args[0]), "")
	flag.StringVar(&config.Service.Main, "service.main", path.Base(os.Args[0]), "")
	flag.StringVar(&config.Service.Host, "service.host", "127.0.0.1", "")
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
	return config, nil
}
