package internal

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
)

func CreateQRIS(payload CreateQRISParams) *CreateQRISResponse {
	client := GetQRISCli()
	response, _ := client.QRIS.CreateQRIS(&payload)

	filename := fmt.Sprintf("../static/%s", response.ID)
	err := qrcode.WriteFile(response.QRString, qrcode.Medium, 256, filename)
	if err != nil {
		log.Fatal("Error generate file")
	}

	return &CreateQRISResponse{QRIS: response, Filename: filename}

}

func GetQRISByExternalID(payload GetQRISParam) *QRIS {
	client := GetQRISCli()
	response, _ := client.QRIS.GetQRIS(&payload)
	return response
}

func SimulatePaymentWithQRIS(payload PaymentParam) *PaymentResponse {
	client := GetQRISCli()
	response, _ := client.QRIS.SimulatePayment(&payload)
	return response
}
