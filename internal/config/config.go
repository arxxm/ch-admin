package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	CHBulk ClickhouseBulk `mapstructure:"clickhousebulk"`
	Nats   Nats           `mapstructure:"nats"`
	Ctx    struct {
		Ttl time.Duration `mapstructure:"ttl"`
	} `mapstructure:"ctx"`
}

type ClickhouseBulk struct {
	Url        string
	BufferTime time.Duration `mapstructure:"buffer_time"`
}

type Nats struct {
	Address          string
	SmsSubjectCreate string
	SmsMaxWorkers    int64
}

func New(folder, filename string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("clickhouse_bulk", &cfg.CHBulk); err != nil {
		return nil, err
	}

	if err := envconfig.Process("nats", &cfg.Nats); err != nil {
		return nil, err
	}

	return cfg, nil
}
