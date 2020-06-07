package internal

import (
	"fmt"
	"github.com/lithammer/shortuuid"
	"testing"
)

var (
	createQRISParam = CreateQRISParams{
		ExternalID:  fmt.Sprintf("QR_%s", shortuuid.New()),
		Type:        "DYNAMIC",
		CallbackURL: "https://naufalihsan.co.id/payment/qris/callback",
		Amount:      50000,
	}

	getQRISByExternalID = GetQRISParam{
		ExternalID: createQRISParam.ExternalID,
	}

	simulatePaymentWithQRIS = PaymentParam{
		ExternalID: createQRISParam.ExternalID,
		Amount:     createQRISParam.Amount,
	}
)

func TestCreateQRIS(t *testing.T) {
	resp := CreateQRIS(createQRISParam)
	t.Logf(resp.QRIS.ExternalID)
	t.Logf(resp.Filename)
}

func TestGetQRISByExternalID(t *testing.T) {
	resp := GetQRISByExternalID(getQRISByExternalID)
	t.Logf(resp.QRString)
}

func TestSimulatePaymentWithQRIS(t *testing.T) {
	resp := SimulatePaymentWithQRIS(simulatePaymentWithQRIS)
	t.Logf(resp.ID)
	t.Logf(resp.Status)
}
