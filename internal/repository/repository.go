package repository

import (
	"pingocean-front/clickhouse-admin-sms-consumer/internal/config"
	clickhousebulkrepo "pingocean-front/clickhouse-admin-sms-consumer/internal/repository/clickhousebulk"
)

type Repository struct {
	clickhousebulkrepo.Clickhousebulk
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		Clickhousebulk: clickhousebulkrepo.NewClickhousebulkRepository(cfg),
	}
}
