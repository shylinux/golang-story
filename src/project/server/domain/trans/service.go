package trans

import (
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
)

func ServiceDTO(service *model.Service) *pb.Service {
	return reflect.Trans(&pb.Service{Machine: MachineDTO(&service.Machine)}, service).(*pb.Service)
}
