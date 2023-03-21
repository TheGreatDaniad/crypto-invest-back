package models

import (
	"github.com/labstack/echo/v4"
	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"
)

type IUserService interface {
	GetUserById(e echo.Context) error
	GetUsers(e echo.Context) error
	LoginUser(e echo.Context) error
	Logout(e echo.Context) error
	RegisterUser(e echo.Context) error

	VerifyUser(e echo.Context) error

	UserCanAccessUser_(c *Carrier, targetUserId uint, requesterUserId uint) (bool, error)
	GetUserById_(c Carrier, id uint) (userModels.User, error)
	SetServices(Services)
}


type IPostmanService interface {
	SendAccountVerificationEmail_(c *Carrier, name string, link string, to string) error
	SetServices(Services)
}

type Services struct {
	UserService    IUserService
	PostmanService IPostmanService
}
