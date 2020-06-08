package wrapper

import "github.com/xendit/xendit-go/disbursement"

type QRIS struct {
	ID          string  `json:"id"`
	ExternalID  string  `json:"external_id"`
	Amount      float64 `json:"amount"`
	QRString    string  `json:"qr_string"`
	CallbackURL string  `json:"callback_url"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	Created     string  `json:"created"`
	Updated     string  `json:"updated"`
}

type CreateQRISResponse struct {
	QRIS     *QRIS  `json:"qris"`
	Filename string `json:"filename"`
}

type PaymentResponse struct {
	ID      string  `json:"id"`
	Amount  float64 `json:"amount"`
	Created string  `json:"created"`
	QRCode  QRIS    `json:"qr_code"`
	Status  string  `json:"status"`
}

type CreateQRISParams struct {
	ForUserID   string  `json:"-"`
	ExternalID  string  `json:"external_id" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	CallbackURL string  `json:"callback_url",validate:"required,url"`
	Amount      float64 `json:"amount,omitempty"`
}

type GetQRISParam struct {
	ExternalID string `json:"external_id" validate:"required"`
}

type PaymentParam struct {
	ExternalID string  `json:"external_id" validate:"required"`
	Amount     float64 `json:"amount,omitempty"`
}

type QRISCallbackParam struct {
	Event   string  `json:"event" validate:"required"`
	ID      string  `json:"id" validate:"required"`
	Amount  float64 `json:"amount" validate:"required"`
	Created string  `json:"created" validate:"required"`
	QRCode  QRIS    `json:"qr_code" validate:"required"`
	Status  string  `json:"status" validate:"required"`
}

type VACallbackParam struct {
	PaymentID                string `json:"payment_id" validate:"required"`
	CallbackVirtualAccountID string `json:"callback_virtual_account_id" validate:"required"`
	OwnerID                  string `json:"owner_id" validate:"required"`
	ExternalID               string `json:"external_id" validate:"required"`
	AccountNumber            string `json:"account_number" validate:"required"`
	BankCode                 string `json:"bank_code" validate:"required"`
	Amount                   string `json:"amount" validate:"required"`
	MerchantCode             string `json:"merchant_code" validate:"required"`
	ID                       string `json:"id" validate:"required"`
	TransactionTimestamp     string `json:"transaction_timestamp" validate:"required"`
}

type VAUpdateCallbackParam struct {
	OwnerID         string `json:"owner_id" validate:"required"`
	ExternalID      string `json:"external_id" validate:"required"`
	BankCode        string `json:"bank_code" validate:"required"`
	MerchantCode    string `json:"merchant_code" validate:"required"`
	Name            string `json:"name" validate:"required"`
	AccountNumber   string `json:"account_number" validate:"required"`
	SuggestedAmount string `json:"suggested_amount,omitempty"`
	IsClosed        bool   `json:"is_closed" validate:"required"`
	ExpectedAmount  string `json:"expected_amount,omitempty"`
	ID              string `json:"id" validate:"required"`
	IsSingleUse     bool   `json:"is_single_use" validate:"required"`
	Status          string `json:"status" validate:"required"`
}

type DisbursementCallback struct {
	IsInstant               string   `json:"is_instant" validate:"required"`
	UserID                  string   `json:"user_id" validate:"required"`
	ExternalID              string   `json:"external_id" validate:"required"`
	Amount                  float64  `json:"amount" validate:"required"`
	BankCode                string   `json:"bank_code" validate:"required"`
	AccountHolderName       string   `json:"account_holder_name" validate:"required"`
	DisbursementDescription string   `json:"disbursement_description" validate:"required"`
	Status                  string   `json:"status" validate:"required"`
	FailureCode             string   `json:"failure_code,omitempty"`
	ID                      string   `json:"id" validate:"required"`
	EmailTo                 []string `json:"email_to,omitempty"`
	EmailCc                 []string `json:"email_cc,omitempty"`
	EmailBcc                []string `json:"email_bcc,omitempty"`
}

type BatchDisbursementCallback struct {
	Created              string                        `json:"created" validate:"required"`
	Updated              string                        `json:"updated" validate:"required"`
	Reference            string                        `json:"reference" validate:"required"`
	TotalUploadedCount   float64                       `json:"total_uploaded_count" validate:"required"`
	TotalUploadedAmount  float64                       `json:"total_uploaded_amount" validate:"required"`
	ApprovedAt           string                        `json:"approved_at" validate:"required"`
	ApproverID           string                        `json:"approver_id" validate:"required"`
	Status               string                        `json:"status" validate:"required"`
	ID                   string                        `json:"id" validate:"required"`
	UserID               string                        `json:"user_id" validate:"required"`
	TotalErrorCount      float64                       `json:"total_error_count" validate:"required"`
	TotalErrorAmount     float64                       `json:"total_error_amount" validate:"required"`
	TotalDisbursedCount  float64                       `json:"total_disbursed_count" validate:"required"`
	TotalDisbursedAmount float64                       `json:"total_disbursed_amount" validate:"required"`
	Disbursements        disbursement.DisbursementItem `json:"disburse,omitempty"`
}
