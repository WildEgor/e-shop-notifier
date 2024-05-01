package services

import (
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/services/template"
	"github.com/google/wire"
)

var ServicesSet = wire.NewSet(
	template.NewTemplateService,
)
