package API

import (
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func (s Service) SendAccountVerificationEmail_(c *gatewayModels.Carrier, name string, link string, to string) error {
	ce := Handler.SendAccountVerificationEmail(c, name, link, to)
	if ce.Err != nil {
		return ce.Err
	}
	c.Finilize()
	return nil
}
