package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
)

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	begin := time.Now()
	logger := log.With()
	logger.Infof("access %s [%+s]", info.FullMethod, req, ctx)
	if resp, err = handler(ctx, req); err == nil {
		logger.Infof("result %s cost: %s resp: [%+v]", info.FullMethod, time.Now().Sub(begin), resp, ctx)
	} else {
		logger.Infof("result %s cost: %s err: %s", info.FullMethod, time.Now().Sub(begin), err, ctx)
	}
	return resp, err
}
func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	begin := time.Now()
	logger := log.With("target", cc.Target())
	logger.Infof("request %s [%+v]", method, req, ctx)
	if err = invoker(ctx, method, req, reply, cc, opts...); err == nil {
		logger.Infof("response %s cost: %s reply: [%+v]", method, time.Now().Sub(begin), reply, ctx)
	} else {
		logger.Warnf("response %s cost: %s err: %s", method, time.Now().Sub(begin), err, ctx)
	}
	return err
}
