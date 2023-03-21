package API

import (
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/postman/handlers"
	"github.com/thegreatdaniad/crypto-invest/services/postman/impl"
)

var Handler = handlers.Handler{
	Email:    impl.EmailAgent{},
	Services: gatewayModels.Services{},
}

type Service struct{}

func (s Service) SetServices(services gatewayModels.Services) {
	Handler.Services = services
}
