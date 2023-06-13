package test

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
)

type UserServiceTestSuite struct {
	*tests.Suite
	ctx    context.Context
	client pb.UserServiceClient
}

func (s *UserServiceTestSuite) SetupTest() {
	s.client = pb.NewUserServiceClient(s.Conn(s.ctx, pb.UserService_ServiceDesc.ServiceName))
}

func (s *UserServiceTestSuite) TestCreate() {
	cases := []struct {
		OK       bool   `yaml:"ok"`
		Username string `yaml:"username"`
		Email    string `yaml:"email"`
	}{}
	s.Load("testdata/UserService/Create.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Create(s.ctx, &pb.UserCreateRequest{Username: c.Username, Email: c.Email})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *UserServiceTestSuite) TestInfo() {
	cases := []struct {
		OK     bool  `yaml:"ok"`
		UserID int64 `yaml:"userID"`
	}{}
	s.Load("testdata/UserService/Info.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Info(s.ctx, &pb.UserInfoRequest{UserID: c.UserID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *UserServiceTestSuite) TestList() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Page  int64  `yaml:"page"`
		Count int64  `yaml:"count"`
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}{}
	s.Load("testdata/UserService/List.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.List(s.ctx, &pb.UserListRequest{Page: c.Page, Count: c.Count, Key: c.Key, Value: c.Value})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *UserServiceTestSuite) TestRemove() {
	cases := []struct {
		OK     bool  `yaml:"ok"`
		UserID int64 `yaml:"userID"`
	}{}
	s.Load("testdata/UserService/Remove.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Remove(s.ctx, &pb.UserRemoveRequest{UserID: c.UserID})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *UserServiceTestSuite) TestRename() {
	cases := []struct {
		OK       bool   `yaml:"ok"`
		UserID   int64  `yaml:"userID"`
		Username string `yaml:"username"`
	}{}
	s.Load("testdata/UserService/Rename.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Rename(s.ctx, &pb.UserRenameRequest{UserID: c.UserID, Username: c.Username})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *UserServiceTestSuite) TestSearch() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Page  int64  `yaml:"page"`
		Count int64  `yaml:"count"`
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}{}
	s.Load("testdata/UserService/Search.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Search(s.ctx, &pb.UserSearchRequest{Page: c.Page, Count: c.Count, Key: c.Key, Value: c.Value})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func TestUserServiceTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &UserServiceTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
