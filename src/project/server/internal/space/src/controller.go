package space

import (
	"context"
	"fmt"

	"go.uber.org/dig"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

func Init(container *dig.Container) {
	container.Provide(NewSpaceController)
	container.Provide(NewSpaceService)
	container.Provide(NewUserConsumer)
}

type SpaceController struct {
	pb.UnimplementedSpaceServiceServer
	Main    *infrastructure.MainServer
	service *SpaceService
	name    string
}

func NewSpaceController(config *config.Config, mainServer *infrastructure.MainServer, proxy *infrastructure.Proxy, server *grpc.Server, consumer *UserConsumer, service *SpaceService) *SpaceController {
	controller := &SpaceController{Main: mainServer, service: service, name: pb.SpaceService_ServiceDesc.ServiceName}
	if config.Internal[enums.Service.User].Export {
		consul.Tags = append(consul.Tags, pb.SpaceService_ServiceDesc.ServiceName)
		if config.Gateway.Export {
			proxy.Register(controller.name, controller)
		}
		pb.RegisterSpaceServiceServer(server, controller)
	}
	return controller
}
func (s *SpaceController) Create(ctx context.Context, req *pb.SpaceCreateRequest) (*pb.SpaceCreateReply, error) {
	if req.Name == "" || len(req.Name) < 5 {
		return nil, errors.NewInvalidParams(fmt.Errorf(enums.Field.Name))
	}
	space, err := s.service.Create(ctx, req.Name)
	return &pb.SpaceCreateReply{Data: SpaceDTO(space)}, errors.NewCreateFailResp(err)
}
func (s *SpaceController) Remove(ctx context.Context, req *pb.SpaceRemoveRequest) (*pb.SpaceRemoveReply, error) {
	if req.Id < 1 {
		return nil, errors.NewInvalidParams(fmt.Errorf(enums.Field.ID))
	}
	return &pb.SpaceRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.Id))
}
func (s *SpaceController) Info(ctx context.Context, req *pb.SpaceInfoRequest) (*pb.SpaceInfoReply, error) {
	if req.Id < 1 {
		return nil, errors.NewInvalidParams(fmt.Errorf(enums.Field.ID))
	}
	space, err := s.service.Info(ctx, req.Id)
	return &pb.SpaceInfoReply{Data: SpaceDTO(space)}, errors.NewInfoFailResp(err)
}
func (s *SpaceController) List(ctx context.Context, req *pb.SpaceListRequest) (*pb.SpaceListReply, error) {
	if req.Page < 1 || req.Count < 1 {
		return nil, errors.NewInvalidParams(fmt.Errorf("page or count"))
	}
	list, total, err := s.service.List(ctx, req.Page, req.Count)
	data := []*pb.Space{}
	for _, space := range list {
		data = append(data, SpaceDTO(space))
	}
	return &pb.SpaceListReply{Data: data, Total: total}, errors.NewListFailResp(err)
}
