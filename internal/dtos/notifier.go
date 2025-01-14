package dtos

import (
	"errors"
	"time"
)

type NotifierResendRequestDto struct {
	Req     NotifierPayloadDto `json:"request"`
	Error   string             `json:"error"`
	TimeReq string             `json:"time_req"`
}

type NotifierEmailSettings struct {
	Email    string      `json:"email"`
	Subject  string      `json:"subject"`
	Template string      `json:"template,omitempty"`
	Text     string      `json:"text,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type NotifierPayloadDto struct {
	Type         string                `json:"type"`
	EmailSetting NotifierEmailSettings `json:"email_setting,omitempty"`
	PhoneSetting struct {
		Number string `json:"phone"`
		Text   string `json:"text"`
	} `json:"phone_setting,omitempty"`
	PushSetting struct {
		To       string      `json:"to"`
		Platform string      `json:"platform"`
		Image    string      `json:"image,omitempty"`
		Title    string      `json:"title,omitempty"`
		Message  string      `json:"message,omitempty"`
		Template string      `json:"template,omitempty"`
		Data     interface{} `json:"data,omitempty"`
	} `json:"push_settings,omitempty"`
	Data         interface{} `json:"data"`
	Error        error       `json:"-"`
	TimeReqStart time.Time   `json:"-"`
}

func (r *NotifierPayloadDto) IsSms() bool {
	return r.Type == "sms"
}

func (r *NotifierPayloadDto) IsEmail() bool {
	return r.Type == "email"
}

func (r *NotifierPayloadDto) IsPush() bool {
	return r.Type == "push"
}

func (r *NotifierPayloadDto) IsForAndroid() bool {
	return r.Type == "push" && r.PushSetting.Platform == "ANDROID"
}

func (r *NotifierPayloadDto) IsForIOS() bool {
	return r.Type == "push" && r.PushSetting.Platform == "IOS"
}

func (r *NotifierPayloadDto) WithTemplate() bool {
	return len(r.EmailSetting.Template) != 0 || len(r.PushSetting.Template) != 0
}

// TODO: use validators
// Validate old fashion validation
func (r *NotifierPayloadDto) Validate() bool {
	if r.Type != "sms" && r.Type != "email" && r.Type != "push" {
		r.Error = errors.New("incorrect type")
		return false
	}

	if r.IsSms() {
		return r.ValidateSms()
	}

	if r.IsEmail() {
		return r.ValidateEmail()
	}

	if r.IsPush() {
		return r.ValidatePush()
	}

	return false
}

func (r *NotifierPayloadDto) ValidateEmail() bool {
	if len(r.EmailSetting.Email) == 0 {
		r.Error = errors.New("error email param")
		return false
	}
	if len(r.EmailSetting.Subject) == 0 {
		r.Error = errors.New("error pass subject param")
		return false
	}
	if len(r.EmailSetting.Template) == 0 {
		r.Error = errors.New("error template param")
		return false
	}

	return true
}

func (r *NotifierPayloadDto) ValidateSms() bool {
	if len(r.PhoneSetting.Number) == 0 || len(r.PhoneSetting.Text) == 0 {
		r.Error = errors.New("error sms params")
		return false
	}
	return true
}

func (r *NotifierPayloadDto) ValidatePush() bool {
	if len(r.PushSetting.To) == 0 {
		r.Error = errors.New("error push to params")
		return false
	}

	return true
}

func (r *NotifierPayloadDto) HasError() bool {
	return r.Error != nil
}
