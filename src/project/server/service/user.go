package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type UserService struct {
	queue  repository.Queue
	cache  repository.Cache
	engine repository.Engine
}

func NewUserService(queue repository.Queue, cache repository.Cache, engine repository.Engine) (*UserService, error) {
	engine.AutoMigrate(&model.User{})
	return &UserService{queue, cache, engine}, nil
}
func (s *UserService) Create(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{Name: name}
	if err := s.engine.Insert(ctx, user); err != nil {
		return user, err
	}
	QueueSend(s.queue, ctx, enums.Operate.Create, user)
	CacheSet(s.cache, user)
	return user, nil
}
func (s *UserService) Remove(ctx context.Context, id int64) error {
	user := &model.User{Common: model.Common{ID: id}}
	QueueSend(s.queue, ctx, enums.Operate.Remove, user)
	CacheDel(s.cache, user)
	return s.engine.Delete(ctx, user, id)
}
func (s *UserService) Info(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	if err := CacheGet(s.cache, user); err == nil {
		return user, nil
	}
	data, err := s.engine.SelectOne(ctx, user, id)
	CacheSet(s.cache, data.(*model.User))
	return data.(*model.User), err
}
func (s *UserService) List(ctx context.Context, page int64, count int64) (res []*model.User, err error) {
	return res, s.engine.SelectList(ctx, &model.User{}, &res, page, count)
}
