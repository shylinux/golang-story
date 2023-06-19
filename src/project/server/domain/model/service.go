package model

import (
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type Service struct {
	Common
	MachineID int64
	ServiceID int64
	Status    int32
	Mirror    string
	Config    string
	Dir       string
	Cmd       string
	Arg       string
	Env       string
	Machine   Machine `gorm:"foreignkey:MachineID;references:MachineID"`
}

func (s Service) TableName() string { return enums.Table.Service }
func (s Service) GetKey() string    { return enums.Table.Service + "_id" }
func (s Service) GetID() string     { return fmt.Sprintf("%d", s.ServiceID) }
