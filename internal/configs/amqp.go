package configs

import (
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
)

type AMQPConfig struct {
	URI string `env:"AMQP_URI,required"`

	Queue    string `env:"NOTIFIER_QUEUE,required"`
	Exchange string `env:"NOTIFIER_EXCHANGE,required"`
}

func NewAMQPConfig(
	c *Configurator,
) *AMQPConfig {
	cfg := AMQPConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("amqp config parse error", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	return &cfg
}
