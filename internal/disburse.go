package internal

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
)

func DisbursementAvailableBanks() []xendit.DisbursementBank {
	client := GetXenditCli()
	response, _ := client.Disbursement.GetAvailableBanks()
	return response
}

func CreateDisbursement(payload disbursement.CreateParams) *xendit.Disbursement {
	client := GetXenditCli()
	response, err := client.Disbursement.Create(&payload)
	if err != nil {
		panic(err)
	}
	return response
}

func GetDisbursementByExternalID(payload disbursement.GetByExternalIDParams) []xendit.Disbursement {
	client := GetXenditCli()
	response, _ := client.Disbursement.GetByExternalID(&payload)
	return response
}

func CreateBatchDisbursement(payload disbursement.CreateBatchParams) *xendit.BatchDisbursement {
	client := GetXenditCli()
	response, err := client.Disbursement.CreateBatch(&payload)
	if err != nil {
		panic(err)
	}
	return response
}
