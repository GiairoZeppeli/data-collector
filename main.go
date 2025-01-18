package main

import (
	"context"
	"data-collector/config"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	fmt.Println(logger.Level())

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()

	settings, err := config.NewSettings()
	if err != nil {
		logger.Errorf("fail to load settings: %s", err)
	}

	app := NewApp(ctx, logger, settings)
	if err := app.InitDatabase(); err != nil {
		logger.Errorf("fail to init database: %s", err)
	}

	if err := app.InitMQ(); err != nil {
		logger.Errorf("fail to init mq: %s", err)
	}

	app.InitService()
	if err = app.Run(); err != nil {
		logger.Errorf("fail to run: %s", err)
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan

	if err = app.Shutdown(); err != nil {
		logger.Errorf("fail to shutdown: %s", err)
		return
	}
}
