package main

import (
	"github.com/thegreatdaniad/crypto-invest/connectors/database"
	"github.com/thegreatdaniad/crypto-invest/gateway"
)

func main() {
	database.Migrate()
	gateway.Start()
}
