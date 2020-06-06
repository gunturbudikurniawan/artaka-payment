package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/models"
	"github.com/naufalihsan/artaka-payment/routers"
)

func main() {

	adapter.Authorize()

	db := adapter.Connect()
	models.AutoMigrate()
	defer db.Close()

	r := gin.Default()
	r.Static("/static", "./static")

	payment := r.Group("/payment")
	qr := payment.Group("/qris")
	{
		routers.RouterQRIS(qr)
	}

	va := payment.Group("/va")
	{
		routers.RouterVA(va)
	}

	r.Run(":8005")
}
