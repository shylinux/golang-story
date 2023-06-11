package space

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type SpaceService struct {
	storage repository.Storage
	*uuid.Generate
}

func NewSpaceService(config *config.Config, storage repository.Storage) *SpaceService {
	storage.AutoMigrate(&Space{})
	return &SpaceService{storage, uuid.New(config.Consul.WorkID)}
}
func (s *SpaceService) Create(ctx context.Context, name string) (*Space, error) {
	space := &Space{SpaceID: s.Generate.GenID(), Name: name}
	return space, errors.NewCreateFail(s.storage.Insert(ctx, space))
}
func (s *SpaceService) Remove(ctx context.Context, spaceID int64) error {
	return errors.NewRemoveFail(s.storage.Delete(ctx, &Space{SpaceID: spaceID}))
}
func (s *SpaceService) Info(ctx context.Context, spaceID int64) (*Space, error) {
	space := &Space{SpaceID: spaceID}
	return space, errors.NewInfoFail(s.storage.SelectOne(ctx, space))
}
func (s *SpaceService) List(ctx context.Context, page, count int64, key, value string) (list []*Space, total int64, err error) {
	condition, arg := service.Clause(key != "" && value != "", key+" = ? and ", key, value)
	total, err = s.storage.SelectList(ctx, &Space{}, &list, page, count, condition, arg...)
	return list, total, errors.NewListFail(err)
}
