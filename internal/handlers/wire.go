package handlers

import (
	eh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/errors"
	hch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/health_check"
	notify_handler "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/notify"
	rch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ready_check"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/services"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	services.ServicesSet,
	eh.NewErrorsHandler,
	hch.NewHealthCheckHandler,
	rch.NewReadyCheckHandler,
	notify_handler.NewNotifyHandler,
)
