package notify_handler

import (
	sms_adapter "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/adapters/sms"
	smtp_adapter "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/adapters/smtp"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/domains"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/dtos"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/services/template"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/validators"
	"github.com/wagslane/go-rabbitmq"
	"log/slog"
	"time"
)

type NotifyHandler struct {
	smtpa *smtp_adapter.SMTPAdapter
	smsa  *sms_adapter.SMSAdapter
	tmls  *template.TemplateService
}

func NewNotifyHandler(
	smtpa *smtp_adapter.SMTPAdapter,
	smsa *sms_adapter.SMSAdapter,
	tmls *template.TemplateService,
) *NotifyHandler {
	return &NotifyHandler{
		smtpa,
		smsa,
		tmls,
	}
}

func (h *NotifyHandler) Handle(d rabbitmq.Delivery) rabbitmq.Action {
	req := &dtos.NotifierPayloadDto{
		TimeReqStart: time.Now(),
	}

	if err := validators.ParseAndValidateBytes(d.Body, req); err != nil {
		slog.Error("handle validation notify error", err)
		return rabbitmq.Ack
	}

	if req.IsEmail() {
		n := &domains.EmailNotification{
			Email:   req.EmailSetting.Email,
			Subject: req.EmailSetting.Subject,
			Message: req.EmailSetting.Text,
		}

		if err := domains.ValidateEmailNotification(n); err != nil {
			slog.Error(
				"Not valid email email",
				slog.String("error", err.Error()),
			)
			return rabbitmq.Ack
		}

		if req.WithTemplate() {
			msg, err := h.tmls.Build(req.EmailSetting.Template, req.EmailSetting.Data)

			if err != nil {
				slog.Error("template parse error", err)
				return rabbitmq.Ack
			}

			n.Message = msg
		}

		if err := h.smtpa.Send(n); err != nil {
			//slog.Error("failed send to", req.EmailSetting.Email)
		}
	}

	if req.IsSms() {
		n := &domains.SMSNotification{
			Phone:   req.PhoneSetting.Number,
			Message: req.PhoneSetting.Text,
		}

		if err := domains.ValidateSMSNotification(n); err != nil {
			slog.Error(
				"not valid sms sms",
				slog.String("error", err.Error()),
			)
			return rabbitmq.Ack
		}

		if err := h.smsa.Send(n); err != nil {
			//slog.Error("Failed send to", req.PhoneSetting.Number)
		}
	}

	return rabbitmq.Ack
}
