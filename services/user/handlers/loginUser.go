package handlers

import (
	"errors"

	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/impl"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
	"gorm.io/gorm"
)

func (h Handler) AuthenticateUser(c *gatewayModels.Carrier, u models.UserLoginInfo) (uint, customErrors.Error) {
	c.InitializeWithData(u, gatewayModels.UserService, "Login the user")
	err := u.Validate()
	if err != nil {
		return 0, customErrors.New(err, customErrors.InvalidInputs)
	}
	user, err := h.Db.GetUserByEmail(c, u.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, customErrors.New(err, customErrors.UserNotFound)
		}
		return 0, customErrors.New(err, customErrors.DatabaseConnectionError)

	}
	passwordIsCorrect, err := impl.CheckPasswordHash(c, u.Password, user.Password)
	if err != nil {
		return 0, customErrors.New(err, customErrors.InternalServerError)
	}
	if passwordIsCorrect {
		return user.ID, customErrors.New(nil, customErrors.NoError)
	}

	err = errors.New("password is not correct")

	return 0, customErrors.New(err, customErrors.PasswordIsNotCorrect)

}
