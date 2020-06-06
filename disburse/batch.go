package disburse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"github.com/xendit/xendit-go/disbursement"
	"net/http"
)

func CreateBatchDisbursement(c *gin.Context) {
	var payload disbursement.CreateBatchParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetXenditCli()
	response, err := client.Disbursement.CreateBatch(&payload)
	if err != nil {
		fmt.Println(err.Message)
	}

	c.JSON(http.StatusOK, response)
}

func BatchDisbursementCallback(c *gin.Context) {
	var payload wrapper.BatchDisbursementCallback
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payload)
}
