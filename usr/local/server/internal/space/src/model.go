package space

import (
	"fmt"

	"gorm.io/gorm"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type Space struct {
	gorm.Model
	SpaceID int64 `gorm:"uniqueIndex:idx_spaceid"`
	Name    string
	Repos   string
	Binary  string
}

func (s Space) TableName() string { return enums.Service.Space }
func (s Space) GetKey() string    { return enums.Service.Space + "_id" }
func (s Space) GetID() string     { return fmt.Sprintf("%d", s.SpaceID) }
