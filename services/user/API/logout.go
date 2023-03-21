package API

import (
	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func (s Service) Logout(e echo.Context) error {

	c := e.Get("Carrier").(gatewayModels.Carrier)

	c.EndPointEntry(nil, "[GET]/logout")
	ec := GenerateEmptyAuthCookie(e)
	e.SetCookie(ec)
	c.Finilize()

	return httpResponses.Success(e, 200, []interface{}{})

}
