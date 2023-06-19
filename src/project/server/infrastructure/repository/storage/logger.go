package storage

import (
	"fmt"

	"gorm.io/gorm"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func check(conf config.Storage) error {
	if conf.Username == "" {
		return fmt.Errorf("not found config storage username")
	} else if conf.Password == "" {
		return fmt.Errorf("not found config storage password")
	} else if conf.Database == "" {
		return fmt.Errorf("not found config storage database")
	} else if conf.Port == 0 {
		return fmt.Errorf("not found config storage port")
	} else {
		return nil
	}
}
func hooks(db *gorm.DB) {
	db.Callback().Create().After("gorm:after_create").Register("some:after", after)
	db.Callback().Delete().After("gorm:after_delete").Register("some:after", after)
	db.Callback().Update().After("gorm:after_update").Register("some:after", after)
	db.Callback().Query().After("gorm:after_query").Register("some:after", after)
}
func after(db *gorm.DB) {
	logs.Infof(db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...), db.Statement.Context)
}
