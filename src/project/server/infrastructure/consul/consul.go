package consul

import (
	"fmt"
	"strings"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type Service struct {
	Name string
	Host string
	Port int
}

func (s Service) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Consul interface {
	Register(service Service) error
	Resolve(name string) ([]Service, error)
	Address(target string) string
}
type consul struct {
	*api.Client
	address  string
	interval string
	*config.Config
}

var Meta = map[string]string{}
var Tags = []string{}

func New(config *config.Config, logs logs.Logger) (Consul, error) {
	conf := api.DefaultConfig()
	conf.Address = config.Consul.Addr
	client, err := api.NewClient(conf)
	return &consul{client, config.Consul.Addr, config.Consul.Interval, config}, err
}

func (s *consul) Register(service Service) error {
	if service.Host == "" {
		service.Host = "127.0.0.1"
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = service.Name
	registration.Tags = Tags
	registration.Meta = Meta
	registration.Port = service.Port
	registration.Address = service.Host
	registration.ID = fmt.Sprintf("%s-%s-%d", service.Name, service.Host, service.Port)
	registration.Check = &api.AgentServiceCheck{
		Interval: s.interval, DeregisterCriticalServiceAfter: s.interval,
		GRPC: fmt.Sprintf("%s/%s", service.Address(), registration.Name),
	}
	logs.With().Infof("register service %+v", service)
	return s.Client.Agent().ServiceRegister(registration)
}
func (s *consul) Resolve(name string) (res []Service, err error) {
	list, err := s.Client.Agent().Services()
	if err != nil {
		return
	}
	for _, v := range list {
		if v.Service == name {
			res = append(res, Service{Name: v.ID, Host: v.Address, Port: v.Port})
		}
	}
	logs.With().Infof("resolve service %s %+v", name, res)
	return
}
func (s *consul) Address(target string) string {
	list := strings.Split(target, ".")
	if len(list) > 1 {
		list = list[:len(list)-1]
	}
	return fmt.Sprintf("consul://%s/%s", s.address, strings.Join(list, "."))
}
