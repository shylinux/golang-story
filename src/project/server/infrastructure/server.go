package infrastructure

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

type MainServer struct {
	*config.Config
	consul.Consul
	*grpc.Server
	*gin.Engine
	Port int
}

func NewMainServer(config *config.Config, consul consul.Consul, server *grpc.Server, engine *gin.Engine) *MainServer {
	return &MainServer{config, consul, server, engine, config.Service.Port}
}
func (s *MainServer) RegisterService(key string, name string, port int) {
	if name == "" {
		name = s.Config.Service.Name + "." + key
	}
	s.Consul.Register(consul.Service{Name: name, Host: s.Config.Service.Host, Port: port})
}
func (s *MainServer) Run() error {
	if s.Config.Service.Main == s.Config.Service.Name {
		for k, v := range s.Config.Internal {
			if v.Export {
				s.RegisterService(k, v.Name, s.Config.Service.Port)
			}
		}
	} else {
		v := s.Config.Internal[s.Config.Service.Main]
		s.RegisterService(s.Config.Service.Main, v.Name, v.Port)
		s.Port = v.Port
	}
	if s.Config.Service.Type == "http" {
		return errors.New(s.Engine.Run(fmt.Sprintf(":%d", s.Port)), "start gin failure")
	} else if l, e := net.Listen("tcp", fmt.Sprintf(":%d", s.Port)); e != nil {
		return errors.New(e, "start rpc failure")
	} else {
		return errors.New(s.Server.Serve(l), "start rpc failure")
	}
}
