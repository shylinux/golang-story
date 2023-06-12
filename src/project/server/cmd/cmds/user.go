package cmds

import (
	"context"
	"fmt"
	"time"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
)

type UserCmds struct {
	client pb.UserServiceClient
}

func (s *UserCmds) Create(ctx context.Context, arg ...string) {
	if res, err := s.client.Create(ctx, &pb.UserCreateRequest{Username: time.Now().Format("20060102150405.000")}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("userID: %d\n", res.Data.UserID)
		fmt.Printf("username: %s\n", res.Data.Username)
	}
}

func (s *UserCmds) List(ctx context.Context, arg ...string) {
	if res, err := s.client.List(ctx, &pb.UserListRequest{}); err != nil {
		fmt.Println(err)
	} else {
		for _, user := range res.Data {
			fmt.Printf("userID: %d\n", user.UserID)
			fmt.Printf("username: %s\n", user.Username)
			fmt.Printf("\n")
		}
	}
}

func NewUserCmds(cmds *cmds.Cmds, ctx context.Context, config *config.Config, consul consul.Consul) (*UserCmds, error) {
	conn, err := grpc.NewConn(ctx, consul.Address(pb.UserService_ServiceDesc.ServiceName))
	if err != nil {
		return nil, err
	}
	user := &UserCmds{client: pb.NewUserServiceClient(conn)}
	cmds.Register("user", `user service client
  create username
    username length > 6
  list page count
    page value > 0
    count value > 10
`, user)
	return user, nil
}
