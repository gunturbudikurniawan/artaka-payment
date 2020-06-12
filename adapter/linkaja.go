package adapter

import (
	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"os"
)

var (
	linkSecretKey string
	cidKey        string
	linkajaCli    *wrapper.ClientLinkAja
)

func AuthorizeLinkAja() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	secretKey := os.Getenv("LINK_SECRET_KEY")
	cid := os.Getenv("CID")

	linkSecretKey = secretKey
	cidKey = cid

	linkajaCli = wrapper.NewClient(linkSecretKey, cidKey)

}

func GetLinkSecretKey() string {
	return linkSecretKey
}

func GetLinkCID() string {
	return cidKey
}

func GetLinkAjaCli() *wrapper.ClientLinkAja {
	return linkajaCli
}
