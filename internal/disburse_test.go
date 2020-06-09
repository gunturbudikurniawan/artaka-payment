package internal

import (
	"fmt"
	"github.com/lithammer/shortuuid"
	"github.com/xendit/xendit-go/disbursement"
	"strconv"
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

func TestCreateThousandBatchDisbursement(t *testing.T) {
	const stress = 1000
	var testingItems []disbursement.DisbursementItem
	for i := 0; i < stress; i++ {
		items := disbursement.DisbursementItem{
			ExternalID:        fmt.Sprintf("DB_%s", shortuuid.New()),
			Amount:            float64(Random(10000, 500000)),
			BankCode:          "BCA",
			BankAccountName:   "Naufal Ihsan",
			BankAccountNumber: strconv.FormatInt(int64(Random(1000000000, 9999999999)), 10),
			Description:       fmt.Sprintf("Reimbursement for payment X-%v", i),
		}

		testingItems = append(testingItems, items)
	}

	stressBatch := disbursement.CreateBatchParams{
		Reference:     "Stress Batch",
		Disbursements: testingItems,
	}

	resp := CreateBatchDisbursement(stressBatch)
	t.Logf(resp.Status)
	t.Logf(fmt.Sprintf("%f", resp.TotalUploadedAmount))

}
