package impl

import (
	"errors"

	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(c *gatewayModels.Carrier, password string) (string, error) {

	select {
	case <-c.Context.Done():
		err := errors.New("hashing operation canceled from a parent process")
		c.SetError(err)
		return "", err
	default:
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err

	}
}

func CheckPasswordHash(c *gatewayModels.Carrier, password, hash string) (bool, error) {

	select {
	case <-c.Context.Done():
		err := errors.New("checking password operation canceled from a parent process")
		c.SetError(err)
		return false, err
	default:
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		return err == nil, nil
	}
}
