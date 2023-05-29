package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type queue struct {
	pulsar.Client
}

func (s *queue) Send(topic, key string, payload []byte) (string, error) {
	log.Infof("send message %s %s %s", topic, key, string(payload))
	p, e := s.Client.CreateProducer(pulsar.ProducerOptions{Topic: fmt.Sprintf("persistent://%s", topic)})
	if e != nil {
		return "", e
	}
	log.Infof("send message %s %s %s", topic, key, string(payload))
	if msgid, e := p.Send(context.TODO(), &pulsar.ProducerMessage{Key: key, Payload: payload}); e != nil {
		return "", e
	} else {
		return msgid.String(), nil
	}
}

func New(config *config.Config) (repository.Queue, error) {
	conf := config.Storage.Queue
	options := pulsar.ClientOptions{URL: fmt.Sprintf("pulsar://%s:%s", conf.Host, conf.Port)}
	if conf.Token != "" {
		options.Authentication = pulsar.NewAuthenticationToken(conf.Token)
	}
	client, e := pulsar.NewClient(options)
	return &queue{client}, e
}
