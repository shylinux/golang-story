package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
)

type MachineService struct {
	*uuid.Generate
	storage repository.Storage
}

func NewMachineService(generate *uuid.Generate, storage repository.Storage) *MachineService {
	storage.AutoMigrate(&model.Machine{})
	return &MachineService{generate, storage}
}
func (s *MachineService) Create(ctx context.Context, hostname, workpath string, status int32) (*model.Machine, error) {
	machine := &model.Machine{MachineID: s.Generate.GenID(), Hostname: hostname, Workpath: workpath, Status: status}
	return machine, errors.NewCreateFail(s.storage.Insert(ctx, machine))
}
func (s *MachineService) Remove(ctx context.Context, machineID int64) error {
	return errors.NewRemoveFail(s.storage.Delete(ctx, &model.Machine{MachineID: machineID}))
}
func (s *MachineService) Change(ctx context.Context, machineID int64, status int32) error {
	machine := &model.Machine{MachineID: machineID, Status: status}
	return errors.NewModifyFail(s.storage.Update(ctx, machine))
}
func (s *MachineService) Info(ctx context.Context, machineID int64) (*model.Machine, error) {
	machine := &model.Machine{MachineID: machineID}
	return machine, errors.NewInfoFail(s.storage.SelectOne(ctx, machine))
}
func (s *MachineService) List(ctx context.Context, page, count int64, key, value string) (list []*model.Machine, total int64, err error) {
	condition, arg := Clause(key != "" && value != "", key+" = ?", key, value)
	total, err = s.storage.SelectList(ctx, &model.Machine{}, &list, page, count, condition, arg...)
	return list, total, errors.NewListFail(err)
}
func (s *MachineService) Find(ctx context.Context, hostname, workpath string) (list []*model.Machine, total int64, err error) {
	total, err = s.storage.SelectList(ctx, &model.Machine{}, &list, 1, 10, "hostname = ? AND workpath = ?", hostname, workpath)
	return list, total, errors.NewListFail(err)
}
