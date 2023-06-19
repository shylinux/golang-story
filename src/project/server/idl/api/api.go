package api

import "shylinux.com/x/golang-story/src/project/server/infrastructure/container"

func Init(c *container.Container) {

	c.Provide(NewAuthServiceClient)

	c.Provide(NewMachineServiceClient)

	c.Provide(NewServiceServiceClient)

	c.Provide(NewUserServiceClient)

}
