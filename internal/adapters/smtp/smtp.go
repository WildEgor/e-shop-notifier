package smtp_adapter

import (
	"encoding/base64"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/configs"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/domains"
	"log/slog"
	"net/smtp"
	"strings"
)

type SMTPAdapter struct {
	config *configs.SMTPConfig
}

func NewSMTPAdapter(
	config *configs.SMTPConfig,
) *SMTPAdapter {
	return &SMTPAdapter{
		config,
	}
}

// Send send single email
func (s *SMTPAdapter) Send(email *domains.EmailNotification) (err error) {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	c, err := smtp.Dial(s.config.Addr())
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer c.Close()

	if err = c.Mail(r.Replace(s.config.From)); err != nil {
		slog.Error(err.Error())
		return err
	}

	to := []string{email.Email}

	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + s.config.From + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(email.Message))

	_, err = w.Write([]byte(msg))
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = w.Close()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return c.Quit()
}
