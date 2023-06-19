package trans

import (
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
)

func UserDTO(user *model.User) *pb.User {
	return reflect.Trans(&pb.User{}, user).(*pb.User)
}
