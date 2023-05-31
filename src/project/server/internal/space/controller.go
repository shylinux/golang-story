package space

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	user "shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceController struct {
	service *SpaceService
	pb.UnimplementedSpaceServiceServer
	user user.UserServiceClient
}

func NewSpaceController(config *config.Config, csl consul.Consul, server *grpc.Server, consumer *UserConsumer, service *SpaceService) *SpaceController {
	controller := &SpaceController{service: service}
	if pb.RegisterSpaceServiceServer(server, controller); config.Internal[enums.Service.Space].Export {
		csl.Register(consul.Service{Name: strings.Join([]string{config.Service.Name, enums.Service.Space}, "."), Port: config.Service.Port})
	}
	return controller
}
func (s *SpaceController) Create(ctx context.Context, req *pb.SpaceCreateRequest) (*pb.SpaceCreateReply, error) {
	space, err := s.service.Create(ctx, req.Name)
	return &pb.SpaceCreateReply{Data: SpaceDTO(space)}, errors.NewCreateFailResp(err)
}
func (s *SpaceController) Remove(ctx context.Context, req *pb.SpaceRemoveRequest) (*pb.SpaceRemoveReply, error) {
	return &pb.SpaceRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.Id))
}
func (s *SpaceController) Info(ctx context.Context, req *pb.SpaceInfoRequest) (*pb.SpaceInfoReply, error) {
	space, err := s.service.Info(ctx, req.Id)
	return &pb.SpaceInfoReply{Data: SpaceDTO(space)}, errors.NewInfoFailResp(err)
}
func (s *SpaceController) List(ctx context.Context, req *pb.SpaceListRequest) (*pb.SpaceListReply, error) {
	list, err := s.service.List(ctx, req.Page, req.Count)
	data := []*pb.Space{}
	for _, space := range list {
		data = append(data, SpaceDTO(space))
	}
	return &pb.SpaceListReply{Data: data}, errors.NewListFailResp(err)
}
