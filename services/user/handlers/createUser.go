package handlers

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"

	"github.com/thegreatdaniad/crypto-invest/services/user/impl"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
	"github.com/thegreatdaniad/crypto-invest/utils"
)

func (h Handler) CreateUser(c *gatewayModels.Carrier, u models.UserRegisterInfo) (models.User, customErrors.Error) {
	c.InitializeWithData(u, gatewayModels.UserService, "Create the user")
	err := u.Validate()
	if err != nil {
		return models.User{}, customErrors.New(err, customErrors.InvalidInputs)
	}
	userExists, err := h.Db.UserDoesNotExist(c, u)
	if err != nil {
		return models.User{}, customErrors.New(err, customErrors.InternalServerError)
	}
	if userExists {
		return models.User{}, customErrors.New(errors.New(customErrors.UserAlreadyExists), customErrors.InvalidInputs)
	}
	hashedPassword, err := impl.HashPassword(c, u.Password)
	if err != nil {
		return models.User{}, customErrors.New(err, customErrors.InternalServerError)
	}
	u.Password = hashedPassword
	newUser, err := h.Db.CreateUser(c, u)
	if err != nil {
		return models.User{}, customErrors.New(err, customErrors.InternalServerError)
	}
	vl, err := generateVerificationLink(newUser)
	if err != nil {
		return models.User{}, customErrors.New(err, customErrors.InternalServerError)
	}
	//TODO: change the hard coded user sting here
	h.Services.PostmanService.SendAccountVerificationEmail_(c, "user", vl, u.Email)

	return newUser, customErrors.New(nil, customErrors.NoError)

}

func generateVerificationLink(u models.User) (string, error) {
	token, err := utils.GenerateUserVerificationJwt(u.ID)
	if err != nil {
		return "", err
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return "", err
	}
	host := os.Getenv("CLIENT_BASE_URL")

	link := host + "/user/verification?token=" + token
	return link, nil
}
