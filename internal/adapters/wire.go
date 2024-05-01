package adapters

import (
	sms_adapter "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/adapters/sms"
	smtp_adapter "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/adapters/smtp"
	"github.com/google/wire"
)

var AdaptersSet = wire.NewSet(
	sms_adapter.NewSMSAdapter,
	smtp_adapter.NewSMTPAdapter,
)
