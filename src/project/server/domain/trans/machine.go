package trans

import (
	"strings"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
)

func MachineDTO(machine *model.Machine) *pb.Machine {
	return reflect.Trans(&pb.Machine{StatusName: strings.TrimPrefix(pb.MachineStatus_name[machine.Status], "MACHINE_")}, machine).(*pb.Machine)
}
