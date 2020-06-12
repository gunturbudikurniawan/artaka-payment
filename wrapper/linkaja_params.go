package wrapper

type LinkAjaQRParams struct {
	Fee               string `json:"fee"`
	Amount            string `json:"amount"`
	City              string `json:"city"`
	PostalCode        string `json:"postalCode"`
	MerchantName      string `json:"merchantName"`
	MerchantID        string `json:"merchantID"`
	MerchantCriteria  string `json:"merchantCriteria"`
	MerchantTrxID     string `json:"merchantTrxID"`
	MerchantPan       string `json:"merchantPan"`
	PartnerMerchantID string `json:"partnerMerchantID"`
}

type LinkAjaQR struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	QRString        string `json:"qrString"`
	MerchantTrxID   string `json:"merchantTrxID"`
}

type InformPaymentParams struct {
	Merchant string  `json:"merchant"`
	Terminal string  `json:"terminal"`
	Pwd      string  `json:"pwd"`
	TrxType  string  `json:"trx_type"`
	TrxDate  int64   `json:"trx_date"`
	TrxID    string  `json:"trx_id"`
	Msisdn   string  `json:"msisdn"`
	Msg      string  `json:"msg"`
	Amount   float64 `json:"amount"`
}

type InformPayment struct {
	ResponseCode        string `json:"responseCode"`
	TransactionID       string `json:"transactionID"`
	NotificationMessage string `json:"notificationMessage"`
}

type ReverseTransactionParams struct {
	CommandID              string `json:"commandID"`
	OriginatorConversionID string `json:"originatorConversationID"`
	OriginalReceiptNumber  string `json:"originalReceiptNumber"`
	TrxAmount              string `json:"trxAmount"`
}

type ReverseTransaction struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	Body   ResponseBody   `json:"res:Body"`
	Header ResponseHeader `json:"res:Header"`
}

type ResponseHeader struct {
	OriginatorConversationID string `json:"res:OriginatorConversationID"`
	Version                  int64  `json:"res:Version"`
	ConversationID           string `json:"res:ConversationID"`
}

type ResponseBody struct {
	ResultCode              int64                   `json:"res:ResultCode"`
	RevertTransactionResult RevertTransactionResult `json:"res:ReverseTransactionResult"`
	ReferenceData           string                  `json:"res:ReferenceData"`
	ResultType              int64                   `json:"res:ResultType"`
	ResultDesc              string                  `json:"res:ResultDesc"`
}

type RevertTransactionResult struct {
	Amount                float64 `json:"res:Amount"`
	TransactionStatus     string  `json:"res:TransactionStatus"`
	CreditPartyPublicName string  `json:"res:CreditPartyPublicName"`
	DebitPartyPublicName  string  `json:"res:DebitPartyPublicName"`
	ReceiptNumber         string  `json:"res:ReceiptNumber"`
	BOCompletedTime       int64   `json:"res:BOCompletedTime"`
	TransCompletedTime    int64   `json:"res:TransCompletedTime"`
	CreditAccountBalance  string  `json:"res:CreditAccountBalance"`
	DebitAccountBalance   string  `json:"res:DebitAccountBalance"`
	Charge                float64 `json:"res:Charge"`
	OriginalAmount        float64 `json:"res:OriginalAmount"`
	OriginalReceiptNumber string  `json:"res:OriginalReceiptNumber"`
}
