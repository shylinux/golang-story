package controller

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
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

func (s *MainController) Run(addr string) error {
	// return s.Engine.Run(addr)
	if l, e := net.Listen("tcp", addr); e != nil {
		return e
	} else {
		return s.Server.Serve(l)
	}
}

func NewMainController(logger log.Logger, server *grpc.Server, engine *gin.Engine, user *UserController) *MainController {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(logger, logger, logger, 99))
	pb.RegisterUserServiceServer(server, user)
	router.Register(engine, "user", user)
	engine.GET("/metrics", func(ctx *gin.Context) { promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request) })
	return &MainController{server, engine}
}
