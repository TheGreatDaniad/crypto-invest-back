package gateway

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/gateway/models"
	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Wellcome to crypto-invest api!\nyou may find the api docs at: https://crypto-invest.stoplight.io/docs/crypto-invest")
}
func logReq(c echo.Context) error {
	value, err := c.FormFile("files[0]")
	if err != nil {
		fmt.Println(err)
	}
	println(value.Filename)
	return c.String(http.StatusOK, "Wellcome to crypto-invest api!\nyou may find the api docs at: https://crypto-invest.stoplight.io/docs/crypto-invest")
}

var services models.Services = GenerateServices()

func addRoutes(gw models.Gateway) {
	gw.Handler.GET("/", hello)
	gw.Handler.POST("/", logReq)

	//users paths
	gw.Handler.GET("/users/:id", services.UserService.GetUserById, auth(userModels.NotVerifiedUser, false))

	gw.Handler.GET("/users", services.UserService.GetUsers, auth(userModels.AdminUser, false))
	gw.Handler.POST("/users", services.UserService.RegisterUser)

	//auth paths
	gw.Handler.POST("/login", services.UserService.LoginUser)
	gw.Handler.POST("/logout", services.UserService.Logout)

	gw.Handler.POST("/users/verification", services.UserService.VerifyUser)

	//wallet paths
	gw.Handler.GET("/wallets/0", services.WalletService.GetWalletsByCustomerID, auth(userModels.DisabledUser, false))
	gw.Handler.GET("/wallets", services.WalletService.GetWallets, auth(userModels.AdminUser, false))
	gw.Handler.POST("/wallets", services.WalletService.CreateWallet, auth(userModels.StandardUser, false))
	gw.Handler.PUT("/wallets/:id", services.WalletService.UpdateWallet, auth(userModels.AdminUser, false))
	gw.Handler.DELETE("/wallets/:id", services.WalletService.DeleteWallet, auth(userModels.StandardUser, false))

}
