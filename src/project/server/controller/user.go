package controller

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
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
	return &pb.UserCreateReply{Data: UserDTO(user)}, errors.NewResp(err, 1, "user create failure")
}
func (s *UserController) Remove(ctx context.Context, req *pb.UserRemoveRequest) (*pb.UserRemoveReply, error) {
	return &pb.UserRemoveReply{}, errors.NewResp(s.service.Remove(ctx, req.Id), 2, "user remove failure")
}
func (s *UserController) Info(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	user, err := s.service.Info(ctx, req.Id)
	return &pb.UserInfoReply{Data: UserDTO(user)}, errors.NewResp(err, 3, "user info failure")
}
func (s *UserController) List(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
	list, err := s.service.List(ctx, req.Page, req.Count)
	data := []*pb.User{}
	for _, user := range list {
		data = append(data, UserDTO(user))
	}
	return &pb.UserListReply{Data: data}, errors.NewResp(err, 4, "user list failure")
}
func UserDTO(user *model.User) *pb.User {
	if user == nil {
		return nil
	}
	return &pb.User{Id: user.ID, Name: user.Name, Email: user.Email}
}
