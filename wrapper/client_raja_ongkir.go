package wrapper

import (
	"fmt"
	"net/http"
)

func (c *ClientRajaOngkir) GetProvinces() (*ProvinceBaseAll, *error) {
	response := &ProvinceBaseAll{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"GET",
		fmt.Sprintf("%s/province", c.BaseURL),
		header,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) GetProvinceQuery(query string) (*ProvinceBaseSingle, *error) {
	response := &ProvinceBaseSingle{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"GET",
		fmt.Sprintf("%s/province%s", c.BaseURL, query),
		header,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) GetCities() (*CityBaseAll, *error) {
	response := &CityBaseAll{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"GET",
		fmt.Sprintf("%s/city", c.BaseURL),
		header,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) GetCityQuery(query string) (*CityBaseSingle, *error) {
	response := &CityBaseSingle{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"GET",
		fmt.Sprintf("%s/city%s", c.BaseURL, query),
		header,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) GetSubdistrictByCity(query string) (*SubdistrictBaseAll, *error) {
	response := &SubdistrictBaseAll{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"GET",
		fmt.Sprintf("%s/subdistrict%s", c.BaseURL, query),
		header,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) GetSubdistrictByID(query string) (*SubdistrictBaseSingle, *error) {
	response := &SubdistrictBaseSingle{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"GET",
		fmt.Sprintf("%s/subdistrict%s", c.BaseURL, query),
		header,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) CountCostEstimation(data *CostEstimationParams) (*CostEstimationBase, *error) {
	response := &CostEstimationBase{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"POST",
		fmt.Sprintf("%s/cost", c.BaseURL),
		header,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ClientRajaOngkir) CheckWayBill(data *WayBillParams) (*WayBillBase, *error) {
	response := &WayBillBase{}
	header := &http.Header{}

	header.Set("key", c.ApiKey)

	err := Call(
		"POST",
		fmt.Sprintf("%s/waybill", c.BaseURL),
		header,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
