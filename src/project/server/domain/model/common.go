package model

import "time"

type Model interface {
	TableName() string
	GetID() int64
}

type Common struct {
	ID       int64
	Deleted  bool
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}

func (s Common) GetID() int64 { return s.ID }
