package controller

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type UserController struct {
	service *service.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}
func (s *UserController) Create(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateReply, error) {
	user, err := s.service.Create(ctx, req.Name)
	return &pb.UserCreateReply{Data: trans.UserDTO(user)}, errors.NewCreateFailResp(err)
}
func (s *UserController) Remove(ctx context.Context, req *pb.UserRemoveRequest) (*pb.UserRemoveReply, error) {
	return &pb.UserRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.Id))
}
func (s *UserController) Info(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	user, err := s.service.Info(ctx, req.Id)
	return &pb.UserInfoReply{Data: trans.UserDTO(user)}, errors.NewInfoFailResp(err)
}
func (s *UserController) List(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
	list, err := s.service.List(ctx, req.Page, req.Count)
	data := []*pb.User{}
	for _, user := range list {
		data = append(data, trans.UserDTO(user))
	}
	return &pb.UserListReply{Data: data}, errors.NewListFailResp(err)
}
