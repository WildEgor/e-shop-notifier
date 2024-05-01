package domains

import (
	"errors"
)

type EmailNotification struct {
	Email   string `json:"email,omitempty"`
	Subject string `json:"subject,omitempty"`
	Message string `json:"text,omitempty"`
}

func ValidateEmailNotification(d *EmailNotification) error {
	if d.Email == "" {
		return errors.New("email must defined")
	}

	// FIXME: regex not correct for test@mail.ru
	//if len(regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).FindStringSubmatch(d.Email)) < 2 {
	//	return errors.New("email incorrect format")
	//}

	//if len(d.Subject) == 0 || len(d.Message) == 0 {
	//	return errors.New("empty subject or message")
	//}

	return nil
}
