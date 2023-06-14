package service

import (
	"context"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/token"
)

type AuthService struct {
	token *token.Token
	user  *UserService
}

func NewAuthService(token *token.Token, user *UserService) *AuthService {
	return &AuthService{token: token, user: user}
}

func (s *AuthService) Register(ctx context.Context, username, password string, email, phone string) (string, error) {
	if _, err := s.user.Create(ctx, username, password, email, phone); err != nil {
		return "", err
	}
	return s.Login(ctx, username, password)
}
func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	list, _, err := s.user.List(ctx, 1, 1, enums.Field.Username, username)
	if err != nil || len(list) == 0 {
		return "", errors.NewNotFoundUser(fmt.Errorf("%v", err))
	}
	if password != list[0].Password {
		return "", errors.NewIncorrectPassword(fmt.Errorf("password incorrect"))
	}
	return s.token.Signed(username)
}
func (s *AuthService) Verify(ctx context.Context, token string) (string, error) {
	return s.token.Verify(token)
}
