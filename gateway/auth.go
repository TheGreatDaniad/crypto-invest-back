package gateway

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	dbConnector "github.com/thegreatdaniad/crypto-invest/connectors/database"
	"github.com/thegreatdaniad/crypto-invest/constants"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/gateway/models"
	userHandlers "github.com/thegreatdaniad/crypto-invest/services/user/handlers"
	"github.com/thegreatdaniad/crypto-invest/services/user/impl"
	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"

	"github.com/thegreatdaniad/crypto-invest/utils"
)

func auth(minRank string, accessByToken bool) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			carrier := e.Get("Carrier")
			c, ok := carrier.(models.Carrier)
			if !ok {
				return customErrors.InternalServerErrorHandler(e, []string{customErrors.InternalServerError})
			}
			c.InitializeWithData(nil, constants.GatewayService, "Authorization")

			token := e.QueryParam("auth")
			var claims utils.Claims
			var err error

			if accessByToken && token != "" {

				claims, err = utils.ParseJwt(token)
				if err != nil {
					c.SetError(err)
					return customErrors.UnauthorizedErrorHandler(e, []string{err.Error()})
				}
				if claims.Purpose != utils.FileOperations {
					return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.Unauthorized})
				}

			} else {
				cookie, err := e.Cookie("Auth")
				if err != nil {
					c.SetError(err)
					return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.Unauthorized})
				}
				jwt := parseCookie(cookie)
				claims, err = utils.ParseJwt(jwt)
				if err != nil {
					c.SetError(err)
					return customErrors.UnauthorizedErrorHandler(e, []string{err.Error()})
				}
			}

			db := impl.Database{
				Postgres: dbConnector.GetDBConnection(),
			}

			h := userHandlers.Handler{
				Db: db,
			}
			u, err := h.Db.GetUserById(&c, claims.Id)
			if err != nil {
				c.SetError(err)
				return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.Unauthorized})
			}
			c.User = u
			if !userModels.IsHigherOrEqualRank(c.User.Rank, minRank) {
				return customErrors.UnauthorizedErrorHandler(e, []string{customErrors.Unauthorized})
			}

			e.Set("Carrier", c)
			c.SetSuccess("Authorized")
			return next(e)
		}
	}
}

func parseCookie(cookie *http.Cookie) string {
	rawCookie := cookie.String()
	token := strings.Split(rawCookie, " ")[1]
	token = token[:len(token)-1]
	return token

}

// func GenerateAuthToken(e echo.Context) error {

// 	c := e.Get("Carrier").(models.Carrier)

// 	c.EndPointEntry(nil, "[GET]/files/token")

// 	token, err := utils.GenerateTimedAuthJwt(c.User.ID, 3)
// 	if err != nil {
// 		return customErrors.InternalServerErrorHandler(e, []string{customErrors.InternalServerError})
// 	}
// 	c.Finilize()
// 	return httpResponses.Success(e, 200, []interface{}{token})
// }

func GenerateAuthToken_(c models.Carrier, userId uint) (string, error) {

	c.EndPointEntry(userId, "[INTERNAL]/auth/token{generateAuthToken}")

	token, err := utils.GenerateTimedAuthJwt(userId, 24*30*3)
	if err != nil {
		return "", err
	}
	c.Finilize()
	return token, nil
}
