package API

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	"github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
)

func (s Service) GetUsers(e echo.Context) error {

	c := e.Get("Carrier").(gatewayModels.Carrier)
	messagesInvolved, paymentsInvolved, filesInvolved, limit := false, false, false, defaultQueryLimit
	registeredBefore, registeredAfter, dateOfBirth := time.Time{}, time.Time{}, time.Time{}
	if strings.ToLower(e.QueryParam("messages")) == "true" {
		messagesInvolved = true
	}
	if strings.ToLower(e.QueryParam("payments")) == "true" {
		paymentsInvolved = true
	}
	if strings.ToLower(e.QueryParam("files")) == "true" {
		filesInvolved = true
	}
	if !reflect.ValueOf(e.QueryParam("limit")).IsZero() {
		l, err := strconv.Atoi(e.QueryParam("limit"))
		if err == nil {

			limit = l
		}
	}
	var err error
	if e.QueryParam("registeredBefore") != "" {
		registeredBefore, err = time.Parse(time.RFC3339, e.QueryParam("registeredBefore"))
		if err != nil {
			fmt.Println(err)
			return customErrors.InvalidInputsErrorHandler(e, []string{"registeredBefore is invalid"})
		}
	}
	if e.QueryParam("registeredAfter") != "" {
		registeredAfter, err = time.Parse(time.RFC3339, e.QueryParam("registeredAfter"))
		if err != nil {
			return customErrors.InvalidInputsErrorHandler(e, []string{"registeredAfter is invalid"})
		}
	}

	if e.QueryParam("dateOfBirth") != "" {
		dateOfBirth, err = time.Parse(time.RFC3339, e.QueryParam("dateOfBirth"))
		if err != nil {
			return customErrors.InvalidInputsErrorHandler(e, []string{"birthday is invalid"})
		}
	}

	queryParams := models.UserQueryInfo{
		Email:            e.QueryParam("email"),
		FirstName:        e.QueryParam("firstName"),
		LastName:         e.QueryParam("lastName"),
		Country:          e.QueryParam("country"),
		Plan:             e.QueryParam("plan"),
		Gender:           e.QueryParam("gender"),
		Rank:             e.QueryParam("rank"),
		DateOfBirth:      dateOfBirth,
		RegisteredBefore: registeredBefore,
		RegisteredAfter:  registeredAfter,
		Messages:         messagesInvolved,
		Payments:         paymentsInvolved,
		Files:            filesInvolved,
		Limit:            limit,
	}
	users, ce := Handler.GetUsers(&c, queryParams)
	fmt.Println(users, ce)
	if ce.Err != nil {
		c.SetError(ce.Err)
		return customErrors.HandleCommonErrors(e, ce)
	}

	c.Finilize()
	return httpResponses.Success(e, 200, []interface{}{users})

}

const (
	defaultQueryLimit = 100
)
