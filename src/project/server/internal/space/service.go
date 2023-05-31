package space

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type SpaceService struct {
	engine repository.Engine
}

func NewSpaceService(engine repository.Engine) *SpaceService {
	engine.AutoMigrate(&Space{})
	return &SpaceService{engine}
}
func (s *SpaceService) Create(ctx context.Context, name string) (*Space, error) {
	space := &Space{Name: name}
	return space, errors.NewCreateFail(s.engine.Insert(ctx, space))
}
func (s *SpaceService) Remove(ctx context.Context, id int64) error {
	return errors.NewRemoveFail(s.engine.Delete(ctx, &Space{}, id))
}
func (s *SpaceService) Info(ctx context.Context, id int64) (*Space, error) {
	data, err := s.engine.SelectOne(ctx, &Space{}, id)
	return data.(*Space), errors.NewInfoFail(err)
}
func (s *SpaceService) List(ctx context.Context, page int64, count int64) (res []*Space, err error) {
	return res, errors.NewListFail(s.engine.SelectList(ctx, &Space{}, &res, page, count))
}
