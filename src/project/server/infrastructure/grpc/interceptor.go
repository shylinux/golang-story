package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.FullMethod == "/grpc.health.v1.Health/Check" {
		return handler(ctx, req)
	}
	begin := time.Now()
	logs.Infof("access %s [%+s]", info.FullMethod, req, ctx)
	if resp, err = handler(ctx, req); err == nil {
		logs.Infof("result %s cost: %s resp: [%+v]", info.FullMethod, time.Now().Sub(begin), resp, ctx)
	} else {
		logs.Warnf("result %s cost: %s err: %s", info.FullMethod, time.Now().Sub(begin), err, ctx)
	}
	return resp, errors.NewResp(err, enums.Errors.Unknown, "result failure")
}
func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	begin := time.Now()
	logs.Infof("request %s [%+v]", method, req, ctx)
	if err = invoker(ctx, method, req, reply, cc, opts...); err == nil {
		logs.Infof("response %s cost: %s reply: [%+v]", method, time.Now().Sub(begin), reply, ctx)
	} else {
		logs.Warnf("response %s cost: %s err: %s", method, time.Now().Sub(begin), err, ctx)
	}
	return errors.New(err, "request %s failure", method)
}
