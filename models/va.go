package models

import "github.com/jinzhu/gorm"

type Va struct {
	gorm.Model
	VaID                 string
	Name                 string
	AccountNumber        string
	ExternalID           string
	BankCode             string
	MerchantCode         string
	IsClosed             *bool
	IsSingleUse          *bool
	PaymentID            string
	TransactionTimestamp string
	Status               string
}
