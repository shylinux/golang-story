package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"shylinux.com/x/golang-story/src/project/server/domain"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type engine struct {
	db *gorm.DB
}

func (s engine) Insert(obj interface{}) error {
	return s.db.Create(obj).Error
}
func (s engine) Delete(obj interface{}, id int64) error {
	return s.db.Model(obj).Where("id = ?", id).Update("deleted", "1").Error
}
func (s engine) Update(obj interface{}, id int64) error {
	return s.db.Model(obj).Where("id = ?", id).Updates(obj).Error
}
func (s engine) SelectOne(obj interface{}, id int64) (interface{}, error) {
	res := s.db.Model(obj).Where("id = ?", id).First(obj)
	return obj, res.Error
}
func (s engine) SelectList(obj interface{}, res interface{}, page, count int64) (err error) {
	return s.db.Model(obj).Where("deleted = 0").Offset(int((page - 1) * count)).Limit(int(count)).Find(res).Error
}

func NewEngine(config *config.Config, log log.Logger) (repository.Engine, error) {
	conf := config.Storage.Engine
	if conf.Password == "" {
		return nil, fmt.Errorf("not found config password")
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database,
	)), &gorm.Config{Logger: logger.New(log, logger.Config{LogLevel: logger.Info})})
	db.AutoMigrate(&domain.User{})
	return &engine{db: db}, err
}
