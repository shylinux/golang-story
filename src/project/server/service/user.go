package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
)

type UserService struct {
	*uuid.Generate
	cache   repository.Cache
	queue   repository.Queue
	search  repository.Search
	storage repository.Storage
}

func NewUserService(generate *uuid.Generate, cache repository.Cache, queue repository.Queue, search repository.Search, storage repository.Storage) (*UserService, error) {
	storage.AutoMigrate(&model.User{})
	return &UserService{generate, cache, queue, search, storage}, nil
}
func (s *UserService) Create(ctx context.Context, username, password, email, phone string) (*model.User, error) {
	user := &model.User{UserID: s.Generate.GenID(), Username: username, Password: password, Email: email, Phone: phone}
	if err := s.storage.Insert(ctx, user); err != nil {
		return user, errors.NewCreateFail(err)
	}
	CacheSet(ctx, s.cache, enums.Cache.User, user)
	SearchUpdate(ctx, s.search, enums.Search.User, user.UserID, trans.UserDTO(user))
	QueueSend(ctx, s.queue, enums.Topic.User, enums.Operate.Create, trans.UserDTO(user))
	return user, nil
}
func (s *UserService) Remove(ctx context.Context, userID int64) error {
	user := &model.User{UserID: userID}
	if err := s.storage.Delete(ctx, user); err != nil {
		return errors.NewRemoveFail(err)
	}
	CacheDel(ctx, s.cache, enums.Cache.User, user)
	SearchDelete(ctx, s.search, enums.Search.User, user.UserID)
	QueueSend(ctx, s.queue, enums.Topic.User, enums.Operate.Remove, trans.UserDTO(user))
	return nil
}
func (s *UserService) Rename(ctx context.Context, userID int64, username string) error {
	user := &model.User{UserID: userID, Username: username}
	if err := s.storage.Update(ctx, user); err != nil {
		return errors.NewModifyFail(err)
	}
	CacheDel(ctx, s.cache, enums.Cache.User, user)
	SearchUpdate(ctx, s.search, enums.Search.User, user.UserID, trans.UserDTO(user))
	return nil
}
func (s *UserService) Search(ctx context.Context, page, count int64, key, value string) (res []*model.User, total int64, err error) {
	total, err = s.search.Query(ctx, enums.Search.User, &res, page, count, key, value)
	return res, total, errors.NewSearchFail(err)
}
func (s *UserService) Info(ctx context.Context, userID int64) (*model.User, error) {
	user := &model.User{UserID: userID}
	if err := CacheGet(ctx, s.cache, enums.Cache.User, user); err == nil {
		return user, nil
	}
	if err := s.storage.SelectOne(ctx, user); err != nil {
		return nil, errors.NewInfoFail(err)
	}
	CacheSet(ctx, s.cache, enums.Cache.User, user)
	return user, nil
}
func (s *UserService) List(ctx context.Context, page int64, count int64, key, value string) (list []*model.User, total int64, err error) {
	condition, arg := Clause(key != "" && value != "", key+" = ? and ", value)
	total, err = s.storage.SelectList(ctx, &model.User{}, &list, page, count, condition, arg...)
	return list, total, errors.NewListFail(err)
}
