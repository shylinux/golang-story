package test

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceServiceTestSuite struct {
	*tests.Suite
	ctx    context.Context
	client pb.SpaceServiceClient
}

func (s *SpaceServiceTestSuite) SetupTest() {
	s.client = pb.NewSpaceServiceClient(s.Conn(s.ctx, pb.SpaceService_ServiceDesc.ServiceName))
}

func (s *SpaceServiceTestSuite) TestCreate() {
	cases := []struct {
		OK     bool   `yaml:"ok"`
		Name   string `yaml:"name"`
		Repos  string `yaml:"repos"`
		Binary string `yaml:"binary"`
	}{}
	s.Load("testdata/SpaceService/Create.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Create(s.ctx, &pb.SpaceCreateRequest{Name: c.Name, Repos: c.Repos, Binary: c.Binary})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *SpaceServiceTestSuite) TestInfo() {
	cases := []struct {
		OK      bool  `yaml:"ok"`
		SpaceID int64 `yaml:"spaceID"`
	}{}
	s.Load("testdata/SpaceService/Info.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Info(s.ctx, &pb.SpaceInfoRequest{SpaceID: c.SpaceID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *SpaceServiceTestSuite) TestList() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Page  int64  `yaml:"page"`
		Count int64  `yaml:"count"`
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}{}
	s.Load("testdata/SpaceService/List.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.List(s.ctx, &pb.SpaceListRequest{Page: c.Page, Count: c.Count, Key: c.Key, Value: c.Value})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *SpaceServiceTestSuite) TestRemove() {
	cases := []struct {
		OK      bool  `yaml:"ok"`
		SpaceID int64 `yaml:"spaceID"`
	}{}
	s.Load("testdata/SpaceService/Remove.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Remove(s.ctx, &pb.SpaceRemoveRequest{SpaceID: c.SpaceID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func TestSpaceServiceTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &SpaceServiceTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
