package model

import (
	"fmt"
	"time"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type Model interface {
	TableName() string
	GetKey() string
	GetID() string
}

type Common struct {
	ID       int64
	Deleted  bool
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}

func (s Common) GetKey() string { return enums.Field.ID }
func (s Common) GetID() string  { return fmt.Sprintf("%d", s.ID) }
