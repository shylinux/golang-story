package space

import "shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"

func SpaceDTO(space *Space) *pb.Space {
	if space == nil {
		return nil
	}
	return &pb.Space{Id: space.ID, Name: space.Name, Email: space.Email}
}
