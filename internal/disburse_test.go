package internal

import (
	"fmt"
	"github.com/lithammer/shortuuid"
	"github.com/xendit/xendit-go/disbursement"
	"testing"
)

var (
	createDisburseParam = disbursement.CreateParams{
		ExternalID:        fmt.Sprintf("DB_%s", shortuuid.New()),
		Amount:            50000,
		BankCode:          "BCA",
		AccountHolderName: "Naufal Ihsan Pratama",
		AccountNumber:     "1231242315",
		Description:       "Reimbursement for payment X-123",
	}

	getDisburseByExternalID = disbursement.GetByExternalIDParams{
		ExternalID: createDisburseParam.ExternalID,
	}

	createBatchDisburseParam = disbursement.CreateBatchParams{
		Reference: "Batch-7-Juni",
		Disbursements: []disbursement.DisbursementItem{
			{
				ExternalID:        fmt.Sprintf("DB_%s", shortuuid.New()),
				Amount:            51000,
				BankCode:          "BCA",
				BankAccountName:   "Naufal Ihsan",
				BankAccountNumber: "12312424310",
				Description:       "Reimbursement for payment X-125",
			},
		},
	}
)

func TestDisbursementAvailableBanks(t *testing.T) {
	resp := DisbursementAvailableBanks()
	for _, r := range resp {
		t.Logf("Bank Name:%s", r.Name)
		t.Logf("Can Disburse:%v", r.CanDisburse)
	}
}

func TestCreateDisbursement(t *testing.T) {
	resp := CreateDisbursement(createDisburseParam)
	t.Logf(resp.ID)
	t.Logf(resp.Status)
}

func TestGetDisbursementByExternalID(t *testing.T) {
	resp := GetDisbursementByExternalID(getDisburseByExternalID)
	for _, r := range resp {
		t.Logf(r.ID)
		t.Logf(r.AccountHolderName)
		t.Logf(r.Status)
	}
}

func TestCreateBatchDisbursement(t *testing.T) {
	resp := CreateBatchDisbursement(createBatchDisburseParam)
	t.Logf(resp.Status)
	t.Logf(fmt.Sprintf("%f", resp.TotalUploadedAmount))

}
