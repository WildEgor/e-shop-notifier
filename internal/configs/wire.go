package configs

import "github.com/google/wire"

var ConfigsSet = wire.NewSet(
	NewAppConfig,
	NewConfigurator,
)
