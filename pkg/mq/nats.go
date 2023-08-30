package mq

import (
	"fmt"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/config"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/logger"
	"time"

	"github.com/nats-io/nats.go"
)

func NewNats(cfg *config.Config) (*nats.Conn, error) {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts := nats.GetDefaultOptions()
	opts.Name = "NATS Sample Subscriber"
	opts.ReconnectWait = reconnectDelay
	opts.AllowReconnect = true
	opts.MaxReconnect = int(totalWait / reconnectDelay)
	opts.Url = cfg.Nats.Address

	opts.DisconnectedCB = func(nc *nats.Conn) {
		logger.Warn(fmt.Sprintf("NATS disconnected, will attempt reconnects for %v", totalWait.Minutes()))
	}
	opts.ReconnectedCB = func(nc *nats.Conn) {
		logger.Warn(fmt.Sprintf("NATS reconnected, url: %v", nc.ConnectedUrl()))
	}
	opts.ClosedCB = func(nc *nats.Conn) {
		logger.Warn("NATS closed")
	}

	nc, err := opts.Connect()
	return nc, err
}
