package kafka

import (
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
)

var errUnknownType = errors.New("unknown event type")

const flushTimeout = 10000 // 10 sec

type Producer interface {
	Produce(message, topic string) error
	Close()
}

type ProducerImpl struct {
	producer *kafka.Producer
}

func NewProducer(address []string) (*ProducerImpl, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
	}

	p, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return &ProducerImpl{p}, nil
}

func (p *ProducerImpl) Produce(message, topic string) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny},
		Value: []byte(message),
		Key:   nil,
	}
	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}

	event := <-kafkaChan
	switch ev := event.(type) {
	case *kafka.Message:
		return nil
	case *kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (p *ProducerImpl) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
