package models

import (
	"fmt"
	"log"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/joho/godotenv"
)

type Email struct {
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	To           string `json:"to"`
	SenderName   string `json:"senderName"`
	ReceiverName string `json:"reiceverName"`
}

func (e Email) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Subject, validation.Required, validation.Length(1, 200)),
		validation.Field(&e.SenderName, validation.Length(1, 100)),
		validation.Field(&e.SenderName, validation.Length(1, 100)),
		validation.Field(&e.SenderName, validation.Required, is.Email),
	)
}

func (e Email) GetFullSender() string {

	return fmt.Sprintf("%v <%v>", e.SenderName, getSenderAddress())
}

func getSenderAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("EMAIL_SENDER_ADDRESS")

}
