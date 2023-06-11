package grpc

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	if info.FullMethod == "/grpc.health.v1.Health/Check" {
		return handler(ctx, req)
	}
	begin := time.Now()
	logs.Infof("access %s %s", info.FullMethod, logs.Marshal(req), ctx)
	echo := func(res interface{}, err error) (interface{}, error) {
		if err != nil && err.Error() != "" {
			err = errors.ParseResp(err, "result failure").ToGRPC()
			logs.Warnf("result %s %s cost:%s", info.FullMethod, err, logs.Cost(begin), ctx)
			return nil, err
		} else {
			logs.Infof("result %s %s cost:%s", info.FullMethod, logs.Marshal(res), logs.Cost(begin), ctx)
			return res, nil
		}
	}
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			echo(nil, err.(error))
		}
	}()
	if valid, ok := req.(interface{ Validate() error }); ok {
		if err := valid.Validate(); err != nil && err.Error() != "" {
			return echo(nil, err)
		}
	}
	res, err = handler(ctx, req)
	return echo(res, err)
}
func clientInterceptor(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	begin := time.Now()
	logs.Infof("request %s %s", method, logs.Marshal(req), ctx)
	if err = invoker(ctx, method, req, res, cc, opts...); err != nil && err.Error() != "" {
		if status, ok := status.FromError(err); ok {
			err = errors.NewResp(fmt.Errorf(method), int64(status.Code()), status.Message())
		} else {
			err = errors.ParseResp(err, "response failure")
		}
		logs.Warnf("response %s %s cost:%s", method, err, logs.Cost(begin), ctx)
	} else {
		logs.Infof("response %s %s cost:%s", method, logs.Marshal(res), logs.Cost(begin), ctx)
	}
	return err
}
