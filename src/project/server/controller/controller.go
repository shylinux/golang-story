package controller

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/router"
)

func Init(container *dig.Container) {
	container.Provide(NewMainController)
	container.Provide(grpc.NewServer)
	container.Provide(gin.New)
	container.Provide(NewUserController)
}

type MainController struct {
	*grpc.Server
	*gin.Engine
}

func (s *MainController) Run(port int) error {
	// return s.Engine.Run(addr)
	if l, e := net.Listen("tcp", fmt.Sprintf(":%d", port)); e != nil {
		return e
	} else {
		return s.Server.Serve(l)
	}
}

func NewMainController(logger log.Logger, consul consul.Consul, server *grpc.Server, engine *gin.Engine, user *UserController) *MainController {
	// grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(logger, logger, logger, 99))
	grpc_health_v1.RegisterHealthServer(server, &HealthController{})
	pb.RegisterUserServiceServer(server, user)
	router.Register(engine, "user", user)
	engine.GET("/metrics", func(ctx *gin.Context) { promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request) })
	return &MainController{server, engine}
}
