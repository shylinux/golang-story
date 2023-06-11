package mysql

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type storage struct {
	*gorm.DB
}

func New(config *config.Config, consul consul.Consul) (repository.Storage, error) {
	conf := config.Engine.Storage
	if list, err := consul.Resolve(config.WithDef(conf.Name, "mysql")); err == nil && len(list) > 0 {
		conf.Host, conf.Port = list[0].Host, list[0].Port
	}
	if err := check(conf); err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database,
	)), &gorm.Config{})
	if err != nil {
		logs.Errorf("engine connect mysql %s:%d/%s %s", conf.Host, conf.Port, conf.Database, err)
		return nil, errors.New(err, "engine connect mysql failure")
	} else {
		hooks(db)
		logs.Infof("engine connect mysql %s:%d/%s", conf.Host, conf.Port, conf.Database)
		return &storage{db}, nil
	}
}
func (s storage) Insert(ctx context.Context, obj model.Model) error {
	err := s.DB.WithContext(ctx).Create(obj).Error
	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		return errors.NewAlreadyExists(err)
	}
	return errors.New(err, "mysql insert failure")
}
func (s storage) Delete(ctx context.Context, obj model.Model) error {
	return errors.New(s.DB.WithContext(ctx).Model(obj).Where(obj.GetKey()+" = ?", obj.GetID()).Update("deleted", "1").Error, "mysql delete failure")
}
func (s storage) Update(ctx context.Context, obj model.Model) error {
	err := s.DB.WithContext(ctx).Model(obj).Where(obj.GetKey()+" = ?", obj.GetID()).Updates(obj).Error
	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		return errors.NewAlreadyExists(err)
	}
	return errors.New(err, "mysql update failure")
}
func (s storage) SelectOne(ctx context.Context, obj model.Model) error {
	return errors.New(s.DB.WithContext(ctx).Model(obj).Where(obj.GetKey()+" = ?", obj.GetID()).First(obj).Error, "mysql select failure")
}
func (s storage) SelectList(ctx context.Context, obj model.Model, res interface{}, page, count int64, condition string, arg ...interface{}) (total int64, err error) {
	db := s.DB.WithContext(ctx).Model(obj).Where(condition+" deleted = 0", arg...)
	db.Count(&total)
	return total, errors.New(db.Offset(int((page-1)*count)).Limit(int(count)).Find(res).Error, "mysql select failure")
}
func (s *storage) AutoMigrate(obj ...interface{}) error {
	return errors.New(s.DB.AutoMigrate(obj...), "mysql migrate failure")
}
