package handlers

import (
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/postman/models"
)

func (h Handler) SendEmail(c *gatewayModels.Carrier, subject string, content string, to string, senderName string, receiverName string) customErrors.Error {
	email := models.Email{
		Subject:      subject,
		Content:      content,
		To:           to,
		SenderName:   senderName,
		ReceiverName: receiverName,
	}
	c.InitializeWithData(email, gatewayModels.UserService, "send email")
	err := email.Validate()
	if err != nil {
		c.SetError(err)
		return customErrors.New(err, customErrors.InvalidInputs)
	}

	err = h.Email.SendEmail(c, email)
	if err != nil {
		return customErrors.New(err, customErrors.EmailNotSent)
	}

	return customErrors.New(nil, customErrors.NoError)
}
