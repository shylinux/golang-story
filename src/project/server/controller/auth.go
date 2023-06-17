package controller

import (
	"context"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/metadata"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type AuthController struct {
	pb.UnimplementedAuthServiceServer
	service *service.AuthService
	name    string
}

func NewAuthController(config *config.Config, server *server.MainServer, service *service.AuthService) *AuthController {
	controller := &AuthController{service: service, name: pb.AuthService_ServiceDesc.ServiceName}
	if !config.Internal[enums.Service.Auth].Export {
		return controller
	}
	server.Proxy.Auth = controller.auth
	server.Proxy.Register(controller.name, controller)
	server.Server.Register(&pb.AuthService_ServiceDesc, controller)
	consul.Tags = append(consul.Tags, controller.name)
	return controller
}
func (s *AuthController) auth(ctx context.Context, api string, token string) (context.Context, error) {
	if strings.Contains(api, s.name) {
		return ctx, nil
	} else if info, err := s.service.Verify(ctx, token); err == nil {
		return metadata.SetValue(ctx, metadata.USERNAME, info.Username), nil
	} else {
		return ctx, err
	}
}
func (s *AuthController) Register(ctx context.Context, req *pb.AuthRegisterRequest) (*pb.AuthRegisterReply, error) {
	token, err := s.service.Register(ctx, req.Username, req.Password, req.Email, req.Phone)
	return &pb.AuthRegisterReply{Token: token}, errors.NewCreateFailResp(err)
}
func (s *AuthController) Login(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginReply, error) {
	token, err := s.service.Login(ctx, req.Username, req.Password)
	return &pb.AuthLoginReply{Token: token}, errors.NewNotAuth(err)
}
func (s *AuthController) Logout(ctx context.Context, req *pb.AuthLogoutRequest) (*pb.AuthLogoutReply, error) {
	return &pb.AuthLogoutReply{}, errors.NewRemoveFailResp(nil)
}
func (s *AuthController) Verify(ctx context.Context, req *pb.AuthVerifyRequest) (*pb.AuthVerifyReply, error) {
	info, err := s.service.Verify(ctx, req.Token)
	return &pb.AuthVerifyReply{Username: info.Username}, errors.NewInfoFailResp(err)
}
