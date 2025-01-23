package kafka

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var errUnknownType = errors.New("unknown event type")

const flushTimeout = 10 * time.Second // 10 sec

type Producer interface {
	Produce(message, topic string) error
	Close()
}

type ProducerImpl struct {
	writer *kafka.Writer
}

func NewProducer(address []string) (*ProducerImpl, error) {
	if len(address) == 0 {
		return nil, errors.New("не указаны адреса брокеров Kafka")
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP(strings.Join(address, ",")),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return &ProducerImpl{writer: writer}, nil
}

func (p *ProducerImpl) Produce(message, topic string) error {
	if p.writer == nil {
		return errors.New("продюсер не инициализирован")
	}

	kafkaMsg := kafka.Message{
		Topic: topic,
		Value: []byte(message),
		Time:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), flushTimeout)
	defer cancel()

	err := p.writer.WriteMessages(ctx, kafkaMsg)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProducerImpl) Close() {
	if p.writer != nil {
		_ = p.writer.Close()
	}
}
