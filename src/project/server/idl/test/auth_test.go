package test

import (
	"context"
	"testing"
	"time"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
)

type AuthTestSuite struct {
	*tests.Suite
	ctx  context.Context
	auth pb.AuthServiceClient
	uid  int64
}

func (s *AuthTestSuite) SetupTest() {
	s.auth = pb.NewAuthServiceClient(s.Conn(s.ctx, pb.AuthService_ServiceDesc.ServiceName))
	s.uid = 13749123010001
}
func (s *AuthTestSuite) TestRegister() {
	cases := []struct {
		ok       bool
		username string
		password string
	}{
		{ok: false, username: "", password: ""},
		{ok: true, username: time.Now().Format("20060102150405.000"), password: "123456"},
	}
	for i, c := range cases {
		res, err := s.auth.Register(s.ctx, &pb.AuthRegisterRequest{Username: c.username, Password: c.password})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *AuthTestSuite) TestLogin() {
	cases := []struct {
		ok       bool
		username string
		password string
	}{
		{ok: false, username: "", password: ""},
		// {ok: true, username: "hi", password: "he"},
	}
	for i, c := range cases {
		res, err := s.auth.Login(s.ctx, &pb.AuthLoginRequest{Username: c.username, Password: c.password})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *AuthTestSuite) TestVerify() {
	cases := []struct {
		ok    bool
		token string
	}{
		{ok: false, token: ""},
		{ok: false, token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImhpIiwiaXNzIjoiZGVtby5hdXRoIiwiZXhwIjoxNjg2MzkwMzQ5fQ.BiFD-EV57gT4aDGUdQgC9-bXr9LppnDlh5S_K62HJzY"},
	}
	for i, c := range cases {
		res, err := s.auth.Verify(s.ctx, &pb.AuthVerifyRequest{Token: c.token})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func TestAuthTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &AuthTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
