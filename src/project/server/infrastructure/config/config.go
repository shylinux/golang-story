package config

import (
	"flag"
	"os"
	"path"

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
	Local  bool
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
	Addr     string
	Interval string
	WorkID   int
}
type Server struct {
	Name string
	Type string
	Main string
	Host string
	Port int
}
type Service struct {
	Export bool
	Name   string
	Host   string
	Port   int
}
type Cache struct {
	Name     string
	Password string
	Host     string
	Port     int
}
type Queue struct {
	Name  string
	Token string
	Host  string
	Port  int
}
type Search struct {
	Name     string
	Username string
	Password string
	Index    string
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
	Search
	Queue
	Cache
}
type Matrix struct {
	Generate
	Install
}
type Config struct {
	file string
	Logs
	Proxy
	Token
	Consul
	Server
	Internal map[string]Service
	Engine
	Matrix
}

var config = &Config{}

func init() {
	flag.StringVar(&config.file, "config.file", "config/service.yaml", "")
	flag.StringVar(&config.Logs.Pid, "logs.pid", "log/service.pid", "")
	flag.StringVar(&config.Logs.Path, "logs.path", "log/service.log", "")
	flag.StringVar(&config.Token.Issuer, "token.issuer", "demo.auth", "")
	flag.StringVar(&config.Token.Expire, "token.expire", "24h", "")
	flag.StringVar(&config.Consul.Addr, "consul.addr", "127.0.0.1:8500", "")
	flag.IntVar(&config.Consul.WorkID, "consul.workid", 1, "")
	flag.StringVar(&config.Server.Name, "server.name", path.Base(os.Args[0]), "")
	flag.StringVar(&config.Server.Main, "server.main", path.Base(os.Args[0]), "")
	flag.StringVar(&config.Server.Host, "server.host", "127.0.0.1", "")
	flag.IntVar(&config.Server.Port, "server.port", 9090, "")
}
func New() (*Config, error) {
	flag.Parse()
	defer flag.Parse()
	load(config.file)
	load("config/install.yaml")
	load("config/generate.yaml")
	if config.Server.Main != config.Server.Name {
		if v := config.Internal[config.Server.Main]; v.Port > 0 {
			config.Server.Port = v.Port
		}
	}
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
