package domain

import "time"

type Common struct {
	ID       int64
	Deleted  bool
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}
