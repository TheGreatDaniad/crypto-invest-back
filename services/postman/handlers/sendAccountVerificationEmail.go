package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/postman/models"
)

func (h Handler) SendAccountVerificationEmail(c *gatewayModels.Carrier, name string, link string, to string) customErrors.Error {

	data := emailData{
		Name: name,
		Link: link,
	}
	content, err := generateEmailTemplate(data)
	if err != nil {
		fmt.Println(err)
		return customErrors.New(err, customErrors.EmailNotSent)

	}
	email := models.Email{
		Subject:      "Account verification at crypto-invest",
		Content:      content,
		To:           to,
		SenderName:   "crypto-invest",
		ReceiverName: data.Name,
	}
	c.InitializeWithData(email, gatewayModels.UserService, "send account verification email")
	err = h.Email.SendEmail(c, email)
	if err != nil {
		return customErrors.New(err, customErrors.EmailNotSent)
	}

	return customErrors.New(nil, customErrors.NoError)
}

func generateEmailTemplate(ed emailData) (string, error) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "services/postman/handlers/accountVerification.html"))
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, ed); err != nil {
		return "", err
	}
	return buf.String(), nil

}

type emailData struct {
	Name string
	Link string
}
