package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	USERNAME = "username"
)

func SetValue(ctx context.Context, key, value string) context.Context {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return metadata.NewIncomingContext(ctx, metadata.Join(md, metadata.Pairs(key, value)))
	} else {
		return metadata.NewIncomingContext(ctx, metadata.Pairs(key, value))
	}
}
func GetValue(ctx context.Context, key string) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if value := md.Get(key); len(value) > 0 {
			return value[0]
		}
	}
	return ""
}
