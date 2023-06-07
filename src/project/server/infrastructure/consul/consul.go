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
	*config.Config
	*api.Client
}

func New(config *config.Config) (Consul, error) {
	conf := api.DefaultConfig()
	conf.Address = config.Consul.Addr
	client, err := api.NewClient(conf)
	logs.Infof("find service consul %s", conf.Address)
	return &consul{config, client}, err
}

var Tags = []string{}
var Meta = map[string]string{}

func (s *consul) Register(service Service) error {
	interval := config.ValueWithDef(s.Config.Consul.Interval, "10s")
	registration := new(api.AgentServiceRegistration)
	registration.Tags = Tags
	registration.Meta = Meta
	registration.Name = service.Name
	registration.Port = service.Port
	registration.Address = service.Host
	registration.ID = fmt.Sprintf("%s-%s-%d", service.Name, service.Host, service.Port)
	registration.Check = &api.AgentServiceCheck{
		Interval: interval, DeregisterCriticalServiceAfter: interval,
		GRPC: fmt.Sprintf("%s/%s", service.Address(), registration.Name),
	}
	logs.Infof("register service %+v %s", service, logs.FileLine(2))
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
	logs.Infof("resolve service %s %+v %s", name, res, logs.FileLine(2))
	return
}
func (s *consul) Address(target string) string {
	list := strings.Split(target, ".")
	if len(list) > 1 {
		list = list[:len(list)-1]
	}
	return fmt.Sprintf("consul://%s/%s", s.Config.Consul.Addr, strings.Join(list, "."))
}
