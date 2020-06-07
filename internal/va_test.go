package internal

import (
	"fmt"
	"github.com/lithammer/shortuuid"
	"github.com/xendit/xendit-go/virtualaccount"
	"testing"
)

var (
	isClosed, isSingleUse = true, true
	createFixedVA         = virtualaccount.CreateFixedVAParams{
		ExternalID:     fmt.Sprintf("VA_%s", shortuuid.New()),
		BankCode:       "BNI",
		Name:           "Naufal Ihsan Pratama",
		IsClosed:       &isClosed,
		ExpectedAmount: 50000,
		IsSingleUse:    &isSingleUse,
	}

	getFixedVA = virtualaccount.GetFixedVAParams{
		ID: "5edc57b4b3020349a27a9a69",
	}
)

func TestGetAvailableBanks(t *testing.T) {
	resp := GetAvailableBanks()
	for _, r := range resp {
		t.Logf("Bank Name:%s", r.Name)
		t.Logf("Bank Code:%s", r.Code)
	}
}

func TestCreateFixedVA(t *testing.T) {
	resp := CreateFixedVA(createFixedVA)
	t.Logf("Virtual Account ID :%s", resp.ID)
}

func TestGetFixedVA(t *testing.T) {
	resp := GetFixedVA(getFixedVA)
	t.Logf(resp.ID)
	t.Logf(resp.Name)
}
