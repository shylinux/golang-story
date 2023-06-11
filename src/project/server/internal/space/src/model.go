package space

import (
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
)

type Space struct {
	model.Common
	SpaceID int64 `gorm:"uniqueIndex:idx_spaceid"`
	Name    string
	Repos   string
	Binary  string
}

func (s Space) TableName() string { return enums.Service.Space }
func (s Space) GetKey() string    { return "space_id" }
func (s Space) GetID() string     { return fmt.Sprintf("%d", s.SpaceID) }
