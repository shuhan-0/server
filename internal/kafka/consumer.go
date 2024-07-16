package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.ConsumerGroup
}

func NewConsumer(brokers []string, groupID string) (*Consumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}
	return &Consumer{consumer: consumer}, nil
}

func (c *Consumer) Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error {
	return c.consumer.Consume(ctx, topics, handler)
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
