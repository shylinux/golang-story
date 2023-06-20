package controller

import (
	"context"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/metadata"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type ServiceController struct {
	pb.UnimplementedServiceServiceServer
	Main    *server.MainServer
	service *service.ServiceService
	machine *service.MachineService
	name    string
}

func NewServiceController(config *config.Config, server *server.MainServer, service *service.ServiceService, machine *service.MachineService) *ServiceController {
	controller := &ServiceController{Main: server, service: service, machine: machine, name: pb.ServiceService_ServiceDesc.ServiceName}
	if !config.Internal[enums.Service.Mesh].Export {
		return controller
	}
	server.Proxy.Register(controller.name, controller)
	server.Server.Register(&pb.ServiceService_ServiceDesc, controller)
	consul.Tags = append(consul.Tags, controller.name)
	return controller
}
func (s *ServiceController) Create(ctx context.Context, req *pb.ServiceCreateRequest) (*pb.ServiceCreateReply, error) {
	service, err := s.service.Create(ctx, req.MachineID, req.Mirror, req.Config, req.Dir, req.Cmd, req.Arg, req.Env)
	return &pb.ServiceCreateReply{Data: trans.ServiceDTO(service)}, errors.NewCreateFailResp(err)
}
func (s *ServiceController) Remove(ctx context.Context, req *pb.ServiceRemoveRequest) (*pb.ServiceRemoveReply, error) {
	return &pb.ServiceRemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.ServiceID))
}
func (s *ServiceController) Inputs(ctx context.Context, req *pb.ServiceInputsRequest) (*pb.ServiceInputsReply, error) {
	res := &pb.ServiceInputsReply{Data: []*pb.ServiceInputsItem{}}
	switch req.Key {
	case "machineID":
		list, _, _ := s.machine.List(ctx, 1, 100, "", "")
		for _, item := range list {
			res.Data = append(res.Data, &pb.ServiceInputsItem{
				Value: fmt.Sprintf("%d", item.MachineID),
				Name:  fmt.Sprintf("%s-%s", item.Hostname, item.Workpath),
			})
		}
	case "mirror":
		list, _ := system.ReadDir("usr/mirror")
		for _, item := range list {
			res.Data = append(res.Data, &pb.ServiceInputsItem{Value: fmt.Sprintf("%s", item.Name())})
		}
	case "config":
		list, _ := system.ReadDir("usr/config")
		for _, item := range list {
			res.Data = append(res.Data, &pb.ServiceInputsItem{Value: fmt.Sprintf("%s", item.Name())})
		}
	}
	return res, errors.NewListFailResp(nil)
}
func (s *ServiceController) Deploy(ctx context.Context, req *pb.ServiceDeployRequest) (*pb.ServiceDeployReply, error) {
	return &pb.ServiceDeployReply{}, errors.NewModifyFailResp(s.service.Change(ctx, req.ServiceID, int32(pb.ServiceStatus_SERVICE_DEPLOY)))
}
func (s *ServiceController) Info(ctx context.Context, req *pb.ServiceInfoRequest) (*pb.ServiceInfoReply, error) {
	space, err := s.service.Info(ctx, req.ServiceID)
	return &pb.ServiceInfoReply{Data: trans.ServiceDTO(space)}, errors.NewInfoFailResp(err)
}
func (s *ServiceController) List(ctx context.Context, req *pb.ServiceListRequest) (*pb.ServiceListReply, error) {
	ctx = metadata.SetValue(ctx, metadata.PRELOAD, "Machine")
	list, total, err := s.service.List(ctx, req.Page, req.Count, req.Key, req.Value, req.MachineID)
	data := []*pb.Service{}
	reflect.TransList(list, trans.ServiceDTO, &data)
	return &pb.ServiceListReply{Data: data, Total: total}, errors.NewListFailResp(err)
}
