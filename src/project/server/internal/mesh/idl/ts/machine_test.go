package test

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/tests"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/idl/pb"
)

type MachineServiceTestSuite struct {
	*tests.Suite
	ctx    context.Context
	client pb.MachineServiceClient
}

func (s *MachineServiceTestSuite) SetupTest() {
	s.client = pb.NewMachineServiceClient(s.Conn(s.ctx, pb.MachineService_ServiceDesc.ServiceName))
}

func (s *MachineServiceTestSuite) TestCreate() {
	cases := []struct {
		OK   bool   `yaml:"ok"`
		Name string `yaml:"name"`
	}{}
	s.Load("testdata/MachineService/Create.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Create(s.ctx, &pb.MachineCreateRequest{Name: c.Name})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *MachineServiceTestSuite) TestInfo() {
	cases := []struct {
		OK        bool  `yaml:"ok"`
		MachineID int64 `yaml:"MachineID"`
	}{}
	s.Load("testdata/MachineService/Info.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Info(s.ctx, &pb.MachineInfoRequest{MachineID: c.MachineID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *MachineServiceTestSuite) TestList() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Page  int64  `yaml:"page"`
		Count int64  `yaml:"count"`
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}{}
	s.Load("testdata/MachineService/List.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.List(s.ctx, &pb.MachineListRequest{Page: c.Page, Count: c.Count, Key: c.Key, Value: c.Value})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *MachineServiceTestSuite) TestRemove() {
	cases := []struct {
		OK        bool  `yaml:"ok"`
		MachineID int64 `yaml:"MachineID"`
	}{}
	s.Load("testdata/MachineService/Remove.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Remove(s.ctx, &pb.MachineRemoveRequest{MachineID: c.MachineID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func TestMachineServiceTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &MachineServiceTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
