package configs

import (
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
)

type SMSConfig struct {
	BaseURL  string `env:"SMS_BASE_URL"`
	Username string `env:"SMS_USERNAME"`
	Password string `env:"SMS_PASSWORD"`
}

func NewSMSConfig(c *Configurator) *SMSConfig {
	cfg := SMSConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("sms config parse error", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	return &cfg
}
