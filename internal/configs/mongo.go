package configs

import (
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
)

type MongoDBConfig struct {
	URI    string `env:"MONGODB_URI,required"`
	DbName string `env:"MONGODB_NAME,required"`
}

func NewMongoDBConfig(c *Configurator) *MongoDBConfig {
	cfg := MongoDBConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("mongo config parse error", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	return &cfg
}
