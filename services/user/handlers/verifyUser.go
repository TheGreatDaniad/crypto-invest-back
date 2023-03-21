package handlers

import (
	"errors"

	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
	"github.com/thegreatdaniad/crypto-invest/utils"
)

func (h Handler) VerifyUser(c *gatewayModels.Carrier, token string) customErrors.Error {
	c.InitializeWithData(token, gatewayModels.UserService, "user verification")
	claims, err := utils.ParseJwt(token)
	if err != nil {
		return customErrors.New(err, customErrors.CannotParseJwt)
	}

	verifiedUser := models.User{
		Rank: models.StandardUser,
	}
	u, err := h.Db.GetUserById(c, claims.Id)

	if err != nil {
		return customErrors.New(err, customErrors.InternalServerError)
	}
	if u.Rank == models.StandardUser {
		return customErrors.New(errors.New(customErrors.UserAlreadyVerified), customErrors.UserAlreadyVerified)
	}
	err = h.Db.UpdateUser(c, claims.Id, verifiedUser)
	if err != nil {
		return customErrors.New(err, customErrors.InternalServerError)
	}
	return customErrors.New(nil, customErrors.NoError)
}
