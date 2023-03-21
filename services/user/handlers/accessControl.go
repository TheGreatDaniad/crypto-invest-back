package handlers

import (
	"errors"

	"github.com/thegreatdaniad/crypto-invest/connectors/database"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func (h Handler) UserCanAccessUser(c *gatewayModels.Carrier, targetUserId uint, requesterUserId uint) (bool, customErrors.Error) {

	tu, err := h.Db.GetUserById(c, targetUserId)
	if err != nil {
		if database.IsRecordNotFound(err) {
			return false, customErrors.New(errors.New(customErrors.UserNotFound), customErrors.NotFound)
		}
		return false, customErrors.New(err, customErrors.DatabaseErrors)
	}
	ru, err := h.Db.GetUserById(c, requesterUserId)
	if err != nil {
		if database.IsRecordNotFound(err) {
			return false, customErrors.New(errors.New(customErrors.UserNotFound), customErrors.NotFound)
		}
		return false, customErrors.New(err, customErrors.DatabaseErrors)
	}

	if !tu.CanBeModifiedBy(ru) {
		return false, customErrors.New(nil, customErrors.NoError)
	}

	return true, customErrors.New(nil, customErrors.NoError)

}
