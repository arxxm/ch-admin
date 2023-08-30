package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/config"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/repository"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/domain"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/domain/task"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/logger"

	"time"

	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type Consumer struct {
	cfg            *config.Config
	nc             *nats.Conn
	repo           *repository.Repository
	delayedChannel chan domain.NatsMessage
}

func New(cfg *config.Config, nc *nats.Conn, repo *repository.Repository) *Consumer {
	return &Consumer{
		cfg:  cfg,
		nc:   nc,
		repo: repo,
	}
}

func (c *Consumer) resendToNats(message domain.NatsMessage) {
	time.AfterFunc(5*time.Second, func() {
		c.delayedChannel <- message
	})
}

func (c *Consumer) ListenCreateSMS(ctx context.Context) error {
	subject := c.cfg.Nats.SmsSubjectCreate
	sem := semaphore.NewWeighted(c.cfg.Nats.SmsMaxWorkers) // limit concurrent tasks
	errListenGroup, errListenGroupCtx := errgroup.WithContext(ctx)

	subscribe, err := c.nc.Subscribe(subject, func(m *nats.Msg) {
		err := sem.Acquire(errListenGroupCtx, int64(1))
		if err != nil {
			logger.Error(fmt.Sprintf("ListenCreateSMS acquire err: %v", err))
			return
		}

		var sms task.Sms
		if err = json.Unmarshal(m.Data, &sms); err != nil {
			logger.Error(fmt.Sprintf("Unmarshal err: %v, data: %v", err, m.Data))
			return
		}

		errListenGroup.Go(func() error {
			defer sem.Release(int64(1))

			if err := c.repo.Clickhousebulk.InsertSmsTask(ctx, sms); err != nil {
				logger.Error(fmt.Sprintf("the creating of the sms occurred with an error, err: %v", err))
			}

			return nil
		})
	})

	if err != nil {
		logger.Error(fmt.Sprintf("ListenCreateSMS subscribe err: %v", err))
		return err
	}

	logger.Info("ListenCreateSMS subscribed")

	// wait program shutdown
	<-ctx.Done()
	err = subscribe.Unsubscribe()
	if err != nil {
		logger.Error(fmt.Sprintf("ListenCreateSMS unsubscribe err: %v", err))
		return err
	}

	// wait all ListenCreateSMS goroutines shutdown
	if err = errListenGroup.Wait(); err != nil {
		return err
	}
	logger.Info("ListenCreateSMS unsubscribed")

	return nil
}
