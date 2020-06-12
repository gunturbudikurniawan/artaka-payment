package wrapper

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *ClientLinkAja) CreateDynamicQR(data *LinkAjaQRParams) (*LinkAjaQR, *error) {
	response := &LinkAjaQR{}
	header := &http.Header{}

	payload, _ := json.Marshal(data)
	signature := fmt.Sprintf("%s:%s:%s", c.CID, string(payload), c.SecretKey)

	hash := hmac.New(sha512.New, []byte(c.SecretKey))
	hash.Write([]byte(signature))

	sha := hex.EncodeToString(hash.Sum(nil))

	header.Set("cid", c.CID)
	header.Set("signature", sha)

	err := Call(
		"POST",
		fmt.Sprintf("%s/v1/qr/generate", c.BaseURL),
		header,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientLinkAja) ReverseTransactionQR(data *ReverseTransactionParams) (*ReverseTransaction, *error) {
	response := &ReverseTransaction{}
	header := &http.Header{}

	err := Call(
		"POST",
		fmt.Sprintf("%s/pinless/v1/transactions", c.BaseURL),
		header,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
