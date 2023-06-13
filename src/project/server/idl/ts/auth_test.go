package test

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
)

type AuthServiceTestSuite struct {
	*tests.Suite
	ctx    context.Context
	client pb.AuthServiceClient
}

func (s *AuthServiceTestSuite) SetupTest() {
	s.client = pb.NewAuthServiceClient(s.Conn(s.ctx, pb.AuthService_ServiceDesc.ServiceName))
}

func (s *AuthServiceTestSuite) TestLogin() {
	cases := []struct {
		OK       bool   `yaml:"ok"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}{}
	s.Load("testdata/AuthService/Login.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Login(s.ctx, &pb.AuthLoginRequest{Username: c.Username, Password: c.Password})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *AuthServiceTestSuite) TestLogout() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Token string `yaml:"token"`
	}{}
	s.Load("testdata/AuthService/Logout.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Logout(s.ctx, &pb.AuthLogoutRequest{Token: c.Token})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *AuthServiceTestSuite) TestRefresh() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Token string `yaml:"token"`
	}{}
	s.Load("testdata/AuthService/Refresh.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Refresh(s.ctx, &pb.AuthRefreshRequest{Token: c.Token})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *AuthServiceTestSuite) TestRegister() {
	cases := []struct {
		OK       bool   `yaml:"ok"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Email    string `yaml:"email"`
		Phone    string `yaml:"phone"`
	}{}
	s.Load("testdata/AuthService/Register.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Register(s.ctx, &pb.AuthRegisterRequest{Username: c.Username, Password: c.Password, Email: c.Email, Phone: c.Phone})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func (s *AuthServiceTestSuite) TestVerify() {
	cases := []struct {
		OK    bool   `yaml:"ok"`
		Token string `yaml:"token"`
	}{}
	s.Load("testdata/AuthService/Verify.yaml", &cases)
	for i, c := range cases {
		res, err := s.client.Verify(s.ctx, &pb.AuthVerifyRequest{Token: c.Token})
		s.ConveySo(i, c.OK, c, res, err)
	}
}

func TestAuthServiceTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &AuthServiceTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
