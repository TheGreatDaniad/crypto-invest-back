package httpResponses

import (
	"github.com/labstack/echo/v4"
)

type ResponseBody struct {
	Data     []interface{} `json:"data"`
	Success  bool          `json:"success"`
	Errors   []string      `json:"errors"`
	Warnings []string      `json:"warnings"`
}

func (r *ResponseBody) SetSuccess(data []interface{}) {
	r.Data = data
	r.Success = true
}

func (r *ResponseBody) SetErrors(errors []string) {
	r.Success = false
	r.Errors = errors
}

func (r *ResponseBody) SetWarnings(warnings []string) {
	r.Success = false
	r.Errors = warnings
}

func Success(e echo.Context, statusCode int, data []interface{}) error {
	r := ResponseBody{
		Success: true,
		Data:    data,
	}
	e.JSON(200, r)
	return nil
}
