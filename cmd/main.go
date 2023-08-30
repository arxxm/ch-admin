package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/config"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/repository"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/server"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/server/consumer"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/logger"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/mq"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "prod"
)

// init loggers
func init() {
	logger.ZapLoggerInit()
}

func main() {
	// init configs
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize config, err: %v", err))
	}

	// init context
	ctx, cancel := context.WithCancel(context.Background())
	errConsGroup, errConsGroupCtx := errgroup.WithContext(ctx)
	defer cancel()

	// init nats client
	nc, err := mq.NewNats(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("nats: failed to connect: %v", err.Error()))
	}
	// close nats connection
	defer nc.Close()

	// init dep-s
	r := repository.New(cfg)

	// init consumer
	cons := consumer.New(cfg, nc, r)

	//init server
	srv := server.NewServer(cons, errConsGroup, r, nc, cfg)

	go srv.ListenAndServe(errConsGroupCtx)

	logger.Info("clickhouse-admin-sms-consumer started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-quit

	// wait finish all consumer
	if err = srv.GracefulStop(); err != nil {
		logger.Fatal(fmt.Sprintf("consumer shutdown with err, call_type: %v, err: %v", osSignal, err))
	}

	logger.Info(fmt.Sprintf("program shutdown... call_type: %v", osSignal))

}
