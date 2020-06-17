package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/shipment"
)

func RouterShipment(api *gin.RouterGroup) {
	api.GET("/province", shipment.GetAvailableProvince)
	api.GET("/city", shipment.GetAvailableCity)
	api.GET("/subdistrict", shipment.GetAvailableSubdistrict)
	api.POST("/cost", shipment.CountCostEstimation)

}
