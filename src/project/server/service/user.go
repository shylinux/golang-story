package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type UserService struct {
	queue   repository.Queue
	cache   repository.Cache
	storage repository.Storage
}

func NewUserService(queue repository.Queue, cache repository.Cache, storage repository.Storage) (*UserService, error) {
	storage.AutoMigrate(&model.User{})
	return &UserService{queue, cache, storage}, nil
}
func (s *UserService) Create(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{Name: name}
	if err := s.storage.Insert(ctx, user); err != nil {
		return user, errors.NewCreateFail(err)
	}
	if err := QueueSend(s.queue, ctx, enums.Topic.User, enums.Operate.Create, trans.UserDTO(user)); err != nil {
		logs.Errorf(errors.New(err, "send message failure topic: %s operate: %s payload: %+s",
			enums.Topic.User, enums.Operate.Create, trans.UserDTO(user)).Error(), ctx)
	}
	if err := CacheSet(s.cache, user); err != nil {
		logs.Warnf(errors.New(err, "set cache failure").Error(), ctx)
	}
	return user, nil
}
func (s *UserService) Remove(ctx context.Context, id int64) error {
	user := &model.User{Common: model.Common{ID: id}}
	if err := QueueSend(s.queue, ctx, enums.Topic.User, enums.Operate.Remove, trans.UserDTO(user)); err != nil {
		logs.Errorf(errors.New(err, "send message failure topic: %s operate: %s payload: %+s",
			enums.Topic.User, enums.Operate.Remove, trans.UserDTO(user)).Error(), ctx)
	}
	if err := CacheDel(s.cache, user); err != nil {
		logs.Warnf(errors.New(err, "del cache failure").Error(), ctx)
	}
	return errors.NewRemoveFail(s.storage.Delete(ctx, user, id))
}
func (s *UserService) Info(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	if err := CacheGet(s.cache, user); err == nil {
		return user, nil
	}
	data, err := s.storage.SelectOne(ctx, user, id)
	if err := CacheSet(s.cache, data.(*model.User)); err != nil {
		logs.Warnf(errors.New(err, "set cache failure").Error(), ctx)
	}
	return data.(*model.User), errors.NewInfoFail(err)
}
func (s *UserService) List(ctx context.Context, page int64, count int64) (res []*model.User, err error) {
	return res, errors.NewListFail(s.storage.SelectList(ctx, &model.User{}, &res, page, count))
}
