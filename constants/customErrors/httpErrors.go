package customErrors

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	responseModels "github.com/thegreatdaniad/crypto-invest/constants/httpResponses"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
)

func BadRequestErrorHandler(e echo.Context, errs []string) error {
	r := responseModels.ResponseBody{
		Success: false,
		Errors:  errs,
	}
	e.JSON(400, r)
	return nil

}
func InvalidInputsErrorHandler(e echo.Context, errs []string) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	c.SetError(errors.New(strings.Join(errs, ",")))
	r := responseModels.ResponseBody{
		Success: false,
		Errors:  errs,
	}
	e.JSON(400, r)
	return nil

}
func UnauthorizedErrorHandler(e echo.Context, errs []string) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	c.SetError(errors.New(strings.Join(errs, ",")))
	r := responseModels.ResponseBody{
		Success: false,
		Errors:  errs,
	}
	e.JSON(401, r)
	return nil

}

func PermissionDeniedErrorHandler(e echo.Context, errs []string) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	c.SetError(errors.New(strings.Join(errs, ",")))
	r := responseModels.ResponseBody{
		Success: false,
		Errors:  errs,
	}
	e.JSON(403, r)
	return nil

}
func NotFoundErrorHandler(e echo.Context, errs []string) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	c.SetError(errors.New(strings.Join(errs, ",")))
	r := responseModels.ResponseBody{
		Success: false,
		Errors:  errs,
	}
	e.JSON(404, r)
	return nil
}
func InternalServerErrorHandler(e echo.Context, errs []string) error {
	c := e.Get("Carrier").(gatewayModels.Carrier)
	c.SetError(errors.New(strings.Join(errs, ",")))
	r := responseModels.ResponseBody{
		Success: false,
		Errors:  errs,
	}
	e.JSON(500, r)
	return nil

}

func HandleCommonErrors(e echo.Context, ce Error) error {

	switch ce.Code {
	case InternalServerError:
		return InternalServerErrorHandler(e, []string{InternalServerError})
	case NotFound:
		return NotFoundErrorHandler(e, []string{ce.Err.Error()})
	case Unauthorized:
		return UnauthorizedErrorHandler(e, []string{ce.Err.Error()})
	case BadRequest:
		return BadRequestErrorHandler(e, []string{ce.Err.Error()})
	case DatabaseConnectionError:
		return InternalServerErrorHandler(e, []string{InternalServerError})
	case InvalidInputs:
		return InvalidInputsErrorHandler(e, []string{ce.Err.Error()})
	case ProccessCanceledByParent:
		return InternalServerErrorHandler(e, []string{ProccessCanceledByParent})
	case UnknownError:
		return InternalServerErrorHandler(e, []string{ce.Err.Error()})
	default:
		return InternalServerErrorHandler(e, []string{InternalServerError})
	}
}
