package mq

import (
	"context"

	"blog/internal/mq/kafka"

	"github.com/google/wire"
)

// ProviderSet is mq providers.
var ProviderSet = wire.NewSet(kafka.NewKafkaClient)

type MessageQueue interface {
	Produce(msg string) error
	Consume(consume func(ctx context.Context, msg string, err error) error) error
}
