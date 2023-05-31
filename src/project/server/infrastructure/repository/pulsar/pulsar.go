package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type queue struct {
	*config.Config
	pulsar.Client
}

func (s *queue) Send(ctx context.Context, topic, key string, payload string) (string, error) {
	logger := log.With("key", key, "payload", payload)
	p, e := s.Client.CreateProducer(pulsar.ProducerOptions{Topic: fmt.Sprintf("persistent://public/default/%s", topic)})
	if e != nil {
		logger.Warnf("send message topic: %s err: %s", topic, e, ctx)
		return "", e
	}
	defer p.Close()
	md, _ := metadata.FromIncomingContext(ctx)
	kv := map[string]string{}
	for k, v := range md {
		kv[k] = v[0]
	}
	if msgid, e := p.Send(ctx, &pulsar.ProducerMessage{Key: key, Payload: []byte(payload), Properties: kv}); e != nil {
		logger.Warnf("send message topic: %s err: %s", topic, e, ctx)
		return "", e
	} else {
		logger.Infof("send message topic: %s msgid: %s", topic, msgid.String(), ctx)
		return msgid.String(), nil
	}
}
func (s *queue) Recv(name, topic string, cb func(ctx context.Context, key string, payload string) error) error {
	ctx := context.TODO()
	logger := log.With()
	p, e := s.Client.Subscribe(pulsar.ConsumerOptions{
		Topic: fmt.Sprintf("persistent://public/default/%s", topic), Type: pulsar.KeyShared,
		SubscriptionName: fmt.Sprintf("%s-%s-%d-%s", s.Config.Service.Name, s.Config.Service.Host, s.Config.Service.Port, name),
	})
	if e != nil {
		logger.Warnf("subs message topic: %s err: %s", topic, e, ctx)
		return e
	} else {
		logger.Infof("subs message topic: %s self: %s", topic, p.Subscription(), ctx)
	}
	go func() {
		for {
			if msg, err := p.Receive(ctx); err == nil {
				kv := []string{}
				for k, v := range msg.Properties() {
					kv = append(kv, k, v)
				}
				ctx := metadata.NewIncomingContext(ctx, metadata.Pairs(kv...))
				otelgrpc.UnaryServerInterceptor()(ctx, nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, req interface{}) (interface{}, error) {
					logger := log.With("key", msg.Key(), "payload", string(msg.Payload()))
					logger.Infof("recv message topic: %s msgid: %s", topic, msg.ID().String(), ctx)
					log.Infof("recv message topic: %s msgid: %s", topic, msg.ID().String(), ctx)
					cb(ctx, msg.Key(), string(msg.Payload()))
					p.Ack(msg)
					return nil, nil

				})
			} else {
				logger.Warnf("recv message topic: %s err: %s", topic, e, ctx)
			}
		}
	}()
	return nil
}
func New(consul consul.Consul, config *config.Config) (repository.Queue, error) {
	conf := config.Storage.Queue
	if list, err := consul.Resolve("pulsar"); err == nil && len(list) > 0 {
		conf.Host = list[0].Host
		conf.Port = list[0].Port
	}
	options := pulsar.ClientOptions{URL: fmt.Sprintf("pulsar://%s:%d", conf.Host, conf.Port)}
	if conf.Token != "" {
		options.Authentication = pulsar.NewAuthenticationToken(conf.Token)
	}
	client, e := pulsar.NewClient(options)
	return &queue{config, client}, e
}
