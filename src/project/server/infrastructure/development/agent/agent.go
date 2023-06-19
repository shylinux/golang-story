package agent

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type AgentCmds struct {
	consul  consul.Consul
	machine pb.MachineServiceClient
	service pb.ServiceServiceClient
}

func (s *AgentCmds) conn(ctx context.Context, arg ...string) {
	if s.machine != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.MachineService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.machine = pb.NewMachineServiceClient(conn)
	if s.service != nil {
		return
	}
	conn, err = grpc.NewConn(ctx, s.consul.Address(pb.ServiceService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.service = pb.NewServiceServiceClient(conn)
}
func (s *AgentCmds) List(ctx context.Context, arg ...string) {
	s.conn(ctx)
	machineID := int64(0)
	if len(arg) > 0 {
		machineID, _ = strconv.ParseInt(arg[0], 10, 64)
	} else {
		if res, err := s.machine.Create(ctx, &pb.MachineCreateRequest{Hostname: system.Hostname(), Workpath: system.Workpath()}); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(logs.MarshalIndent(res))
			machineID = res.Data.MachineID
		}
	}
	for {
		if res, err := s.service.List(ctx, &pb.ServiceListRequest{MachineID: machineID}); err == nil {
			fmt.Println(system.Now())
			fmt.Println(system.MarshalIndent(res))
		} else {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Second)
	}
}
func NewAgentCmds(consul consul.Consul, cmds *cmds.Cmds, deploy *deploy.DeployCmds) *AgentCmds {
	s := &AgentCmds{consul: consul}
	cmds = cmds.Add("agent", "agent runtime cli", s.List)
	return s
}
