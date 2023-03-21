package API

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
)

func (s Service) RegisterUser(e echo.Context) error {

	c := e.Get("Carrier").(gatewayModels.Carrier)
	u := models.UserRegisterInfo{}
	err := json.NewDecoder(e.Request().Body).Decode(&u) // add json data to the user struct
	if err != nil {
		customErrors.InvalidInputsErrorHandler(e, []string{customErrors.InvalidInputs})
		return nil
	}
	user, ce := Handler.CreateUser(&c, u)
	if ce.Err != nil {
		c.SetError(ce.Err)
		if ce.Code == customErrors.UserAlreadyExists {
			return customErrors.InvalidInputsErrorHandler(e, []string{customErrors.UserAlreadyExists})
		}
		return customErrors.HandleCommonErrors(e, ce)
	}
	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{user})

}
