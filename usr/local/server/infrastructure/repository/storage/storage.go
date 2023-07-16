package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/metadata"
)

const (
	SQLITE = "sqlite"
	MYSQL  = "mysql"
)

type storage struct {
	*gorm.DB
}

func New(config *config.Config, consul consul.Consul) (repository.Storage, error) {
	switch conf := config.Engine.Storage; conf.Type {
	case SQLITE:
		db, err := gorm.Open(sqlite.Open(conf.Database+".db"), &gorm.Config{})
		if err != nil {
			logs.Errorf("engine connect sqlite %s %s", conf.Database, err)
			return nil, errors.New(err, "engine connect sqlite failure")
		} else {
			hooks(db)
			logs.Infof("engine connect sqlite %s", conf.Database)
			return &storage{db}, nil
		}
	case MYSQL:
		if list, err := consul.Resolve(config.WithDef(conf.Type, MYSQL)); err == nil && len(list) > 0 {
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
	default:
		return nil, errors.New(fmt.Errorf("not implement storage"), conf.Type)
	}
}
func (s storage) Insert(ctx context.Context, obj model.Model) error {
	if err := s.DB.WithContext(ctx).Create(obj).Error; err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		return errors.NewAlreadyExists(err)
	} else if err != nil && strings.Contains(err.Error(), "UNIQUE constraint") {
		return errors.NewAlreadyExists(err)
	} else {
		return errors.New(err, "storage insert failure")
	}
}
func (s storage) Delete(ctx context.Context, obj model.Model) error {
	return errors.New(s.DB.WithContext(ctx).Model(obj).Where(obj.GetKey()+" = ?", obj.GetID()).Update("deleted_at", time.Now()).Error, "storage delete failure")
}
func (s storage) Update(ctx context.Context, obj model.Model) error {
	err := s.DB.WithContext(ctx).Model(obj).Where(obj.GetKey()+" = ?", obj.GetID()).Updates(obj).Error
	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		return errors.NewAlreadyExists(err)
	}
	return errors.New(err, "storage update failure")
}
func (s storage) SelectOne(ctx context.Context, obj model.Model) error {
	return errors.New(s.DB.WithContext(ctx).Model(obj).Where(obj.GetKey()+" = ?", obj.GetID()).First(obj).Error, "storage select failure")
}
func (s storage) SelectList(ctx context.Context, obj model.Model, res interface{}, page, count int64, condition string, arg ...interface{}) (total int64, err error) {
	db := s.DB.WithContext(ctx).Model(obj).Where(condition, arg...)
	db.Count(&total)
	if preload := metadata.GetValue(ctx, metadata.PRELOAD); preload != "" {
		for _, preload := range strings.Split(preload, ",") {
			db = db.Preload(preload)
		}
	}
	return total, errors.New(db.Offset(int((page-1)*count)).Limit(int(count)).Find(res).Error, "storage select failure")
}
func (s *storage) AutoMigrate(obj ...interface{}) error {
	return errors.New(s.DB.AutoMigrate(obj...), "storage migrate failure")
}
