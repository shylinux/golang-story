package space

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	user "shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	grpcs "shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceController struct {
	service *SpaceService
	pb.UnimplementedSpaceServiceServer
	user user.UserServiceClient
}

func NewSpaceController(config *config.Config, l log.Logger, csl consul.Consul, queue repository.Queue, server *grpc.Server, service *SpaceService) (*SpaceController, error) {
	controller := &SpaceController{service: service}
	if pb.RegisterSpaceServiceServer(server, controller); config.Internal[enums.Service.Space].Export {
		csl.Register(consul.Service{Name: strings.Join([]string{config.Service.Name, enums.Service.Space}, "."), Port: config.Service.Port})
	}
	return controller, controller.UserConsumer(csl, queue)
}
func (s *SpaceController) UserConsumer(csl consul.Consul, queue repository.Queue) error {
	if conn, err := grpcs.NewConn(context.TODO(), csl.Address(enums.Service.User)); err != nil {
		return err
	} else {
		return queue.Recv(enums.Service.Space, enums.Service.User, func(ctx context.Context, key string, payload string) error {
			logger := log.With()
			id, err := strconv.ParseInt(payload, 10, 64)
			if err != nil {
				logger.Warnf("parse int err: %s", err)
				return err
			}
			if resp, err := user.NewUserServiceClient(conn).Info(ctx, &user.UserInfoRequest{Id: id}); err != nil {
				logger.Warnf("get user info err: %s", err, ctx)
			} else {
				logger.Infof("recv message %v %v %v", key, payload, resp.Data, ctx)
			}
			return nil
		})
	}

}
func (s *SpaceController) Create(ctx context.Context, req *pb.SpaceCreateRequest) (*pb.SpaceCreateReply, error) {
	space, err := s.service.Create(ctx, req.Name)
	return &pb.SpaceCreateReply{Data: SpaceDTO(space)}, err
}
func (s *SpaceController) Remove(ctx context.Context, req *pb.SpaceRemoveRequest) (*pb.SpaceRemoveReply, error) {
	return &pb.SpaceRemoveReply{}, s.service.Remove(ctx, req.Id)
}
func (s *SpaceController) Info(ctx context.Context, req *pb.SpaceInfoRequest) (*pb.SpaceInfoReply, error) {
	space, err := s.service.Info(ctx, req.Id)
	return &pb.SpaceInfoReply{Data: SpaceDTO(space)}, err
}
func (s *SpaceController) List(ctx context.Context, req *pb.SpaceListRequest) (*pb.SpaceListReply, error) {
	list, err := s.service.List(ctx, req.Page, req.Count)
	data := []*pb.Space{}
	for _, space := range list {
		data = append(data, SpaceDTO(space))
	}
	return &pb.SpaceListReply{Data: data}, err
}
func SpaceDTO(space *Space) *pb.Space {
	if space == nil {
		return nil
	}
	return &pb.Space{Id: space.ID, Name: space.Name, Email: space.Email}
}
