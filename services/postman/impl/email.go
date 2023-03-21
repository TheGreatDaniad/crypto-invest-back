package impl

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/mail.v2"

	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/postman/models"
)

type EmailAgent struct{}

func (ea EmailAgent) SendEmail(c *gatewayModels.Carrier, email models.Email) error {

	select {
	case <-c.Context.Done():
		err := errors.New("sending email operation canceled from a parent process")
		c.SetError(err)
		return nil

	default:

		SMTP_HOST := os.Getenv("SMTP_HOST")
		SMTP_PORT := os.Getenv("SMTP_PORT")
		SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
		SMTP_EMAIL := os.Getenv("SMTP_EMAIL")

		from := fmt.Sprintf("crypto-invest <%v>", SMTP_EMAIL)
		port, err := strconv.Atoi(SMTP_PORT)
		if err != nil {
			fmt.Println(err)
		}
		// auth := smtp.PlainAuth("", SMTP_EMAIL, SMTP_PASSWORD, SMTP_HOST)
		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", email.To)
		m.SetHeader("Subject", email.Subject)
		m.SetBody("text/html", email.Content)

		d := gomail.NewDialer(SMTP_HOST, port, SMTP_EMAIL, SMTP_PASSWORD)

		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("Email Sent!")
		return nil
	}

}
