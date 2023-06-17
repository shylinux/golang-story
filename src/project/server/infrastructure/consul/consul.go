package consul

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
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

func New(config *config.Config, _ logs.Logger) (Consul, error) {
	if !config.Consul.Enable {
		return &consul{config, nil}, nil
	}
	conf := api.DefaultConfig()
	conf.Address = config.Consul.Addr
	if client, err := api.NewClient(conf); err != nil {
		logs.Errorf("engine connect consul %s %s", conf.Address, err)
		return nil, errors.New(err, "engine connnect consul")
	} else {
		logs.Infof("engine connect consul %s", conf.Address)
		return &consul{config, client}, nil
	}
}

var Tags = []string{}
var Meta = map[string]string{}

func init() {
	wd, _ := os.Getwd()
	Meta["bin"] = os.Args[0]
	Meta["dir"] = wd
}

func (s *consul) Register(service Service) error {
	if !s.Enable {
		return nil
	}
	registration := new(api.AgentServiceRegistration)
	registration.Tags = Tags
	registration.Meta = Meta
	registration.Name = service.Name
	registration.Port = service.Port
	registration.Address = service.Host
	registration.ID = fmt.Sprintf("%s-%s-%d", service.Name, service.Host, service.Port)
	if interval := s.Config.Consul.Interval; interval != "" {
		registration.Check = &api.AgentServiceCheck{
			Interval: interval, DeregisterCriticalServiceAfter: interval,
			GRPC: fmt.Sprintf("%s/%s", service.Address(), registration.Name),
		}
	}
	if err := s.Client.Agent().ServiceRegister(registration); err != nil {
		logs.Errorf("consul register %s %+v %s %s", service.Name, service, errors.FileLine(2), err)
		return errors.New(err, "consul register")
	} else {
		logs.Infof("consul register %s %+v %s", service.Name, service, errors.FileLine(2))
		return nil
	}
}
func (s *consul) Resolve(name string) (res []Service, err error) {
	if !s.Enable {
		return nil, errors.New(fmt.Errorf("consul not found"), name)
	}
	list, err := s.Client.Agent().Services()
	if err != nil {
		logs.Errorf("consul resolve %s %s %s", name, err, errors.FileLine(2))
		return nil, errors.New(err, "consul resolve")
	}
	for _, v := range list {
		if v.Service == name {
			res = append(res, Service{Name: v.ID, Host: v.Address, Port: v.Port})
		}
	}
	logs.Infof("consul resolve %s %+v %s", name, res, errors.FileLine(2))
	return
}
func (s *consul) Address(target string) string {
	list := strings.Split(target, ".")
	if len(list) > 1 {
		list = list[:len(list)-1]
	}
	name := strings.Join(list, ".")
	if !s.Enable {
		service := s.Config.Internal[strings.TrimPrefix(name, s.Config.Server.Name+".")]
		return config.Address(service.Host, service.Port)
	}
	return fmt.Sprintf("consul://%s/%s", s.Config.Consul.Addr, name)
}
