package market

import (
	"data-collector/kafka"
	"github.com/GiairoZeppeli/utils/context"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

const (
	positionInfoTopic  = "positionInfo"
	orderBookTopic     = "orderBook"
	fundingRateTopic   = "fundingRate"
	deliveryPriceTopic = "deliveryPrice"
	coinInfoTopic      = "coinInfo"
	tickersTopic       = "tickers"
)

type Service interface {
	GetPositionInfo(ctx context.MyContext, positionInfoParams map[string]interface{}) (string, error)
	GetOrderBook(ctx context.MyContext, params map[string]interface{}) (string, error)
	GetFundingRate(ctx context.MyContext, params map[string]interface{}) (string, error)
	GetDeliveryPrice(ctx context.MyContext, params map[string]interface{}) (string, error)
	GetCoinInfo(ctx context.MyContext, params map[string]interface{}) (string, error)
	GetTickers(ctx context.MyContext, params map[string]interface{}) (string, error)
}

type Impl struct {
	client   *bybit.Client
	producer kafka.Producer
}

func NewMarketService(client *bybit.Client, producer kafka.Producer) *Impl {
	return &Impl{
		client:   client,
		producer: producer,
	}
}

func (s *Impl) GetPositionInfo(ctx context.MyContext, positionInfoParams map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(positionInfoParams).GetPositionList(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetPositionInfo bybit api err:%v", err)
		return "", err
	}

	positionInfoResponse := bybit.PrettyPrint(serverResult)
	err = s.producer.Produce(positionInfoResponse, positionInfoTopic)
	if err != nil {
		ctx.Logger.Errorf("GetPositionInfo produce err:%v", err)
	}

	return positionInfoResponse, nil
}

func (s *Impl) GetOrderBook(ctx context.MyContext, params map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(params).GetOrderBookInfo(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetOrderBook bybit api err: %v", err)
		return "", err
	}

	orderBookResponse := bybit.PrettyPrint(serverResult)
	err = s.producer.Produce(orderBookResponse, orderBookTopic)
	if err != nil {
		ctx.Logger.Errorf("GetOrderBook produce err: %v", err)
	}

	return orderBookResponse, nil
}

func (s *Impl) GetFundingRate(ctx context.MyContext, params map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(params).GetFundingRateHistory(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetFundingRate bybit api err: %v", err)
		return "", err
	}

	fundingRateResponse := bybit.PrettyPrint(serverResult)
	err = s.producer.Produce(fundingRateResponse, fundingRateTopic)
	if err != nil {
		ctx.Logger.Errorf("GetFundingRate produce err: %v", err)
	}

	return fundingRateResponse, nil
}

func (s *Impl) GetDeliveryPrice(ctx context.MyContext, params map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(params).GetDeliveryPrice(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetDeliveryPrice error: %v", err)
		return "", err
	}

	response := bybit.PrettyPrint(serverResult)
	if err := s.producer.Produce(response, deliveryPriceTopic); err != nil {
		ctx.Logger.Errorf("Produce delivery price error: %v", err)
	}
	return response, nil
}

func (s *Impl) GetCoinInfo(ctx context.MyContext, params map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(params).GetCoinInfo(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetCoinInfo error: %v", err)
		return "", err
	}

	response := bybit.PrettyPrint(serverResult)
	if err := s.producer.Produce(response, coinInfoTopic); err != nil {
		ctx.Logger.Errorf("Produce coin info error: %v", err)
	}
	return response, nil
}

func (s *Impl) GetTickers(ctx context.MyContext, params map[string]interface{}) (string, error) {
	serverResult, err := s.client.NewUtaBybitServiceWithParams(params).GetMarketTickers(ctx.Ctx)
	if err != nil {
		ctx.Logger.Errorf("GetTickers error: %v", err)
		return "", err
	}

	response := bybit.PrettyPrint(serverResult)
	if err := s.producer.Produce(response, tickersTopic); err != nil {
		ctx.Logger.Errorf("Produce tickers error: %v", err)
	}
	return response, nil
}
