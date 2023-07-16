package space

import (
	"context"
	"encoding/json"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type UserConsumer struct{}

func NewUserConsumer(ctx context.Context, client pb.UserServiceClient, queue repository.Queue) (*UserConsumer, error) {
	return &UserConsumer{}, queue.Recv(ctx, enums.Service.Space, enums.Topic.User, func(ctx context.Context, key string, payload []byte) {
		user := &pb.User{}
		if err := json.Unmarshal(payload, user); err != nil {
			logs.Errorf("get user info err: %s", err, ctx)
		} else if resp, err := client.Info(ctx, &pb.UserInfoRequest{UserID: user.UserID}); err != nil {
			logs.Errorf("get user info err: %s", err, ctx)
		} else {
			logs.Infof("message recv %v %v %+v %+v", key, string(payload), resp.Data, user, ctx)
		}
	})
}
