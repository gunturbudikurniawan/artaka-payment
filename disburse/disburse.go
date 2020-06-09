package disburse

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"github.com/xendit/xendit-go/disbursement"
	"net/http"
)

func GetAvailableBanks(c *gin.Context) {
	client := adapter.GetXenditCli()
	response, _ := client.Disbursement.GetAvailableBanks()
	c.JSON(http.StatusOK, response)
}

func CreateDisbursement(c *gin.Context) {
	var payload disbursement.CreateParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetXenditCli()
	response, err := client.Disbursement.Create(&payload)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response)
}

func GetDisbursementByExternalID(c *gin.Context) {
	var payload disbursement.GetByExternalIDParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetXenditCli()
	response, error := client.Disbursement.GetByExternalID(&payload)

	if error != nil {
		c.JSON(http.StatusOK, error)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DisbursementCallback(c *gin.Context) {
	var payload wrapper.DisbursementCallback
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payload)
}
