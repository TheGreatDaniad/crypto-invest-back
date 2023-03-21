package gateway

import (
	"github.com/thegreatdaniad/crypto-invest/gateway/models"
	postmanAPI "github.com/thegreatdaniad/crypto-invest/services/postman/API"
	userAPI "github.com/thegreatdaniad/crypto-invest/services/user/API"
)

func GenerateServices() models.Services {
	// initializing services separately
	var services = models.Services{
		UserService:    userAPI.Service{},
		PostmanService: postmanAPI.Service{},
	}

	// injecting services object to each service
	// so they can access other services as well

	services.UserService.SetServices(services)
	services.PostmanService.SetServices(services)

	return services
}
