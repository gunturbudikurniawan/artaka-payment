package wrapper

type ClientLinkAja struct {
	BaseURL    string
	CID        string
	SecretKey  string
	MerchantID string
}

func NewClient(secretKey string, cid string) *ClientLinkAja {
	client := ClientLinkAja{
		BaseURL:   "https://linkgw-dev.linkaja.com/api",
		CID:       cid,
		SecretKey: secretKey,
	}

	return &client
}
