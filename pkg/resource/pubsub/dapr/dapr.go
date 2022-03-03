package dapr

import (
	"context"
	"os"

	cloudevents "github.com/cloudevents/sdk-go"
	daprSDK "github.com/dapr/go-sdk/client"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	xerrors "github.com/tkeel-io/core/pkg/errors"
	zfield "github.com/tkeel-io/core/pkg/logger"
	"github.com/tkeel-io/core/pkg/resource/pubsub"
	"github.com/tkeel-io/core/pkg/util/dapr"
	"github.com/tkeel-io/kit/log"
)

type daprMetadata struct {
	TopicName    string `json:"topic_name" mapstructure:"topic_name"`
	PubsubName   string `json:"pubsub_name" mapstructure:"pubsub_name"`
	ConsumerType string `json:"consumer_type" mapstructure:"consumer_type"`
}

type daprPubsub struct {
	id           string
	topicName    string
	pubsubName   string
	consumerType string
}

func (d *daprPubsub) ID() string {
	return d.id
}

func (d *daprPubsub) Send(ctx context.Context, event cloudevents.Event) error {
	var (
		err      error
		bytes    []byte
		metadata = make(map[string]string)
	)

	if bytes, err = event.MarshalJSON(); nil != err {
		return errors.Wrap(err, "dapr send")
	}

	log.Debug("pubsub.dapr send message",
		zfield.ID(d.id), zfield.Event(event),
		zfield.Pubsub(d.pubsubName), zfield.Topic(d.topicName))

	var conn dapr.Client
	if conn = dapr.Get().Select(); nil == conn {
		log.Error("nil connection",
			zfield.ID(d.id), zfield.Event(event),
			zfield.Pubsub(d.pubsubName), zfield.Topic(d.topicName))
		return errors.Wrap(xerrors.ErrConnectionNil, "dapr send")
	}

	err = conn.PublishEvent(
		ctx, d.pubsubName, d.topicName, bytes,
		daprSDK.PublishEventWithMetadata(metadata),
		daprSDK.PublishEventWithContentType(cloudevents.ApplicationJSON))
	return errors.Wrap(err, "dapr send")
}

func (d *daprPubsub) Received(ctx context.Context, handler pubsub.EventHandler) error {
	log.Debug("pubsub.dapr start receive message", zfield.ID(d.id))
	err := Get().Register(d.consumerType, d.pubsubName, d.topicName, &Consumer{id: d.id, handler: handler})
	return errors.Wrap(err, "register message handler")
}

func (d *daprPubsub) Commit(v interface{}) error {
	return nil
}

func (d *daprPubsub) Close() error {
	log.Debug("pubsub.dapr close", zfield.ID(d.id))
	err := Get().Unregister(d.consumerType, d.pubsubName, d.topicName, &Consumer{id: d.id})
	return errors.Wrap(err, "unregister message handler")
}

func init() {
	zfield.SuccessStatusEvent(os.Stdout, "Register Resource<pubsub.dapr> successful")
	pubsub.Register("dapr", func(id string, properties map[string]interface{}) (pubsub.Pubsub, error) {
		var daprMeta daprMetadata
		if err := mapstructure.Decode(properties, &daprMeta); nil != err {
			return nil, errors.Wrap(err, "decode pubsub.dapr configuration")
		}

		log.Info("create pubsub.dapr instance", zfield.ID(id))

		return &daprPubsub{
			id:           id,
			topicName:    daprMeta.TopicName,
			pubsubName:   daprMeta.PubsubName,
			consumerType: daprMeta.ConsumerType,
		}, nil
	})
}
