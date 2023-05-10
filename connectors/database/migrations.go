package database

import (
	"log"

	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"
)

func Migrate() {
	if !Connected {
		ConnectToDatabase()
	}

	err := Pg.AutoMigrate(&userModels.User{}, &CryptoWallet{})

	if err != nil {
		log.Fatal(err)
	}

}
