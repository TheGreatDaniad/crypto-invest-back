package handlers

import (
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
)

type Handler struct {
	Db       Database
	Services gatewayModels.Services
}

type Database interface {
	CreateUser(c *gatewayModels.Carrier, user models.UserRegisterInfo) (models.User, error)
	UserDoesNotExist(c *gatewayModels.Carrier, user models.UserRegisterInfo) (bool, error)
	GetUserByEmail(c *gatewayModels.Carrier, email string) (models.User, error)
	GetUserById(c *gatewayModels.Carrier, id uint) (models.User, error)
	GetUsers(c *gatewayModels.Carrier, queryParams models.UserQueryInfo) ([]models.User, error)
	DeleteUser(c *gatewayModels.Carrier, id uint) error
	UpdateUser(c *gatewayModels.Carrier, id uint, u models.User) error
	UpdateUserPassword(c *gatewayModels.Carrier, id uint, u models.User) error
	UpdateProfileImage(c *gatewayModels.Carrier, id uint, imagePath string) error
}
