package cli

import (
	"context"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
)

type AuthServiceCmds struct {
	consul consul.Consul
	client pb.AuthServiceClient
}

func NewAuthServiceCmds(cmds *cmds.Cmds, consul consul.Consul) (*AuthServiceCmds, error) {
	_cmds := &AuthServiceCmds{consul: consul}
	cmds.Register("auth", "auth", _cmds)
	return _cmds, nil
}

func (s *AuthServiceCmds) conn(ctx context.Context, arg ...string) {
	if s.client != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.AuthService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.client = pb.NewAuthServiceClient(conn)
}

func (s *AuthServiceCmds) Register(ctx context.Context, req *pb.AuthRegisterRequest) {
	s.conn(ctx)
	if res, err := s.client.Register(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *AuthServiceCmds) Login(ctx context.Context, req *pb.AuthLoginRequest) {
	s.conn(ctx)
	if res, err := s.client.Login(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *AuthServiceCmds) Logout(ctx context.Context, req *pb.AuthLogoutRequest) {
	s.conn(ctx)
	if res, err := s.client.Logout(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *AuthServiceCmds) Refresh(ctx context.Context, req *pb.AuthRefreshRequest) {
	s.conn(ctx)
	if res, err := s.client.Refresh(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *AuthServiceCmds) Verify(ctx context.Context, req *pb.AuthVerifyRequest) {
	s.conn(ctx)
	if res, err := s.client.Verify(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}
