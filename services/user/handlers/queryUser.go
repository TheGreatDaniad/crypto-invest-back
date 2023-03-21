package handlers

import (
	"errors"

	"github.com/thegreatdaniad/crypto-invest/connectors/database"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
)

func (h Handler) GetUserById(c *gatewayModels.Carrier, id uint) (models.User, customErrors.Error) {
	c.InitializeWithData(id, gatewayModels.UserService, "get user by id")

	u, err := h.Db.GetUserById(c, id)

	if err != nil {
		if database.IsRecordNotFound(err) {
			return models.User{}, customErrors.New(errors.New(customErrors.UserNotFound), customErrors.NotFound) //TODO - check error

		}
		return models.User{}, customErrors.New(err, customErrors.DatabaseErrors)
	}
	if !u.CanBeModifiedBy(c.User) {
		return models.User{}, customErrors.New(errors.New(customErrors.Unauthorized), customErrors.Unauthorized)
	}
	return u, customErrors.New(nil, customErrors.NoError)
}

func (h Handler) GetUsers(c *gatewayModels.Carrier, queryParams models.UserQueryInfo) ([]models.User, customErrors.Error) {
	c.InitializeWithData(queryParams, gatewayModels.UserService, "get users")
	err := queryParams.Validate()
	if err != nil {
		return []models.User{}, customErrors.New(err, customErrors.InvalidInputs)
	}

	users, err := h.Db.GetUsers(c, queryParams)
	if err != nil {
		return []models.User{}, customErrors.New(err, customErrors.UserNotFound) //TODO - check error
	}
	var resultingUsers []models.User
	for _, user := range users {
		user.HidePassword()
		if user.CanBeModifiedBy(c.User) {
			resultingUsers = append(resultingUsers, user)
		}
	}
	return resultingUsers, customErrors.New(nil, customErrors.NoError)

}
