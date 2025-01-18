package history

import (
	"github.com/GiairoZeppeli/utils/context"
	"github.com/GiairoZeppeli/utils/kafka"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type Service interface {
	GetCandleHistory(ctx context.MyContext, candleHistoryParams map[string]interface{}) (string, error)
}

type Impl struct {
	client   *bybit.Client
	producer kafka.Producer
}

func NewHistoryService(client *bybit.Client, producer kafka.Producer) *Impl {
	return &Impl{
		client:   client,
		producer: producer,
	}
}

const (
	topic  = "history"
	broker = "localhost:9092"
)

func (s *Impl) GetCandleHistory(ctx context.MyContext, candleHistoryParams map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(candleHistoryParams).GetMarkPriceKline(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetCandleHistory bybit api err:%v", err)
		return "", err
	}
	candleHistoryResponse := bybit.PrettyPrint(serverResult)
	err = s.producer.Produce(candleHistoryResponse, topic)
	if err != nil {
		ctx.Logger.Errorf("GetCandleHistory produce err:%v", err)
	}

	return candleHistoryResponse, nil
}
