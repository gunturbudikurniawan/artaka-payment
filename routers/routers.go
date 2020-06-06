package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/qris"
	"github.com/naufalihsan/artaka-payment/va"
)

func RouterQRIS(api *gin.RouterGroup) {
	api.POST("/create", qris.CreateQRIS)
	api.POST("/get", qris.GetQRISByExternalID)
	api.POST("/simulate", qris.SimulatePaymentWithQRIS)
	api.POST("/callback", qris.PaymentStatusCallback)
}

func RouterVA(api *gin.RouterGroup) {
	api.GET("/banks", va.GetAvailableBanks)
	api.POST("/create", va.CreateFixedVA)
	api.POST("/get", va.GetFixedVA)
	api.POST("/update", va.UpdateFixedVA)
	api.POST("/callback", va.VACallback)
	api.POST("/update/callback", va.VAUpdateCallback)
	api.POST("/payment", va.GetPaymentVA)
}
