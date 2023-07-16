package controller

import (
	"context"

	dt "shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/service"
)

type MachineController struct {
	pb.UnimplementedMachineServiceServer
	Main    *server.MainServer
	service *service.MachineService
	name    string
}

func NewMachineController(config *config.Config, server *server.MainServer, service *service.MachineService) *MachineController {
	controller := &MachineController{Main: server, service: service, name: pb.MachineService_ServiceDesc.ServiceName}
	if !config.Internal["mesh"].Export {
		return controller
	}
	server.Proxy.Register(controller.name, controller)
	server.Server.Register(&pb.MachineService_ServiceDesc, controller)
	consul.Tags = append(consul.Tags, controller.name)
	return controller
}
func (s *MachineController) Create(ctx context.Context, req *pb.MachineCreateRequest) (*pb.MachineCreateReply, error) {
	space, err := s.service.Create(ctx, req.Name)
	return &pb.MachineCreateReply{Data: trans.MachineDTO(space)}, errors.NewCreateFailResp(err)
}
func (s *MachineController) Remove(ctx context.Context, req *pb.MachineRemoveRequest) (*pb.MachineRemoveReply, error) {
	return &pb.MachineRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.MachineID))
}
func (s *MachineController) Rename(ctx context.Context, req *pb.MachineRenameRequest) (*pb.MachineRenameReply, error) {
	return &pb.MachineRenameReply{}, errors.NewModifyFailResp(s.service.Rename(ctx, req.MachineID, req.Name))
}
func (s *MachineController) Info(ctx context.Context, req *pb.MachineInfoRequest) (*pb.MachineInfoReply, error) {
	space, err := s.service.Info(ctx, req.MachineID)
	return &pb.MachineInfoReply{Data: trans.MachineDTO(space)}, errors.NewInfoFailResp(err)
}
func (s *MachineController) List(ctx context.Context, req *pb.MachineListRequest) (*pb.MachineListReply, error) {
	list, total, err := s.service.List(ctx, req.Page, req.Count, req.Key, req.Value)
	data := []*pb.Machine{}
	dt.ListDTO(list, trans.MachineDTO, &data)
	return &pb.MachineListReply{Data: data, Total: total}, errors.NewListFailResp(err)
}
