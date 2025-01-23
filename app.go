package main

import (
	ctx "context"
	"data-collector/api"
	"data-collector/config"
	"data-collector/kafka"
	"data-collector/service/history"
	"github.com/GiairoZeppeli/utils/context"
	"github.com/redis/go-redis/v9"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"go.uber.org/zap"
	"log"
	"time"
)

const (
	cacheKey = "coin:%s"
	ttl      = 10 * time.Minute
)

type App struct {
	ctx           context.MyContext
	server        *api.Server
	redis         *redis.Client
	settings      config.Settings
	kafkaProducer kafka.Producer
}

func NewApp(ctx ctx.Context, logger *zap.SugaredLogger, settings config.Settings) *App {
	return &App{
		ctx:      context.NewMyContext(ctx, logger),
		settings: settings,
	}
}

func (a *App) InitDatabase() error {
	return nil
}

func (a *App) InitMQ() error {
	producer, err := kafka.NewProducer(a.settings.Kafka.Address)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	a.kafkaProducer = producer

	return nil
}

func (a *App) InitService() {
	a.server = api.NewServer(a.ctx)

	bybitClient := bybit.NewBybitHttpClient("Rd2ENQWlsRwgWKYhUI", "mAZLYExoBZjMmqvlNb9STkhDY3U99L0TK7C9", bybit.WithBaseURL(bybit.TESTNET))

	historyService := history.NewHistoryService(bybitClient, a.kafkaProducer)
	a.server.HandleHistory(a.ctx, historyService)
}

func (a *App) Run() error {
	go func() {
		if err := a.server.Run(); err != nil {
			a.ctx.Logger.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	a.ctx.Logger.Info("run server")
	return nil
}

func (a *App) Shutdown() error {
	err := a.server.Shutdown(a.ctx)
	if err != nil {
		a.ctx.Logger.Errorf("Failed to disconnect from server %v", err)
		return err
	}

	a.ctx.Logger.Info("server shut down successfully")
	return nil
}
