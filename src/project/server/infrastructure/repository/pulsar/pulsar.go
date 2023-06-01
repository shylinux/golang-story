package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type queue struct {
	*config.Config
	pulsar.Client
}

func (s *queue) Send(ctx context.Context, topic, key string, payload []byte) (string, error) {
	logger := logs.With("key", key, "payload", string(payload))
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
	if msgid, e := p.Send(ctx, &pulsar.ProducerMessage{Key: key, Payload: payload, Properties: kv}); e != nil {
		logger.Warnf("send message topic: %s err: %s", topic, e, ctx)
		return "", e
	} else {
		logger.Infof("send message topic: %s msgid: %s", topic, msgid.String(), ctx)
		return msgid.String(), nil
	}
}
func (s *queue) Recv(ctx context.Context, name, topic string, cb func(ctx context.Context, key string, payload []byte) error) error {
	p, e := s.Client.Subscribe(pulsar.ConsumerOptions{
		Topic: fmt.Sprintf("persistent://public/default/%s", topic), SubscriptionName: name, Type: pulsar.Shared,
	})
	if e != nil {
		logs.Warnf("subscribe topic: %s err: %s", topic, e)
		return e
	} else {
		logs.Infof("subscribe topic: %s service: %s %s", topic, p.Subscription(), logs.FileLine(2))
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
					logger := logs.With("key", msg.Key(), "payload", string(msg.Payload()))
					logger.Infof("recv message topic: %s msgid: %s", topic, msg.ID().String(), ctx)
					cb(ctx, msg.Key(), msg.Payload())
					p.Ack(msg)
					return nil, nil
				})
			} else {
				logs.Warnf("recv message topic: %s err: %s", topic, e, ctx)
			}
		}
	}()
	return nil
}
func New(consul consul.Consul, config *config.Config) (repository.Queue, error) {
	conf := config.Engine.Queue
	if list, err := consul.Resolve(conf.Name); err == nil && len(list) > 0 {
		conf.Host = list[0].Host
		conf.Port = list[0].Port
	}
	options := pulsar.ClientOptions{URL: fmt.Sprintf("pulsar://%s:%d", conf.Host, conf.Port), Logger: &logger{}}
	if conf.Token != "" {
		options.Authentication = pulsar.NewAuthenticationToken(conf.Token)
	}
	client, e := pulsar.NewClient(options)
	return &queue{config, client}, e
}

type logger struct{}

func (l *logger) SubLogger(fields log.Fields) log.Logger             { return l }
func (l *logger) WithFields(fields log.Fields) log.Entry             { return l }
func (l *logger) WithField(name string, value interface{}) log.Entry { return l }
func (l *logger) WithError(err error) log.Entry                      { return l }
func (l *logger) Info(args ...interface{})                           {}
func (l *logger) Warn(args ...interface{})                           {}
func (l *logger) Error(args ...interface{})                          {}
func (l *logger) Debug(args ...interface{})                          {}
func (l *logger) Infof(format string, args ...interface{})           {}
func (l *logger) Warnf(format string, args ...interface{})           {}
func (l *logger) Errorf(format string, args ...interface{})          {}
func (l *logger) Debugf(format string, args ...interface{})          {}
