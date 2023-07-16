package model

import (
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
)

type Machine struct {
	model.Common
	MachineID int64
	Name      string
}

func (s Machine) TableName() string { return "machine" }
func (s Machine) GetKey() string    { return "machine_id" }
func (s Machine) GetID() string     { return fmt.Sprintf("%d", s.MachineID) }
