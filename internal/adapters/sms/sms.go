package sms_adapter

import (
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/configs"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/domains"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"log/slog"
	"strings"
)

type SMSAdapter struct {
	cfg *configs.SMSConfig
}

func NewSMSAdapter(
	cfg *configs.SMSConfig,
) *SMSAdapter {
	return &SMSAdapter{
		cfg,
	}
}

// Send implement own logic here
func (s *SMSAdapter) Send(sms *domains.SMSNotification) (err error) {
	//// HINT: some mock sms payload
	//queryParams := url.Values{
	//	"action":      {"sendmessage"},
	//	"username":    {s.cfg.Username},
	//	"password":    {s.cfg.Password},
	//	"recipient":   {sms.Phone},
	//	"messagetype": {"SMS:TEXT"},
	//	"originator":  {""},
	//	"messagedata": {sms.Message},
	//}
	//
	//requestUrl := fmt.Sprintf("%v?%v", url.URL{
	//	Scheme: "https",
	//	Host:   s.cfg.BaseURL,
	//}, queryParams.Encode())
	//
	//_, err = http.NewRequest(http.MethodGet, requestUrl, nil)
	//if err != nil {
	//	slog.Error(
	//		"creating the request failed",
	//		slog.String("error", err.Error()),
	//	)
	//	return nil
	//}

	plen := len(sms.Phone)
	if plen > 4 {
		mask := strings.Repeat("*", plen-4) + sms.Phone[plen-4:]
		msp := strings.Replace(mask, "*", "*", plen-4)
		slog.Debug(
			"send sms to",
			models.LogEntryAttr(&models.LogEntry{
				Props: map[string]interface{}{
					"to":   msp,
					"text": sms.Message,
				},
			}),
		)
		return nil
	}

	slog.Debug(
		"send sms to",
		models.LogEntryAttr(&models.LogEntry{
			Props: map[string]interface{}{
				"to":   sms.Phone,
				"text": sms.Message,
			},
		}),
	)

	return nil
}
