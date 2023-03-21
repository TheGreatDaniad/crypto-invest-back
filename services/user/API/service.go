package API

import (
	dbConnector "github.com/thegreatdaniad/crypto-invest/connectors/database"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/handlers"
	"github.com/thegreatdaniad/crypto-invest/services/user/impl"
)

var db = impl.Database{
	Postgres: dbConnector.GetDBConnection(),
}
var Handler = handlers.Handler{
	Db:       db,
	Services: gatewayModels.Services{},
}

type Service struct{}

func (s Service) SetServices(services gatewayModels.Services) {
	Handler.Services = services
}
