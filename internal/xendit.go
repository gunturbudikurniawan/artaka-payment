package internal

import (
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/client"
	"os"
)

func Init() (string, string) {
	err := godotenv.Load("internal.env")
	if err != nil {
		panic(err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	authorization := os.Getenv("AUTHORIZATION")
	return secretKey, authorization
}

func GetXenditCli() *client.API {
	secretKey, _ := Init()
	return client.New(secretKey)
}

func GetQRISCli() *API {
	secretKey, _ := Init()
	return New(secretKey)
}
