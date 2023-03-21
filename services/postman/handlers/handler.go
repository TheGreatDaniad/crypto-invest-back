package handlers

import (
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/postman/models"
)

type Handler struct {
	Email    EmailAgent
	Sms      SmsAgent
	Services gatewayModels.Services
}

type EmailAgent interface {
	SendEmail(*gatewayModels.Carrier, models.Email) error
}
type SmsAgent interface {
	SendSms(models.Sms) error
}
