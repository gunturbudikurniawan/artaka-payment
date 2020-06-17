package wrapper

type ClientRajaOngkir struct {
	BaseURL string
	ApiKey  string
}

func NewROClient(apiKey string) *ClientRajaOngkir {
	client := ClientRajaOngkir{
		BaseURL: "https://pro.rajaongkir.com/api",
		ApiKey:  apiKey,
	}

	return &client
}
