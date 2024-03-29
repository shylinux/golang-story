package server

import (
	"fmt"
	"os"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/proxy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/goroutine"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type MainServer struct {
	*config.Config
	consul.Consul
	*proxy.Proxy
	*grpc.Server
	*goroutine.Pool
}

func New(config *config.Config, logger logs.Logger, proxy *proxy.Proxy, consul consul.Consul, server *grpc.Server, pool *goroutine.Pool) *MainServer {
	return &MainServer{config, consul, proxy, server, pool}
}
func (s *MainServer) register(key string, name string, host string, port int) {
	s.Consul.Register(consul.Service{Name: config.WithDef(name, key), Host: host, Port: port})
}
func (s *MainServer) Run() error {
	if s.Config.Logs.Pid != "" {
		system.WriteFile(s.Config.Logs.Pid, []byte(fmt.Sprintf("%d", os.Getpid())), 0755)
		s.Pool.Go("signal", system.Watch)
	}
	// if !s.Config.Consul.Enable {
	// 	return s.Proxy.Run()
	// } else
	if s.Config.Proxy.Export {
		s.Pool.Go("proxy", s.Proxy.Run)
	}
	server := s.Config.Server
	if !s.Config.Proxy.Simple {
		for k, v := range s.Config.Internal {
			if v.Export {
				s.register(k, v.Name, server.Host, server.Port)
			}
		}
	}
	addr := config.Address(server.Host, server.Port)
	logs.Infof("server start %s %s", server.Type, addr)
	system.Printfln("server start %s", addr)
	if l, e := system.Listen("tcp", addr); e != nil {
		return errors.New(e, "server start tcp failure")
	} else {
		return errors.New(s.Server.Serve(l), "server start rpc failure")
	}
}
