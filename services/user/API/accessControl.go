package API

import (
	"errors"

	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func (s Service) UserCanAccessUser_(c *gatewayModels.Carrier, targetUserId uint, requesterUserId uint) (bool, error) {

	userCanAccess, ce := Handler.UserCanAccessUser(c, targetUserId, requesterUserId)
	if ce.Err != nil {
		return false, errors.New(ce.Code)
	}

	c.Finilize()
	return userCanAccess, nil
}
