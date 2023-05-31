package mysql

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type engine struct {
	db *gorm.DB
}

func (s engine) Insert(ctx context.Context, obj interface{}) error {
	return errors.New(s.db.WithContext(ctx).Create(obj).Error, "gorm create failure")
}
func (s engine) Delete(ctx context.Context, obj interface{}, id int64) error {
	return s.db.WithContext(ctx).Model(obj).Where("id = ?", id).Update("deleted", "1").Error
}
func (s engine) Update(ctx context.Context, obj interface{}, id int64) error {
	return s.db.WithContext(ctx).Model(obj).Where("id = ?", id).Updates(obj).Error
}
func (s engine) SelectOne(ctx context.Context, obj interface{}, id int64) (interface{}, error) {
	res := s.db.WithContext(ctx).Model(obj).Where("id = ?", id).First(obj)
	return obj, res.Error
}
func (s engine) SelectList(ctx context.Context, obj interface{}, res interface{}, page, count int64) (err error) {
	return s.db.WithContext(ctx).Model(obj).Where("deleted = 0").Offset(int((page - 1) * count)).Limit(int(count)).Find(res).Error
}
func (s *engine) AutoMigrate(obj ...interface{}) error {
	return s.db.AutoMigrate(obj...)
}

func New(consul consul.Consul, config *config.Config, logs logs.Logger) (repository.Engine, error) {
	conf := config.Storage.Engine
	if conf.Password == "" {
		return nil, fmt.Errorf("not found config password")
	}
	if list, err := consul.Resolve(conf.Name); err == nil && len(list) > 0 {
		conf.Host = list[0].Host
		conf.Port = list[0].Port
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database,
	)), &gorm.Config{})
	db.Callback().Create().After("gorm:after_create").Register("some:after", after)
	db.Callback().Update().After("gorm:after_update").Register("some:after", after)
	db.Callback().Delete().After("gorm:after_delete").Register("some:after", after)
	db.Callback().Query().After("gorm:after_query").Register("some:after", after)
	return &engine{db: db}, err
}

func after(db *gorm.DB) {
	logger := logs.With()
	logger.Infof(db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...), db.Statement.Context)
}
