package va

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/models"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"github.com/xendit/xendit-go/virtualaccount"
	"net/http"
)

func GetAvailableBanks(c *gin.Context) {
	client := adapter.GetXenditCli()
	response, _ := client.VirtualAccount.GetAvailableBanks()
	c.JSON(http.StatusOK, response)
}

func CreateFixedVA(c *gin.Context) {
	var payload virtualaccount.CreateFixedVAParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := adapter.GetXenditCli()
	response, _ := client.VirtualAccount.CreateFixedVA(&payload)

	c.JSON(http.StatusOK, response)
}

func GetFixedVA(c *gin.Context) {
	var payload virtualaccount.GetFixedVAParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetXenditCli()
	response, _ := client.VirtualAccount.GetFixedVA(&payload)

	adapter.GetDB().Create(&models.Va{
		VaID:          response.ID,
		Name:          response.Name,
		AccountNumber: response.AccountNumber,
		ExternalID:    response.ExternalID,
		BankCode:      response.BankCode,
		MerchantCode:  response.MerchantCode,
		IsClosed:      response.IsClosed,
		IsSingleUse:   response.IsSingleUse,
		Status:        response.Status,
	})

	c.JSON(http.StatusOK, response)
}

func UpdateFixedVA(c *gin.Context) {
	var payload virtualaccount.UpdateFixedVAParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetXenditCli()
	response, _ := client.VirtualAccount.UpdateFixedVA(&payload)

	va := models.Va{VaID: payload.ID}
	adapter.GetDB().Model(&va).Updates(models.Va{
		VaID:          response.ID,
		Name:          response.Name,
		AccountNumber: response.AccountNumber,
		ExternalID:    response.ExternalID,
		BankCode:      response.BankCode,
		MerchantCode:  response.MerchantCode,
		IsClosed:      response.IsClosed,
		IsSingleUse:   response.IsSingleUse,
		Status:        response.Status,
	})

	c.JSON(http.StatusOK, response)
}

func VACallback(c *gin.Context) {
	var payload wrapper.VACallbackParam
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	va := models.Va{VaID: payload.ID}
	adapter.GetDB().Model(&va).Updates(models.Va{
		PaymentID:            payload.PaymentID,
		TransactionTimestamp: payload.TransactionTimestamp,
	})

	c.JSON(http.StatusOK, payload)
}

func VAUpdateCallback(c *gin.Context) {
	var payload wrapper.VAUpdateCallbackParam
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payload)
}

func GetPaymentVA(c *gin.Context) {
	var payload virtualaccount.GetPaymentParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetXenditCli()
	response, _ := client.VirtualAccount.GetPayment(&payload)

	c.JSON(http.StatusOK, response)
}
