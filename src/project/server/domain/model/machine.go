package model

import (
	"fmt"

	"gorm.io/gorm"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type Machine struct {
	gorm.Model
	MachineID int64
	Hostname  string `gorm:"uniqueIndex:idx_hostname_workpath"`
	Workpath  string `gorm:"uniqueIndex:idx_hostname_workpath"`
	Status    int32
}

func (s Machine) TableName() string { return enums.Table.Machine }
func (s Machine) GetKey() string    { return enums.Table.Machine + "_id" }
func (s Machine) GetID() string     { return fmt.Sprintf("%d", s.MachineID) }
