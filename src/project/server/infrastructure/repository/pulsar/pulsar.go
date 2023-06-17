package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/metadata"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/trace"
)

type queue struct {
	conf config.Queue
	pulsar.Client
}

func New(config *config.Config, consul consul.Consul) (repository.Queue, error) {
	conf := config.Engine.Queue
	if config.Proxy.Simple {
		conf.Enable = false
	}
	if !conf.Enable {
		return &queue{conf, nil}, nil
	}
	if list, err := consul.Resolve(config.WithDef(conf.Name, "pulsar")); err == nil && len(list) > 0 {
		conf.Host, conf.Port = list[0].Host, list[0].Port
	}
	options := pulsar.ClientOptions{URL: fmt.Sprintf("pulsar://%s:%d", conf.Host, conf.Port), Logger: &logger{}}
	if conf.Token != "" {
		options.Authentication = pulsar.NewAuthenticationToken(conf.Token)
	}
	if client, err := pulsar.NewClient(options); err != nil {
		logs.Errorf("engine connect pulsar %s:%d %s", conf.Host, conf.Port, err)
		return nil, errors.New(err, "engine connect pulsar failure")
	} else {
		logs.Infof("engine connect pulsar %s:%d", conf.Host, conf.Port)
		return &queue{conf, client}, nil
	}
}
func (s *queue) Send(ctx context.Context, topic, key string, payload []byte) (string, error) {
	if !s.conf.Enable {
		return "", nil
	}
	echo := func(res string, err error) (string, error) {
		if err != nil && err.Error() != "" {
			logs.Errorf("pulsar send %s:%s %s %s", topic, key, err, string(payload), ctx)
		} else {
			logs.Infof("pulsar send %s:%s %s %s", topic, key, res, string(payload), ctx)
		}
		return res, errors.New(err, "pulsar send failure")
	}
	p, err := s.Client.CreateProducer(pulsar.ProducerOptions{Topic: fmt.Sprintf("persistent://public/default/%s", topic)})
	if err != nil {
		return echo("", err)
	}
	defer p.Close()
	if msgid, err := p.Send(ctx, &pulsar.ProducerMessage{Key: key, Payload: payload, Properties: metadata.Dumps(ctx)}); err != nil {
		return echo("", err)
	} else {
		return echo(msgid.String(), nil)
	}
}
func (s *queue) Recv(ctx context.Context, name, topic string, cb func(ctx context.Context, key string, payload []byte)) error {
	if !s.conf.Enable {
		return nil
	}
	p, err := s.Client.Subscribe(pulsar.ConsumerOptions{Topic: fmt.Sprintf("persistent://public/default/%s", topic), SubscriptionName: name, Type: pulsar.Shared})
	if err != nil {
		logs.Errorf("%s subscribe %s %s %s", name, topic, err, errors.FileLine(2), ctx)
		return errors.New(err, "pulsar subscribe failure")
	} else {
		logs.Infof("%s subscribe %s %s", name, topic, errors.FileLine(2), ctx)
	}
	go func() {
		for {
			if msg, err := p.Receive(ctx); err != nil {
				logs.Errorf("pulsar recv %s %s", topic, err, ctx)
			} else {
				trace.ServerAccess(metadata.Loads(ctx, msg.Properties()), func(ctx context.Context) {
					logs.Infof("pulsar recv %s:%s %s %s", topic, msg.Key(), msg.ID().String(), string(msg.Payload()), ctx)
					cb(ctx, msg.Key(), msg.Payload())
					p.Ack(msg)
				})
			}
		}
	}()
	return nil
}
