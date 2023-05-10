package database

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"
	"gorm.io/gorm"
)

type CryptoWallet struct {
	ID         uint             `gorm:"primary_key"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"-"`
	DeletedAt  gorm.DeletedAt   `json:"-" sql:"index"`
	Name       *string          `json:"name"`
	CustomerID uint             `json:"customer_id"`
	Customer   *userModels.User `json:"customer" gorm:"foreignKey:CustomerID"`
	Address  string           `json:"address"`
	PrivateKey string           `json:"-"`
	Balance    float64          `json:"balance"`

}

func (w CryptoWallet) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&w.CustomerID, validation.Required),
		validation.Field(&w.Address, validation.Required),
		validation.Field(&w.Balance, validation.Min(0)),

	)
}
