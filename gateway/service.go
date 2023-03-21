package gateway

import "github.com/thegreatdaniad/crypto-invest/gateway/models"

func Start() {
	GateWay := models.NewGateway()
	handleMiddlewares(GateWay)
	addRoutes(GateWay)
	// Start server
	GateWay.Handler.Logger.Fatal(GateWay.Handler.Start(":3001"))
}
