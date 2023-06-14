package space

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

func Init(container *container.Container) {
	container.Provide(NewSpaceController)
	container.Provide(NewSpaceService)
	container.Provide(NewUserConsumer)
}

type SpaceController struct {
	pb.UnimplementedSpaceServiceServer
	Main    *server.MainServer
	service *SpaceService
	name    string
}

func NewSpaceController(config *config.Config, server *server.MainServer, consumer *UserConsumer, service *SpaceService) *SpaceController {
	controller := &SpaceController{Main: server, service: service, name: pb.SpaceService_ServiceDesc.ServiceName}
	if !config.Internal[enums.Service.User].Export {
		return controller
	}
	server.Proxy.Register(controller.name, controller)
	server.Server.Register(&pb.SpaceService_ServiceDesc, controller)
	consul.Tags = append(consul.Tags, controller.name)
	return controller
}
func (s *SpaceController) Create(ctx context.Context, req *pb.SpaceCreateRequest) (*pb.SpaceCreateReply, error) {
	space, err := s.service.Create(ctx, req.Name)
	return &pb.SpaceCreateReply{Data: SpaceDTO(space)}, errors.NewCreateFailResp(err)
}
func (s *SpaceController) Remove(ctx context.Context, req *pb.SpaceRemoveRequest) (*pb.SpaceRemoveReply, error) {
	return &pb.SpaceRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.SpaceID))
}
func (s *SpaceController) Info(ctx context.Context, req *pb.SpaceInfoRequest) (*pb.SpaceInfoReply, error) {
	space, err := s.service.Info(ctx, req.SpaceID)
	return &pb.SpaceInfoReply{Data: SpaceDTO(space)}, errors.NewInfoFailResp(err)
}
func (s *SpaceController) List(ctx context.Context, req *pb.SpaceListRequest) (*pb.SpaceListReply, error) {
	list, total, err := s.service.List(ctx, req.Page, req.Count, req.Key, req.Value)
	data := []*pb.Space{}
	trans.ListDTO(list, SpaceDTO, &data)
	return &pb.SpaceListReply{Data: data, Total: total}, errors.NewListFailResp(err)
}
