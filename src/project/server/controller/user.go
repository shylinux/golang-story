package controller

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/router"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
	name    string
}

func NewUserController(config *config.Config, mainServer *infrastructure.MainServer, proxy *infrastructure.Proxy, engine *gin.Engine, server *grpc.Server, service *service.UserService) *UserController {
	controller := &UserController{service: service, name: pb.UserService_ServiceDesc.ServiceName}
	if config.Internal[enums.Service.User].Export {
		consul.Tags = append(consul.Tags, controller.name)
		if config.Gateway.Export {
			proxy.Register(controller.name, controller)
		}
		if config.Service.Type == enums.Service.HTTP {
			router.Register(engine, controller.name, controller)
		} else {
			pb.RegisterUserServiceServer(server, controller)
		}
	}
	return controller
}
func (s *UserController) Create(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateReply, error) {
	if req.Name == "" || len(req.Name) < 5 {
		return nil, errors.NewInvalidParams(fmt.Errorf(enums.Field.Name))
	}
	user, err := s.service.Create(ctx, req.Name)
	return &pb.UserCreateReply{Data: trans.UserDTO(user)}, errors.NewCreateFailResp(err)
}
func (s *UserController) Remove(ctx context.Context, req *pb.UserRemoveRequest) (*pb.UserRemoveReply, error) {
	if req.Id < 1 {
		return nil, errors.NewInvalidParams(fmt.Errorf(enums.Field.ID))
	}
	return &pb.UserRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.Id))
}
func (s *UserController) Info(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	if req.Id < 1 {
		return nil, errors.NewInvalidParams(fmt.Errorf(enums.Field.ID))
	}
	user, err := s.service.Info(ctx, req.Id)
	return &pb.UserInfoReply{Data: trans.UserDTO(user)}, errors.NewInfoFailResp(err)
}
func (s *UserController) List(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
	if req.Page < 1 || req.Count < 1 {
		return nil, errors.NewInvalidParams(fmt.Errorf("page or count"))
	}
	list, err := s.service.List(ctx, req.Page, req.Count)
	data := []*pb.User{}
	for _, user := range list {
		data = append(data, trans.UserDTO(user))
	}
	return &pb.UserListReply{Data: data}, errors.NewListFailResp(err)
}
