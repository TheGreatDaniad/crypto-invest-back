package wallets

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func (s Service) CreateWallet(e echo.Context) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	userID := c.User.ID
	address, err := Handler.Services.BitcoinCoreService.CreateAddress()
	if err != nil {

		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	w := CryptoWallet{
		CustomerID: userID,
		Address:    address,
	}
	wallets, err := db.GetWalletsByCustomerID(&c, userID)
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	if len(wallets) > 0 {
		ce := customErrors.New(errors.New("user already has wallet"), customErrors.BadRequest)
		return customErrors.HandleCommonErrors(e, ce)
	}
	wallet, err := db.CreateWallet(&c, w)
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	c.Finilize()
	return httpResponses.Success(e, 201, []interface{}{wallet})
}

func (s Service) GetWallets(e echo.Context) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	wallets, err := db.GetWallets(&c)
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{wallets})
}

func (s Service) GetWalletsByCustomerID(e echo.Context) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	userID := c.User.ID
	wallets, err := db.GetWalletsByCustomerID(&c, userID)
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	wallet := wallets[0]
	b, err := Handler.Services.BitcoinCoreService.GetAddressBalance(wallet.Address)
	fmt.Println(b, err)

	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	wallet.Balance = b

	t, err := Handler.Services.BitcoinCoreService.GetAddressTransactions(wallet.Address)
	fmt.Println(t, err, "ssss")
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}

	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{wallet})
}

func (s Service) UpdateWallet(e echo.Context) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	w := CryptoWallet{}
	err := json.NewDecoder(e.Request().Body).Decode(&w) // add json data to the wallet struct
	if err != nil {
		fmt.Println(err)
		customErrors.InvalidInputsErrorHandler(e, []string{customErrors.InvalidInputs})
		return nil
	}
	err = w.Validate()
	if err != nil {
		ce := customErrors.New(err, customErrors.InvalidInputs)
		return customErrors.HandleCommonErrors(e, ce)
	}
	err = db.UpdateWallet(&c, w)
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{})
}

func (s Service) DeleteWallet(e echo.Context) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	id, err := strconv.Atoi(e.Param("id"))
	walletID := uint(id)
	if err != nil {
		customErrors.InvalidInputsErrorHandler(e, []string{customErrors.InvalidId})
		return nil
	}
	err = db.DeleteWallet(&c, walletID)
	if err != nil {
		ce := customErrors.New(err, customErrors.InternalServerError)
		return customErrors.HandleCommonErrors(e, ce)
	}
	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{})
}
