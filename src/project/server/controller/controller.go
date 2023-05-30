package controller

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/router"
	"shylinux.com/x/golang-story/src/project/server/internal"
)

func Init(container *dig.Container) {
	container.Provide(NewMainController)
	container.Provide(NewUserController)
}

type MainController struct {
	*config.Config
	*grpc.Server
	*gin.Engine
}

func NewMainController(config *config.Config, csl consul.Consul, server *grpc.Server, engine *gin.Engine, user *UserController, internal *internal.InternalController) *MainController {
	pb.RegisterUserServiceServer(server, user)
	router.Register(engine, enums.Service.User, user)
	csl.Register(consul.Service{Name: config.Service.Name, Port: config.Service.Port})
	return &MainController{config, server, engine}
}
func (s *MainController) Run() error {
	if s.Config.Service.Type == "http" {
		return s.Engine.Run(fmt.Sprintf(":%d", s.Config.Service.Port))
	} else if l, e := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Service.Port)); e != nil {
		return e
	} else {
		return s.Server.Serve(l)
	}
}
