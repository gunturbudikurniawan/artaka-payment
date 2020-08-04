package shipment

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalihsan/artaka-payment/adapter"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"net/http"
)

func GetAvailableProvince(c *gin.Context) {
	query := c.Request.URL.Query()
	isQuery, strQuery := wrapper.ParamToString(query)

	client := adapter.GetRajaOngkirCli()

	if isQuery {
		response, error := client.GetProvinceQuery(strQuery)
		if error != nil {
			c.JSON(http.StatusOK, error)
			return
		}
		c.JSON(http.StatusOK, response)
	} else {
		response, error := client.GetProvinces()
		if error != nil {
			c.JSON(http.StatusOK, error)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetAvailableCity(c *gin.Context) {
	query := c.Request.URL.Query()
	isQuery, strQuery := wrapper.ParamToString(query)

	client := adapter.GetRajaOngkirCli()

	if isQuery {
		response, error := client.GetCityQuery(strQuery)
		if error != nil {
			c.JSON(http.StatusOK, error)
			return
		}
		c.JSON(http.StatusOK, response)
	} else {
		response, error := client.GetCities()
		if error != nil {
			c.JSON(http.StatusOK, error)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetAvailableSubdistrict(c *gin.Context) {
	query := c.Request.URL.Query()
	_, strQuery := wrapper.ParamToString(query)

	existID := c.Query("id") == ""

	client := adapter.GetRajaOngkirCli()

	if existID {
		response, error := client.GetSubdistrictByCity(strQuery)
		if error != nil {
			c.JSON(http.StatusOK, error)
			return
		}
		c.JSON(http.StatusOK, response)
	} else {
		response, error := client.GetSubdistrictByID(strQuery)
		if error != nil {
			c.JSON(http.StatusOK, error)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func CountCostEstimation(c *gin.Context) {
	var payload wrapper.CostEstimationParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetRajaOngkirCli()
	response, error := client.CountCostEstimation(&payload)

	if error != nil {
		c.JSON(http.StatusOK, error)
		return
	}

	c.JSON(http.StatusOK, response)
}

func CheckWayBill(c *gin.Context) {
	var payload wrapper.WayBillParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := adapter.GetRajaOngkirCli()
	response, error := client.CheckWayBill(&payload)

	if error != nil {
		c.JSON(http.StatusOK, error)
		return
	}

	c.JSON(http.StatusOK, response)
}
