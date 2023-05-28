package consul

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

type Consul interface {
}

type consul struct {
}

func New(config *config.Config) (Consul, error) {
	conf := consulapi.DefaultConfig()
	conf.Address = config.Consul.Addr
	client, err := consulapi.NewClient(conf)
	if err != nil {
		return nil, err
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = fmt.Sprintf("%s-%d", config.Service.Name, config.Service.Port)
	registration.Address = fmt.Sprintf("%s:%d", "127.0.0.1", config.Service.Port)
	registration.Name = config.Service.Name
	registration.Port = config.Service.Port
	registration.Check = &consulapi.AgentServiceCheck{
		Interval: config.Consul.Interval, DeregisterCriticalServiceAfter: config.Consul.Interval,
		GRPC: fmt.Sprintf("%s/%s", registration.Address, registration.Name),
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		return nil, err
	}
	return &consul{}, nil
}
