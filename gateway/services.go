package gateway

import (
	"github.com/thegreatdaniad/crypto-invest/gateway/models"
	postmanAPI "github.com/thegreatdaniad/crypto-invest/services/postman/API"
	bitcoinCore "github.com/thegreatdaniad/crypto-invest/services/bitcoinCore"
	wallets "github.com/thegreatdaniad/crypto-invest/services/wallets"

	userAPI "github.com/thegreatdaniad/crypto-invest/services/user/API"
)

func GenerateServices() models.Services {
	// initializing services separately
	var services = models.Services{
		UserService:    userAPI.Service{},
		PostmanService: postmanAPI.Service{},
		BitcoinCoreService:    bitcoinCore.Service{},
		WalletService: wallets.Service{},
    }

	

	// injecting services object to each service
	// so they can access other services as well

	services.UserService.SetServices(services)
	services.PostmanService.SetServices(services)
	services.WalletService.SetServices(services)

	return services
}
