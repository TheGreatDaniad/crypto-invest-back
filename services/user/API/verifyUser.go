package API

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func (s Service) VerifyUser(e echo.Context) error {

	c := e.Get("Carrier").(gatewayModels.Carrier)

	token := e.QueryParam("token")
	ce := Handler.VerifyUser(&c, token)
	fmt.Println(ce)
	if ce.Err != nil {
		c.SetError(ce.Err)
		if ce.Code == customErrors.CannotParseJwt {
			return customErrors.InvalidInputsErrorHandler(e, []string{customErrors.CannotParseToken})

		}
		if ce.Code == customErrors.UserAlreadyVerified {
			return customErrors.InvalidInputsErrorHandler(e, []string{customErrors.UserAlreadyVerified})

		}
		return customErrors.HandleCommonErrors(e, ce)
	}

	c.Finilize()

	return httpResponses.Success(e, 200, []interface{}{})

}
