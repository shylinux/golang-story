package server

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
	"shylinux.com/x/golang-story/src/project/server/infrastructure/proxy"
)

type MainServer struct {
	*config.Config
	consul.Consul
	*proxy.Proxy
	*grpc.Server
	*gin.Engine
}

func New(config *config.Config, logger logs.Logger, consul consul.Consul, proxy *proxy.Proxy, server *grpc.Server, engine *gin.Engine) *MainServer {
	return &MainServer{config, consul, proxy, server, engine}
}
func (s *MainServer) registerService(key string, name string, host string, port int) {
	name = config.WithDef(name, s.Config.Server.Name+"."+key)
	s.Consul.Register(consul.Service{Name: name, Host: host, Port: port})
}
func (s *MainServer) Run() error {
	server := s.Config.Server
	if k := server.Main; k == server.Name {
		for k, v := range s.Config.Internal {
			if v.Export {
				s.registerService(k, v.Name, server.Host, server.Port)
			}
		}
		if conf := s.Config.Proxy; conf.Export {
			go s.Proxy.Run(conf)
		}
	} else {
		v := s.Config.Internal[k]
		s.registerService(k, v.Name, server.Host, server.Port)
	}
	addr := fmt.Sprintf("%s:%d", server.Host, server.Port)
	logs.Infof("server start %s %s %s", server.Name, server.Type, addr)
	if server.Type == enums.Service.HTTP {
		return errors.New(s.Engine.Run(addr), "server start gin failure")
	} else if l, e := net.Listen("tcp", addr); e != nil {
		return errors.New(e, "server start tcp failure")
	} else {
		return errors.New(s.Server.Serve(l), "server start rpc failure")
	}
}
