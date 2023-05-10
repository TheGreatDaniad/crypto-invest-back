package bitcoinCore

import (
	"os"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/joho/godotenv"
	dbConnector "github.com/thegreatdaniad/crypto-invest/connectors/database"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"gorm.io/gorm"
)

type DatabaseActions struct {
	Postgres *gorm.DB
}

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
}

func getBTCCoreClient() (*rpcclient.Client, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	host := os.Getenv("BTC_CORE_HOST")
	user := os.Getenv("BTC_CORE_USER")
	password := os.Getenv("BTC_CORE_PASSWORD")
	var config = &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         password,
		HTTPPostMode: true,
		DisableTLS:   true,
		
		
	}
	return rpcclient.New(config, nil)

}
