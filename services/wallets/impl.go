package wallets

import (
	"errors"

	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"gorm.io/gorm"
)

type DatabaseActions struct {
	Postgres *gorm.DB
}

func (db DatabaseActions) CreateWallet(c *gatewayModels.Carrier, w CryptoWallet) (CryptoWallet, error) {
	select {
	case <-c.Context.Done():
		err := errors.New("inserting the wallet to the database has been canceled from a parent process")
		return CryptoWallet{}, err
	default:
		res := db.Postgres.Create(&w)
		return w, res.Error
	}
}


func (db DatabaseActions) GetWallets(c *gatewayModels.Carrier) ([]CryptoWallet, error) {
	select {
	case <-c.Context.Done():
		err := errors.New("retrieving the wallet from the database has been canceled from a parent process")
		return []CryptoWallet{}, err
	default:
		w := []CryptoWallet{}
		res := db.Postgres.Find(&w)
		return w, res.Error
	}
}

func (db DatabaseActions) GetWalletsByCustomerID(c *gatewayModels.Carrier, id uint) ([]CryptoWallet, error) {
	select {
	case <-c.Context.Done():
		err := errors.New("retrieving the wallets from the database has been canceled from a parent process")
		return []CryptoWallet{}, err
	default:
		wallets := []CryptoWallet{}
		res := db.Postgres.Find(&wallets, "customer_id = ?", id)
		return wallets, res.Error
	}
}

func (db DatabaseActions) UpdateWallet(c *gatewayModels.Carrier, w CryptoWallet) error {
	select {
	case <-c.Context.Done():
		err := errors.New("updating the wallet in the database has been canceled from a parent process")
		return err
	default:
		res := db.Postgres.Save(&w)
		return res.Error
	}
}

func (db DatabaseActions) DeleteWallet(c *gatewayModels.Carrier, id uint) error {
	select {
	case <-c.Context.Done():
		err := errors.New("deleting the wallet from the database has been canceled from a parent process")
		return err
	default:
		res := db.Postgres.Delete(&CryptoWallet{}, id)
		return res.Error
	}
}
