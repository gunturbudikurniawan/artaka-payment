package adapter

import (
	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"os"
)

var (
	apiKey        string
	rajaOngkirCli *wrapper.ClientRajaOngkir
)

func AuthorizeRajaOngkir() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	api := os.Getenv("RAJA_ONGKIR_API_KEY")
	apiKey = api

	rajaOngkirCli = wrapper.NewROClient(apiKey)

}

func GetRajaOngkirApiKey() string {
	return apiKey
}

func GetRajaOngkirCli() *wrapper.ClientRajaOngkir {
	return rajaOngkirCli
}
