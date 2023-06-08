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

type storage struct {
	db *gorm.DB
}

func (s storage) Insert(ctx context.Context, obj interface{}) error {
	return errors.New(s.db.WithContext(ctx).Create(obj).Error, "gorm insert failure")
}
func (s storage) Delete(ctx context.Context, obj interface{}, id int64) error {
	return errors.New(s.db.WithContext(ctx).Model(obj).Where("id = ?", id).Update("deleted", "1").Error, "gorm delete failure")
}
func (s storage) Update(ctx context.Context, obj interface{}, id int64) error {
	return errors.New(s.db.WithContext(ctx).Model(obj).Where("id = ?", id).Updates(obj).Error, "gorm update failure")
}
func (s storage) SelectOne(ctx context.Context, obj interface{}, id int64) (interface{}, error) {
	res := s.db.WithContext(ctx).Model(obj).Where("id = ?", id).First(obj)
	return obj, errors.New(res.Error, "gorm select failure")
}
func (s storage) SelectList(ctx context.Context, obj interface{}, res interface{}, page, count int64, condition string, arg ...interface{}) (total int64, err error) {
	db := s.db.WithContext(ctx).Model(obj).Where(condition+" deleted = 0", arg...)
	db.Count(&total)
	return total, errors.New(db.Offset(int((page-1)*count)).Limit(int(count)).Find(res).Error, "gorm select failure")
}
func (s *storage) AutoMigrate(obj ...interface{}) error {
	return errors.New(s.db.AutoMigrate(obj...), "gorm migrate failure")
}

func New(config *config.Config, consul consul.Consul) (repository.Storage, error) {
	conf := config.Engine.Storage
	if list, err := consul.Resolve(config.ValueWithDef(conf.Name, "mysql")); err == nil && len(list) > 0 {
		conf.Host = list[0].Host
		conf.Port = list[0].Port
	}
	if err := check(conf); err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database,
	)), &gorm.Config{})
	if err != nil {
		return nil, errors.New(err, "connect service mysql failure")
	}
	hooks(db)
	logs.Infof("connect service mysql %s:%d", conf.Host, conf.Port)
	return &storage{db: db}, nil
}

func check(conf config.Storage) error {
	if conf.Username == "" {
		return errors.New(fmt.Errorf("not found config mysql username"), "")
	}
	if conf.Password == "" {
		return errors.New(fmt.Errorf("not found config mysql password"), "")
	}
	if conf.Database == "" {
		return errors.New(fmt.Errorf("not found config mysql database"), "")
	}
	if conf.Port == 0 {
		return errors.New(fmt.Errorf("not found config mysql port"), "")
	}
	return nil
}
func hooks(db *gorm.DB) {
	db.Callback().Create().After("gorm:after_create").Register("some:after", after)
	db.Callback().Update().After("gorm:after_update").Register("some:after", after)
	db.Callback().Delete().After("gorm:after_delete").Register("some:after", after)
	db.Callback().Query().After("gorm:after_query").Register("some:after", after)
}
func after(db *gorm.DB) {
	logs.Infof(db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...), db.Statement.Context)
}
