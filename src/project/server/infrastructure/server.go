package infrastructure

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type MainServer struct {
	*config.Config
	consul.Consul
	*grpc.Server
	*gin.Engine
}

func NewMainServer(config *config.Config, logger logs.Logger, consul consul.Consul, server *grpc.Server, engine *gin.Engine) *MainServer {
	return &MainServer{config, consul, server, engine}
}
func (s *MainServer) Run() error {
	service := s.Config.Service
	if k := service.Main; k == service.Name {
		for k, v := range s.Config.Internal {
			if v.Export {
				s.RegisterService(k, v.Name, service.Host, service.Port)
			}
		}
	} else {
		v := s.Config.Internal[k]
		if v.Port > 0 {
			service.Port = v.Port
		}
		s.RegisterService(k, v.Name, service.Host, service.Port)
	}
	logs.Infof("start service %s %s %s:%d", service.Name, service.Type, service.Host, service.Port)
	if service.Type == enums.Service.HTTP {
		return errors.New(s.Engine.Run(fmt.Sprintf(":%d", service.Port)), "start gin failure")
	} else if l, e := net.Listen("tcp", fmt.Sprintf(":%d", service.Port)); e != nil {
		return errors.New(e, "start rpc failure")
	} else {
		return errors.New(s.Server.Serve(l), "start rpc failure")
	}
}
func (s *MainServer) RegisterService(key string, name string, host string, port int) {
	if name == "" {
		name = s.Config.Service.Name + "." + key
	}
	s.Consul.Register(consul.Service{Name: name, Host: host, Port: port})
}
