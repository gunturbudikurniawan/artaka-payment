package adapter

import (
	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"github.com/xendit/xendit-go/client"
	"os"
)

var (
	secretKey        string
	authorizationKey string
	xenditCli        *client.API
	qrisCli          *wrapper.API
)

func Authorize() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("SECRET_KEY")
	authorization := os.Getenv("AUTHORIZATION")

	secretKey = secret
	authorizationKey = authorization

	xenditCli = client.New(secretKey)
	qrisCli = wrapper.New(secretKey)

}

func GetSecretKey() string {
	return secretKey
}

func GetAuthorizationKey() string {
	return authorizationKey
}

func GetXenditCli() *client.API {
	return xenditCli
}

func GetQRISCli() *wrapper.API {
	return qrisCli
}
