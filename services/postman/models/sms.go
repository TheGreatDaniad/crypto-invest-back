package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Sms struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	To           string `json:"to"`
	SenderName   string `json:"senderName"`
	ReceiverName string `json:"reiceverName"`
}

func (s Sms) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Title, validation.Required, validation.Length(1, 200)),
		validation.Field(&s.SenderName, validation.Length(1, 100)),
		validation.Field(&s.SenderName, validation.Length(1, 100)),
		validation.Field(&s.SenderName, validation.Required, is.Email),
	)
}
