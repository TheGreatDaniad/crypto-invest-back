package wallets

import (
	dbConnector "github.com/thegreatdaniad/crypto-invest/connectors/database"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

var db = DatabaseActions{
	Postgres: dbConnector.GetDBConnection(),
}

type HandlerType struct {
	Db       Database
	Services gatewayModels.Services
}

var Handler = HandlerType{
	Db:       db,
	Services: gatewayModels.Services{},
}

type Service struct{}

func (s Service) SetServices(services gatewayModels.Services) {
	Handler.Services = services
}

type Database interface {
	CreateWallet(c *gatewayModels.Carrier, wallet CryptoWallet) (CryptoWallet, error)
	GetWallets(c *gatewayModels.Carrier) ([]CryptoWallet, error)
	GetWalletsByCustomerID(c *gatewayModels.Carrier, id uint) ([]CryptoWallet, error)
	UpdateWallet(c *gatewayModels.Carrier, wallet CryptoWallet) error
	DeleteWallet(c *gatewayModels.Carrier, id uint) error
}
