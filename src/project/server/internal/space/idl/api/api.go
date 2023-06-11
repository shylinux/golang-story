package api

import "shylinux.com/x/golang-story/src/project/server/infrastructure/container"

func Init(container *container.Container) {

	container.Provide(NewSpaceServiceClient)

}
