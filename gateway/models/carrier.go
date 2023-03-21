package models

import (
	"context"
	"fmt"
	"log"

	userModels "github.com/thegreatdaniad/crypto-invest/services/user/models"
)

type Carrier struct {
	Context context.Context
	Steps   []Payload
	User    userModels.User
	Options Options
}

func (c *Carrier) Push(i Payload) {
	c.Steps = append(c.Steps, i)
}

func (c *Carrier) Pop() {
	c.Steps = c.Steps[:len(c.Steps)-1]
}

func (c *Carrier) GetCurrentStep() Payload {
	return c.Steps[len(c.Steps)-1]
}
func (c *Carrier) SetCurrentStep(p Payload) {
	c.Steps[len(c.Steps)-1] = p
}
func (c *Carrier) GetLogs() string {
	var log string
	for i, payload := range c.Steps {
		log += "\nStep " + fmt.Sprint(i)
		log += "\nAction: " + payload.Action
		log += "\nService: " + payload.Service
		log += "\nStatus: " + payload.Status
		log += "\nReport: " + payload.Report
		log += "\n\n"
	}
	return log
}

func (c *Carrier) InitializeWithData(data interface{}, service string, action string) {
	c.Push(Payload{
		Data:    data,
		Status:  PendingStatus,
		Service: service,
		Action:  action,
	})
}

func (c *Carrier) SetError(err error) {
	if len(c.Steps) < 1 {
		c.Push(Payload{})
	}
	log.Print(err)
	currentStep := c.GetCurrentStep()
	currentStep.Status = FailedStatus
	currentStep.Report = err.Error()
}

func (c *Carrier) SetSuccess(msg string) {
	log.Print(msg)
	currentStep := c.GetCurrentStep()
	currentStep.Status = SuccessStatus
	currentStep.Report = msg
}

func (c *Carrier) EndPointEntry(data interface{}, endpoint string) {
	c.Push(Payload{
		Data:    data,
		Status:  SuccessStatus,
		Report:  "Data has been successfully entered to the REST endpoint",
		Service: "Endpoint:  " + endpoint,
		Action:  "Entering data to the endpoint",
	})
}

func (c *Carrier) Finilize() {
	c.Push(Payload{
		Status: SuccessStatus,
		Report: "Request has been processed with no error",
		Action: "Sending the response to the user",
	})
}

type Payload struct {
	Status  string
	Action  string
	Report  string
	Service string
	Data    interface{}
}

type Options struct {
	Self bool // if the action is relevant to the user who makes it
}
