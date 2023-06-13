package controller

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/router"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
	name    string
}

func NewUserController(config *config.Config, server *server.MainServer, service *service.UserService) *UserController {
	controller := &UserController{service: service, name: pb.UserService_ServiceDesc.ServiceName}
	if !config.Internal[enums.Service.User].Export {
		return controller
	}
	server.Proxy.Register(controller.name, controller)
	consul.Tags = append(consul.Tags, controller.name)
	router.Register(server.Engine, controller.name, controller)
	pb.RegisterUserServiceServer(server.Server, controller)
	return controller
}
func (s *UserController) Create(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateReply, error) {
	user, err := s.service.Create(ctx, req.Username, "", req.Email, "")
	return &pb.UserCreateReply{Data: trans.UserDTO(user)}, errors.NewCreateFailResp(err)
}
func (s *UserController) Remove(ctx context.Context, req *pb.UserRemoveRequest) (*pb.UserRemoveReply, error) {
	return &pb.UserRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.UserID))
}
func (s *UserController) Rename(ctx context.Context, req *pb.UserRenameRequest) (*pb.UserRenameReply, error) {
	return &pb.UserRenameReply{}, errors.NewModifyFailResp(s.service.Rename(ctx, req.UserID, req.Username))
}
func (s *UserController) Search(ctx context.Context, req *pb.UserSearchRequest) (*pb.UserSearchReply, error) {
	list, total, err := s.service.Search(ctx, req.Page, req.Count, req.Key, req.Value)
	data := []*pb.User{}
	trans.ListDTO(list, trans.UserDTO, &data)
	return &pb.UserSearchReply{Data: data, Total: total}, errors.NewSearchFailResp(err)
}
func (s *UserController) Info(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	user, err := s.service.Info(ctx, req.UserID)
	return &pb.UserInfoReply{Data: trans.UserDTO(user)}, errors.NewInfoFailResp(err)
}
func (s *UserController) List(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
	list, total, err := s.service.List(ctx, req.Page, req.Count, req.Key, req.Value)
	data := []*pb.User{}
	trans.ListDTO(list, trans.UserDTO, &data)
	return &pb.UserListReply{Data: data, Total: total}, errors.NewListFailResp(err)
}
