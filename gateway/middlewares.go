package gateway

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func handleMiddlewares(gw models.Gateway) {
	env := os.Getenv("ENV")
	middlewares := []echo.MiddlewareFunc{
		middleware.Secure(),
		middleware.Recover(),
		middleware.Logger(),

		appendCarrier(),
	}
	clientBaseUrl := os.Getenv("CLIENT_BASE_URL")
	if env == "DEV" {
		middlewares = append(middlewares, (middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{clientBaseUrl, "https://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},

			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials},
			AllowCredentials: true,
		})))
	}
	if env == "PROD" {
		middlewares = append(middlewares, (middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{clientBaseUrl},
			AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials},
			AllowCredentials: true,
		})))
	}
	gw.AddMiddlewares(middlewares)

}

func appendCarrier() echo.MiddlewareFunc {
	ctx := context.Background()
	carrier := models.Carrier{
		Context: ctx,
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("Carrier", carrier)
			return next(c)
		}
	}
}

// func prelog() echo.MiddlewareFunc {

// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			fmt.Println("incoming request from " + c.RealIP() + " to " + c.Request().URL.Path)
// 			return next(c)
// 		}
// 	}
// }
