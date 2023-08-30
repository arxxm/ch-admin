package clickhousebulk

import "pingocean-front/clickhouse-admin-sms-consumer/internal/config"

type Clickhousebulk interface {
	Sms
}

type ClickhousebulkRepository struct {
	cfg *config.Config
}

func NewClickhousebulkRepository(cfg *config.Config) *ClickhousebulkRepository {
	return &ClickhousebulkRepository{cfg: cfg}
}
