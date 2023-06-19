package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
)

type ServiceService struct {
	*uuid.Generate
	storage repository.Storage
}

func NewServiceService(generate *uuid.Generate, storage repository.Storage) *ServiceService {
	storage.AutoMigrate(&model.Service{})
	return &ServiceService{generate, storage}
}
func (s *ServiceService) Create(ctx context.Context, machineID int64, mirror, config, dir, cmd, arg, env string) (*model.Service, error) {
	service := &model.Service{ServiceID: s.Generate.GenID(), MachineID: machineID, Mirror: mirror, Config: config, Dir: dir, Cmd: cmd, Arg: arg, Env: env}
	return service, errors.NewCreateFail(s.storage.Insert(ctx, service))
}
func (s *ServiceService) Remove(ctx context.Context, serviceID int64) error {
	return errors.NewRemoveFail(s.storage.Delete(ctx, &model.Service{ServiceID: serviceID}))
}
func (s *ServiceService) Change(ctx context.Context, serviceID int64, status int32) error {
	service := &model.Service{ServiceID: serviceID, Status: status}
	return errors.NewModifyFail(s.storage.Update(ctx, service))
}
func (s *ServiceService) Info(ctx context.Context, serviceID int64) (*model.Service, error) {
	service := &model.Service{ServiceID: serviceID}
	return service, errors.NewInfoFail(s.storage.SelectOne(ctx, service))
}
func (s *ServiceService) List(ctx context.Context, page, count int64, key, value string, machineID int64) (list []*model.Service, total int64, err error) {
	condition, arg := Clause(key != "" && value != "", key+" = ? and ", key, value)
	condition, arg = Clause(machineID != 0, "machine_id = ? and ", machineID)
	total, err = s.storage.SelectList(ctx, &model.Service{}, &list, page, count, condition, arg...)
	return list, total, errors.NewListFail(err)
}
