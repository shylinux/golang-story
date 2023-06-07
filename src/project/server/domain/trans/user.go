package trans

import (
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
)

func UserDTO(user *model.User) *pb.User {
	if user == nil {
		return nil
	}
	return &pb.User{Id: user.ID, Name: user.Name, Email: user.Email}
}