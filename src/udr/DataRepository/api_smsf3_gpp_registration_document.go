/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/gin-gonic/gin"
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udr/logger"
	"free5gc/src/udr/udr_handler/udr_message"
)

// CreateSmsfContext3gpp - Create the SMSF context data of a UE via 3GPP access
func CreateSmsfContext3gpp(c *gin.Context) {
	var smsfRegistration models.SmsfRegistration
	if err := c.ShouldBindJSON(&smsfRegistration); err != nil {
		logger.DataRepoLog.Panic(err.Error())
	}

	req := http_wrapper.NewRequest(c.Request, smsfRegistration)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventCreateSmsfContext3gpp, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// DeleteSmsfContext3gpp - To remove the SMSF context data of a UE via 3GPP access
func DeleteSmsfContext3gpp(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventDeleteSmsfContext3gpp, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// QuerySmsfContext3gpp - Retrieves the SMSF context data of a UE using 3gpp access
func QuerySmsfContext3gpp(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventQuerySmsfContext3gpp, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}