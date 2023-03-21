package models

import (
	"github.com/labstack/echo/v4"
)

type Gateway struct {
	Handler *echo.Echo
}

func (g *Gateway) AddMiddlewares(m []echo.MiddlewareFunc) {
	for _, middleware := range m {
		g.Handler.Use(middleware)
	}
}

func NewGateway() Gateway {
	return Gateway{
		echo.New(),
	}
}


const (
	//statuses
	PendingStatus  = "Pending"
	SuccessStatus  = "Success"
	WarningStatus  = "Warning"
	FailedStatus   = "Failed"
	RejectedStatus = "Rejected"

	//services
	UserService         = "User Service"

)

