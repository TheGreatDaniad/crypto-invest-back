package API

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
	"github.com/thegreatdaniad/crypto-invest/utils"
)

func (s Service) LoginUser(e echo.Context) error {

	c := e.Get("Carrier").(gatewayModels.Carrier)
	u := models.UserLoginInfo{}
	err := json.NewDecoder(e.Request().Body).Decode(&u) // add json data to the user struct
	if err != nil {
		customErrors.InvalidInputsErrorHandler(e, []string{customErrors.InvalidInputs})
		return nil
	}
	id, ce := Handler.AuthenticateUser(&c, u)
	if ce.Err != nil {
		c.SetError(ce.Err)
		if ce.Code == customErrors.PasswordIsNotCorrect {
			return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.InvalidCredentials})
		}
		if ce.Code == customErrors.UserNotFound {
			return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.InvalidCredentials})
		}
		if ce.Code == customErrors.AuthenticationFailed {
			return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.InvalidCredentials})

		}
		return customErrors.HandleCommonErrors(e, ce)
	}
	cookie, err := GenerateAuthCookie(id, e)
	if err != nil {
		c.SetError(ce.Err)
		customErrors.InternalServerErrorHandler(e, []string{customErrors.InternalServerError})
		return nil
	}
	e.SetCookie(cookie)

	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{})

}

func GenerateAuthCookie(id uint, e echo.Context) (*http.Cookie, error) {
	token, err := utils.GenerateAuthJwt(id)
	if err != nil {
		return &http.Cookie{}, err
	}

	cookie := new(http.Cookie)
	cookie.Name = "Auth"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.Value = "Bearer " + token
	cookie.Expires = time.Now().Add(10 * 12 * 30 * 24 * time.Hour) //expires in 10 years
	return cookie, nil
}

func GenerateEmptyAuthCookie(e echo.Context) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "Auth"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.Value = ""
	cookie.MaxAge = -1

	return cookie
}
