package service

import (
	"context"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type UserService struct {
	queue  repository.Queue
	cache  repository.Cache
	engine repository.Engine
}

func NewUserService(queue repository.Queue, cache repository.Cache, engine repository.Engine) *UserService {
	return &UserService{queue, cache, engine}
}
func (s *UserService) Create(ctx context.Context, name string) (*domain.User, error) {
	user := &domain.User{Name: name}
	if err := s.engine.Insert(user); err != nil {
		return user, err
	}
	s.queue.Send("user", "create", []byte(fmt.Sprintf("%d", user.ID)))
	CacheSet(s.cache, user.ID, user)
	return user, nil
}
func (s *UserService) Remove(ctx context.Context, id int64) error {
	user := &domain.User{}
	CacheDel(s.cache, id, user)
	return s.engine.Delete(user, id)
}
func (s *UserService) Info(ctx context.Context, id int64) (*domain.User, error) {
	user := &domain.User{}
	if err := CacheGet(s.cache, id, user); err == nil {
		return user, nil
	}
	data, err := s.engine.SelectOne(user, id)
	CacheSet(s.cache, id, data)
	return data.(*domain.User), err
}
func (s *UserService) List(ctx context.Context, page int64, count int64) (res []*domain.User, err error) {
	return res, s.engine.SelectList(&domain.User{}, &res, page, count)
}
