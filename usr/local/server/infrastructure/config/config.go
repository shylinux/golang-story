package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

type Logs struct {
	Pid     string
	Path    string
	MaxSize int
	MaxAge  int
	Stdout  bool
}
type Proxy struct {
	Export bool
	Simple bool
	Target string
	Root   string
	Host   string
	Port   int
}
type Token struct {
	Issuer string
	Secret string
	Expire string
}
type Consul struct {
	Enable   bool
	Addr     string
	Interval string
	WorkID   int
}
type Server struct {
	Type string
	Host string
	Port int
}
type Consumer struct {
	Enable bool
	Name   string
}
type Service struct {
	Export bool
	Name   string
	Host   string
	Port   int
}
type Queue struct {
	Enable bool
	Type   string
	Token  string
	Host   string
	Port   int
}
type Cache struct {
	Enable   bool
	Type     string
	Password string
	Host     string
	Port     int
}
type Search struct {
	Enable   bool
	Type     string
	Username string
	Password string
	Index    string
	Host     string
	Port     int
}
type Storage struct {
	Type     string
	Username string
	Password string
	Database string
	Host     string
	Port     int
}
type Engine struct {
	Storage
	Search
	Cache
	Queue
}
type Config struct {
	file string
	Logs
	Proxy
	Token
	Consul
	Server
	Consumer map[string]Consumer
	Internal map[string]Service
	Engine
	Install
	Replace    []Replace
	ReplaceMap map[string]string
	Generate
	Product
}

var config = &Config{}

func init() {
	flag.StringVar(&config.file, "config.file", "config/service.yaml", "")
	flag.StringVar(&config.Logs.Pid, "logs.pid", "log/service.pid", "")
	flag.StringVar(&config.Logs.Path, "logs.path", "log/service.log", "")
	flag.StringVar(&config.Proxy.Host, "proxy.Host", "", "")
	flag.IntVar(&config.Proxy.Port, "proxy.port", 8081, "")
	flag.StringVar(&config.Token.Issuer, "token.issuer", "auth", "")
	flag.StringVar(&config.Token.Expire, "token.expire", "72h", "")
	flag.BoolVar(&config.Consul.Enable, "consul.enable", false, "")
	flag.StringVar(&config.Consul.Addr, "consul.addr", Address("", 8500), "")
	flag.IntVar(&config.Consul.WorkID, "consul.workid", 1, "")
	flag.IntVar(&config.Server.Port, "server.port", 9090, "")
	flag.StringVar(&config.Engine.Storage.Type, "engine.storage.type", "sqlite", "")
	flag.StringVar(&config.Engine.Storage.Database, "engine.storage.database", "demo.db", "")
}
func New() (*Config, error) {
	flag.Parse()
	defer flag.Parse()
	load("config/install.yaml")
	load("config/replace.yaml")
	load("config/generate.yaml")
	load("config/internal.yaml")
	load("config/product.yaml")
	load(config.file)
	config.ReplaceMap = map[string]string{}
	for _, v := range config.Replace {
		config.ReplaceMap[v.From] = v.To
	}
	consumer := map[string]Consumer{}
	for k, v := range config.Consumer {
		if v.Name == "" {
			consumer[k] = v
		} else {
			consumer[v.Name] = v
		}
	}
	config.Consumer = consumer
	service := map[string]Service{}
	for k, v := range config.Internal {
		if v.Name == "" {
			service[k] = v
		} else {
			service[v.Name] = v
		}
	}
	config.Internal = service
	return config, nil
}
func load(p string) {
	if _, e := os.Stat(p); os.IsNotExist(e) {
		return
	}
	viper.SetConfigFile(p)
	errors.Assert(viper.ReadInConfig())
	errors.Assert(viper.Unmarshal(config))
}
func (config *Config) WithDef(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
func WithDef(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
func Address(host string, port int) string {
	return fmt.Sprintf("%s:%d", WithDef(host, "127.0.0.1"), port)
}
