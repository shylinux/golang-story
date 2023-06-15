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
func (s *MainServer) registerService(key string, name string, host string, port int) {
	name = config.WithDef(name, s.Config.Server.Name+"."+key)
	s.Consul.Register(consul.Service{Name: name, Host: host, Port: port})
}
func (s *MainServer) Run() error {
	server := s.Config.Server
	if k := server.Main; k == server.Name || k == "server" {
		for k, v := range s.Config.Internal {
			if v.Export {
				s.registerService(k, v.Name, server.Host, server.Port)
			}
		}
		if s.Config.Proxy.Export {
			s.Pool.Go("proxy", s.Proxy.Run)
		}
	} else {
		v := s.Config.Internal[k]
		s.registerService(k, v.Name, server.Host, server.Port)
	}
	if s.Config.Logs.Pid != "" {
		system.WriteFile(s.Config.Logs.Pid, []byte(fmt.Sprintf("%d", os.Getpid())), 0755)
		s.Pool.Go("signal", system.Watch)
	}
	addr := config.Address(server.Host, server.Port)
	logs.Infof("server start %s %s %s", server.Name, server.Type, addr)
	system.Printfln("server start %s", addr)
	if l, e := system.Listen("tcp", addr); e != nil {
		return errors.New(e, "server start tcp failure")
	} else {
		return errors.New(s.Server.Serve(l), "server start rpc failure")
	}
}
