package qris

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/models"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func CreateQRIS(c *gin.Context) {
	var payload wrapper.CreateQRISParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetQRISCli()
	response, _ := client.QRIS.CreateQRIS(&payload)

	filename := fmt.Sprintf("static/%s", response.ID)
	err := qrcode.WriteFile(response.QRString, qrcode.Medium, 256, filename)
	if err != nil {
		log.Fatal("Error generate file")
	}

	c.JSON(http.StatusOK, gin.H{"imgUrl": filename})

}

func GetQRISByExternalID(c *gin.Context) {
	var payload wrapper.GetQRISParam
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetQRISCli()
	response, _ := client.QRIS.GetQRIS(&payload)

	c.JSON(http.StatusOK, response)
}

func SimulatePaymentWithQRIS(c *gin.Context) {
	var payload wrapper.PaymentParam
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetQRISCli()
	response, _ := client.QRIS.SimulatePayment(&payload)

	c.JSON(http.StatusOK, response)
}

func PaymentStatusCallback(c *gin.Context) {
	var payload wrapper.QRISCallbackParam

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, _ := json.Marshal(payload)
	fmt.Println(string(log))

	adapter.GetDB().Create(&models.Qris{
		Event:      payload.Event,
		QrID:       payload.ID,
		Amount:     payload.Amount,
		ExternalID: payload.QRCode.ExternalID,
		QrString:   payload.QRCode.QRString,
		QrType:     payload.QRCode.Type,
		Status:     payload.Status,
	})

	c.JSON(http.StatusOK, payload)
}
