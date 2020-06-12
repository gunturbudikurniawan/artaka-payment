package linkaja

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"net/http"
)

func GenerateDynamicQR(c *gin.Context) {
	var payload wrapper.LinkAjaQRParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetLinkAjaCli()
	response, error := client.CreateDynamicQR(&payload)

	if error != nil {
		c.JSON(http.StatusOK, error)
		return
	}

	c.JSON(http.StatusOK, response)

}

func InformPaymentQR(c *gin.Context) {
	var payload wrapper.InformPaymentParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, _ := json.Marshal(payload)
	fmt.Println(string(log))

	response := wrapper.InformPayment{
		ResponseCode:        "00",
		TransactionID:       payload.TrxID,
		NotificationMessage: fmt.Sprintf("Notification %s", payload.Msg),
	}

	c.JSON(http.StatusOK, response)

}

func ReverseTransaction(c *gin.Context) {
	var payload wrapper.ReverseTransactionParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO
	//client := adapter.GetLinkAjaCli()
	//response, error := client.ReverseTransactionQR(&payload)
	//
	//if error != nil {
	//	c.JSON(http.StatusOK, error)
	//	return
	//}

	c.JSON(http.StatusOK, payload)
}
