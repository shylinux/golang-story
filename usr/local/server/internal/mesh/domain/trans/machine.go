package trans

import (
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/domain/model"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/idl/pb"
)

func MachineDTO(machine *model.Machine) *pb.Machine {
	if machine == nil {
		return nil
	}
	return &pb.Machine{
		MachineID: machine.MachineID,
		Name:      machine.Name,
	}
}
