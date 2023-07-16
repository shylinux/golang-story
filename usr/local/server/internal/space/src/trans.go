package space

import "shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"

func SpaceDTO(space *Space) *pb.Space {
	if space == nil {
		return nil
	}
	return &pb.Space{SpaceID: space.SpaceID, Name: space.Name, Repos: space.Repos, Binary: space.Binary}
}
