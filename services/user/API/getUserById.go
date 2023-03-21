package API

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
)

func (s Service) GetUserById(e echo.Context) error {

	c := e.Get("Carrier").(gatewayModels.Carrier)

	id, err := strconv.Atoi(e.Param("id"))
	userId := uint(id)
	if id == 0 {
		userId = c.User.ID
	}
	if err != nil {
		customErrors.InvalidInputsErrorHandler(e, []string{customErrors.InvalidUserId})
		return nil
	}

	u, ce := Handler.GetUserById(&c, userId)
	u.HidePassword()
	if ce.Err != nil {
		c.SetError(ce.Err)
		return customErrors.HandleCommonErrors(e, ce)
	}

	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{u})

}

func (s Service) GetUserById_(c gatewayModels.Carrier, id uint) (models.User, error) {

	u, ce := Handler.GetUserById(&c, id)
	if ce.Err != nil {
		c.SetError(ce.Err)
		return models.User{}, ce.Err
	}

	return u, nil

}
