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
func Trans(ctx context.Context, key ...string) context.Context {
	kv := []string{}
	if len(key) == 0 {
		for k, v := range Dumps(ctx) {
			kv = append(kv, k, v)
		}
	} else if md, ok := metadata.FromIncomingContext(ctx); ok {
		for _, k := range key {
			kv = append(kv, k, md.Get(k)[0])
		}
	}
	return metadata.AppendToOutgoingContext(ctx, kv...)
}
func Dumps(ctx context.Context) map[string]string {
	md, _ := metadata.FromIncomingContext(ctx)
	meta := map[string]string{}
	for k, v := range md {
		meta[k] = v[0]
	}
	return meta
}
func Loads(ctx context.Context, meta map[string]string) context.Context {
	kv := []string{}
	for k, v := range meta {
		kv = append(kv, k, v)
	}
	return metadata.NewIncomingContext(ctx, metadata.Pairs(kv...))
}
