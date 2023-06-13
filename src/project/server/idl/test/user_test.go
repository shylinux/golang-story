package test

import (
	"context"
	"testing"
	"time"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
)

type UserTestSuite struct {
	*tests.Suite
	ctx      context.Context
	client   pb.UserServiceClient
	userID   int64
	username string
}

func (s *UserTestSuite) newUsername() string {
	return time.Now().Format("20060102150405.000")
}
func (s *UserTestSuite) SetupTest() {
	s.username = s.newUsername()
	s.client = pb.NewUserServiceClient(s.Conn(s.ctx, pb.UserService_ServiceDesc.ServiceName))
	if res, err := s.client.Create(s.ctx, &pb.UserCreateRequest{Username: s.username}); err != nil {
		s.T().Error(err)
	} else {
		s.userID = res.Data.UserID
	}
}
func (s *UserTestSuite) TestCreate() {
	cases := []struct {
		ok       bool
		username string
	}{
		{ok: false, username: ""},
		{ok: false, username: "hi"},
		{ok: false, username: s.username},
		{ok: true, username: s.newUsername()},
	}
	for i, c := range cases {
		res, err := s.client.Create(s.ctx, &pb.UserCreateRequest{Username: c.username})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *UserTestSuite) TestRemove() {
	cases := []struct {
		ok     bool
		userID int64
	}{
		{ok: false, userID: 0},
		{ok: false, userID: -1},
		{ok: true, userID: s.userID},
	}
	for i, c := range cases {
		res, err := s.client.Remove(s.ctx, &pb.UserRemoveRequest{UserID: c.userID})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *UserTestSuite) TestRename() {
	cases := []struct {
		ok       bool
		userID   int64
		username string
	}{
		{ok: false, userID: 0, username: ""},
		{ok: false, userID: 0, username: "hi"},
		{ok: false, userID: s.userID, username: "hi"},
		{ok: true, userID: s.userID, username: s.newUsername()},
	}
	for i, c := range cases {
		res, err := s.client.Rename(s.ctx, &pb.UserRenameRequest{UserID: c.userID, Username: c.username})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *UserTestSuite) TestSearch() {
	cases := []struct {
		ok    bool
		key   string
		value string
	}{
		{ok: false, key: "username", value: ""},
		{ok: true, key: "username", value: "hi"},
		{ok: true, key: "username", value: "he*"},
	}
	for i, c := range cases {
		res, err := s.client.Search(s.ctx, &pb.UserSearchRequest{Key: c.key, Value: c.value})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *UserTestSuite) TestInfo() {
	cases := []struct {
		ok     bool
		userID int64
	}{
		{ok: false, userID: 0},
		{ok: false, userID: -1},
		{ok: true, userID: s.userID},
	}
	for i, c := range cases {
		res, err := s.client.Info(s.ctx, &pb.UserInfoRequest{UserID: c.userID})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *UserTestSuite) TestList() {
	cases := []struct {
		ok    bool
		page  int64
		count int64
	}{
		{ok: true, page: 0, count: 0},
		{ok: true, page: 0, count: 10},
		{ok: true, page: 1, count: 10},
		{ok: true, page: 1, count: 10},
	}
	for i, c := range cases {
		res, err := s.client.List(s.ctx, &pb.UserListRequest{Page: c.page, Count: c.count})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func TestUserTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &UserTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
