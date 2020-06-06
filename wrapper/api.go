package wrapper

import "github.com/xendit/xendit-go"

type API struct {
	opt          xendit.Option
	apiRequester xendit.APIRequester
	QRIS         *Client
}

type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

func (a *API) init() {
	a.QRIS = &Client{Opt: &a.opt, APIRequester: a.apiRequester}
}

func New(secretKey string) *API {
	api := API{
		opt: xendit.Option{
			SecretKey: secretKey,
			XenditURL: "https://api.xendit.co",
		},
		apiRequester: xendit.GetAPIRequester(),
	}

	api.init()

	return &api
}
