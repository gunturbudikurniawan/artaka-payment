package internal

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func GetAvailableBanks() []xendit.VirtualAccountBank {
	client := GetXenditCli()
	response, _ := client.VirtualAccount.GetAvailableBanks()
	return response
}

func CreateFixedVA(payload virtualaccount.CreateFixedVAParams) *xendit.VirtualAccount {
	client := GetXenditCli()
	response, _ := client.VirtualAccount.CreateFixedVA(&payload)
	return response
}

func GetFixedVA(payload virtualaccount.GetFixedVAParams) *xendit.VirtualAccount {
	client := GetXenditCli()
	response, _ := client.VirtualAccount.GetFixedVA(&payload)
	return response
}

func UpdateFixedVA(payload virtualaccount.UpdateFixedVAParams) *xendit.VirtualAccount {
	client := GetXenditCli()
	response, _ := client.VirtualAccount.UpdateFixedVA(&payload)
	return response
}

func GetPaymentVA(payload virtualaccount.GetPaymentParams) *xendit.VirtualAccountPayment {
	client := GetXenditCli()
	response, _ := client.VirtualAccount.GetPayment(&payload)
	return response
}
