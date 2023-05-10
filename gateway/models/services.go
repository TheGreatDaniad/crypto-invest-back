package models

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/labstack/echo/v4"
	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"
)

type IUserService interface {
	GetUserById(e echo.Context) error
	GetUsers(e echo.Context) error
	LoginUser(e echo.Context) error
	Logout(e echo.Context) error
	RegisterUser(e echo.Context) error

	VerifyUser(e echo.Context) error

	UserCanAccessUser_(c *Carrier, targetUserId uint, requesterUserId uint) (bool, error)
	GetUserById_(c Carrier, id uint) (userModels.User, error)
	SetServices(Services)
}

type IPostmanService interface {
	SendAccountVerificationEmail_(c *Carrier, name string, link string, to string) error
	SetServices(Services)
}
type IBitcoinCoreService interface {
	CreateAddress() (string, error)
	GetAddressBalance(address string) (float64, error)
	GetAddressTransactions(addressStr string) ([]btcjson.ListTransactionsResult, error)
	LoadWallet(walletPath string) error
	SetServices(Services)
}

type IWalletService interface {
	CreateWallet(e echo.Context) error
	GetWallets(e echo.Context) error
	GetWalletsByCustomerID(e echo.Context) error
	UpdateWallet(e echo.Context) error
	DeleteWallet(e echo.Context) error
	SetServices(Services)
}
type Services struct {
	UserService        IUserService
	PostmanService     IPostmanService
	BitcoinCoreService IBitcoinCoreService
	WalletService      IWalletService
}
