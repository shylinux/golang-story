package space

import (
	"context"
	"encoding/json"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type UserConsumer struct{}

func NewUserConsumer(ctx context.Context, consul consul.Consul, queue repository.Queue) (*UserConsumer, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.UserService_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		client := pb.NewUserServiceClient(conn)
		return &UserConsumer{}, queue.Recv(ctx, enums.Service.Space, enums.Topic.User, func(ctx context.Context, key string, payload []byte) error {
			user := &pb.User{}
			if err := json.Unmarshal(payload, user); err != nil {
				logs.Errorf("get user info err: %s", err, ctx)
			} else if resp, err := client.Info(ctx, &pb.UserInfoRequest{Id: user.Id}); err != nil {
				logs.Errorf("get user info err: %s", err, ctx)
			} else {
				logs.Infof("recv message %v %v %#v %#v", key, string(payload), resp.Data, user, ctx)
			}
			return nil
		})
	}
}
