package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/disburse"
	"github.com/naufalihsan/artaka-payment/linkaja"
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

func RouterDisbursement(api *gin.RouterGroup) {
	api.GET("/banks", disburse.GetAvailableBanks)
	api.POST("/create", disburse.CreateDisbursement)
	api.POST("/get", disburse.GetDisbursementByExternalID)
	api.POST("/callback", disburse.DisbursementCallback)
	api.POST("/batch/create", disburse.CreateBatchDisbursement)
	api.POST("/batch/callback", disburse.BatchDisbursementCallback)
}

func RouterLink(api *gin.RouterGroup) {
	api.POST("/generate", linkaja.GenerateDynamicQR)
	api.POST("/callback", linkaja.InformPaymentQR)
	api.POST("/reverse", linkaja.ReverseTransaction)
}
