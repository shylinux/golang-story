package space

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type SpaceService struct {
	storage repository.Storage
}

func NewSpaceService(storage repository.Storage) *SpaceService {
	storage.AutoMigrate(&Space{})
	return &SpaceService{storage}
}
func (s *SpaceService) Create(ctx context.Context, name string) (*Space, error) {
	space := &Space{Name: name}
	return space, errors.NewCreateFail(s.storage.Insert(ctx, space))
}
func (s *SpaceService) Remove(ctx context.Context, id int64) error {
	return errors.NewRemoveFail(s.storage.Delete(ctx, &Space{}, id))
}
func (s *SpaceService) Info(ctx context.Context, id int64) (*Space, error) {
	data, err := s.storage.SelectOne(ctx, &Space{}, id)
	return data.(*Space), errors.NewInfoFail(err)
}
func (s *SpaceService) List(ctx context.Context, page int64, count int64) (res []*Space, err error) {
	return res, errors.NewListFail(s.storage.SelectList(ctx, &Space{}, &res, page, count))
}
