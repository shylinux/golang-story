package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/domain/model"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type MachineService struct {
	*uuid.Generate
	storage repository.Storage
}

func NewMachineService(generate *uuid.Generate, storage repository.Storage) *MachineService {
	storage.AutoMigrate(&model.Machine{})
	return &MachineService{generate, storage}
}
func (s *MachineService) Create(ctx context.Context, name string) (*model.Machine, error) {
	space := &model.Machine{
		MachineID: s.Generate.GenID(), Name: name,
	}
	return space, errors.NewCreateFail(s.storage.Insert(ctx, space))
}
func (s *MachineService) Remove(ctx context.Context, machineID int64) error {
	return errors.NewRemoveFail(s.storage.Delete(ctx, &model.Machine{
		MachineID: machineID,
	}))
}
func (s *MachineService) Rename(ctx context.Context, machineID int64, name string) error {
	machine := &model.Machine{
		MachineID: machineID, Name: name,
	}
	if err := s.storage.Update(ctx, machine); err != nil {
		return errors.NewModifyFail(err)
	}
	return nil
}
func (s *MachineService) Info(ctx context.Context, machineID int64) (*model.Machine, error) {
	machine := &model.Machine{
		MachineID: machineID,
	}
	return machine, errors.NewInfoFail(s.storage.SelectOne(ctx, machine))
}
func (s *MachineService) List(ctx context.Context, page, count int64, key, value string) (list []*model.Machine, total int64, err error) {
	condition, arg := service.Clause(key != "" && value != "", key+" = ? and ", key, value)
	total, err = s.storage.SelectList(ctx, &model.Machine{}, &list, page, count, condition, arg...)
	return list, total, errors.NewListFail(err)
}
