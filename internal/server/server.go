package server

import (
	"context"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/config"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/repository"
	"pingocean-front/clickhouse-admin-sms-consumer/internal/server/consumer"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/logger"
	"sync"

	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	cons     *consumer.Consumer
	errGroup *errgroup.Group
	repo     *repository.Repository
	cfg      *config.Config
	wg       *sync.WaitGroup
	nc       *nats.Conn
}

func NewServer(cons *consumer.Consumer, errGroup *errgroup.Group, repo *repository.Repository, nc *nats.Conn, cfg *config.Config) *Server {
	return &Server{
		cons:     cons,
		errGroup: errGroup,
		repo:     repo,
		cfg:      cfg,
		wg:       &sync.WaitGroup{},
		nc:       nc,
	}
}

func (s *Server) ListenAndServe(ctx context.Context) {

	s.errGroup.Go(func() error {
		return s.cons.ListenCreateSMS(ctx)
	})

	logger.Info("all consumers have been started")
}

func (s *Server) GracefulStop() error {
	s.wg.Wait()
	return s.errGroup.Wait()
}
