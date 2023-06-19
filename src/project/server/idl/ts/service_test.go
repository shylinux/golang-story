package test

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/tests"
)

type ServiceServiceTestSuite struct {
	*tests.Suite
	ctx    context.Context
	client pb.ServiceServiceClient
}

func (s *ServiceServiceTestSuite) SetupTest() {
	s.client = pb.NewServiceServiceClient(s.Conn(s.ctx, pb.ServiceService_ServiceDesc.ServiceName))
}

func (s *ServiceServiceTestSuite) TestCreate() {
	cases := []struct {
		OK        bool   `yaml:"ok"`
		MachineID int64  `yaml:"machineID"`
		Mirror    string `yaml:"mirror"`
		Config    string `yaml:"config"`
		Dir       string `yaml:"dir"`
		Cmd       string `yaml:"cmd"`
		Arg       string `yaml:"arg"`
		Env       string `yaml:"env"`
	}{}
	s.Load("testdata/ServiceService/Create.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Create(s.ctx, &pb.ServiceCreateRequest{MachineID: c.MachineID, Mirror: c.Mirror, Config: c.Config, Dir: c.Dir, Cmd: c.Cmd, Arg: c.Arg, Env: c.Env})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *ServiceServiceTestSuite) TestDeploy() {
	cases := []struct {
		OK        bool  `yaml:"ok"`
		ServiceID int64 `yaml:"serviceID"`
	}{}
	s.Load("testdata/ServiceService/Deploy.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Deploy(s.ctx, &pb.ServiceDeployRequest{ServiceID: c.ServiceID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *ServiceServiceTestSuite) TestInfo() {
	cases := []struct {
		OK        bool  `yaml:"ok"`
		ServiceID int64 `yaml:"serviceID"`
	}{}
	s.Load("testdata/ServiceService/Info.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Info(s.ctx, &pb.ServiceInfoRequest{ServiceID: c.ServiceID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *ServiceServiceTestSuite) TestInputs() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}{}
	s.Load("testdata/ServiceService/Inputs.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Inputs(s.ctx, &pb.ServiceInputsRequest{Key: c.Key, Value: c.Value})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *ServiceServiceTestSuite) TestList() {
	cases := []struct {
		OK        bool   `yaml:"ok"`
		Page      int64  `yaml:"page"`
		Count     int64  `yaml:"count"`
		Key       string `yaml:"key"`
		Value     string `yaml:"value"`
		MachineID int64  `yaml:"machineID"`
	}{}
	s.Load("testdata/ServiceService/List.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.List(s.ctx, &pb.ServiceListRequest{Page: c.Page, Count: c.Count, Key: c.Key, Value: c.Value, MachineID: c.MachineID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *ServiceServiceTestSuite) TestRemove() {
	cases := []struct {
		OK        bool  `yaml:"ok"`
		ServiceID int64 `yaml:"serviceID"`
	}{}
	s.Load("testdata/ServiceService/Remove.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Remove(s.ctx, &pb.ServiceRemoveRequest{ServiceID: c.ServiceID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func TestServiceServiceTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &ServiceServiceTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
