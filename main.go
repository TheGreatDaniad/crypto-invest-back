package main

import (
	"github.com/thegreatdaniad/crypto-invest/connectors/database"
	"github.com/thegreatdaniad/crypto-invest/gateway"
	"github.com/thegreatdaniad/crypto-invest/services/bitcoinCore"
)

func main() {
	database.Migrate()
	b := bitcoinCore.Service{}
	b.LoadWallet("wal")
	gateway.Start()

}
