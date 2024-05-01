package configs

import (
	"fmt"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
)

type SMTPConfig struct {
	From     string `env:"SMTP_FROM_EMAIL"`
	Host     string `env:"SMTP_HOST,required"`
	Port     int    `env:"SMTP_PORT,required"`
	Username string `env:"SMTP_USERNAME,required"`
	Password string `env:"SMTP_PASSWORD,required"`
}

func NewSMTPConfig(c *Configurator) *SMTPConfig {
	cfg := SMTPConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("smtp config parse error", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	return &cfg
}

func (c *SMTPConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
