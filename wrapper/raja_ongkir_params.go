package wrapper

type ProvinceBaseSingle struct {
	RajaOngkir ProvinceBodySingle `json:"rajaongkir"`
}

type ProvinceBaseAll struct {
	RajaOngkir ProvinceBodyAll `json:"rajaongkir"`
}

type ProvinceBodySingle struct {
	Query   ProvinceQuery  `json:"query"`
	Status  ResponseStatus `json:"status"`
	Results ProvinceItem   `json:"results"`
}

type ProvinceBodyAll struct {
	Query   []string       `json:"query"`
	Status  ResponseStatus `json:"status"`
	Results []ProvinceItem `json:"results"`
}

type ProvinceQuery struct {
	ID string `json:"id,omitempty"`
}

type ResponseStatus struct {
	Code        int64  `json:"code"`
	Description string `json:"description"`
}

type ProvinceItem struct {
	ProvinceID string `json:"province_id,omitempty"`
	Province   string `json:"province,omitempty"`
}

type CityBaseSingle struct {
	RajaOngkir CityBodySingle `json:"rajaongkir"`
}

type CityBaseAll struct {
	RajaOngkir CityBodyAll `json:"rajaongkir"`
}

type CityBodySingle struct {
	Query   CityQuery      `json:"query"`
	Status  ResponseStatus `json:"status"`
	Results CityItem       `json:"results"`
}

type CityBodyAll struct {
	Query   []string       `json:"query"`
	Status  ResponseStatus `json:"status"`
	Results []CityItem     `json:"results"`
}

type CityQuery struct {
	ID       string `json:"id,omitempty"`
	Province string `json:"province,omitempty"`
}

type CityItem struct {
	CityID     string `json:"city_id,omitempty"`
	ProvinceID string `json:"province_id,omitempty"`
	Province   string `json:"province,omitempty"`
	Type       string `json:"type,omitempty"`
	CityName   string `json:"city_name,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

type SubdistrictBaseSingle struct {
	RajaOngkir SubdistrictBodySingle `json:"rajaongkir"`
}

type SubdistrictBaseAll struct {
	RajaOngkir SubdistrictBodyAll `json:"rajaongkir"`
}

type SubdistrictBodySingle struct {
	Query   SubdistrictQuery `json:"query"`
	Status  ResponseStatus   `json:"status"`
	Results SubdistrictItem  `json:"results"`
}

type SubdistrictBodyAll struct {
	Query   SubdistrictQuery  `json:"query"`
	Status  ResponseStatus    `json:"status"`
	Results []SubdistrictItem `json:"results"`
}

type SubdistrictQuery struct {
	ID       string `json:"id,omitempty"`
	Province string `json:"province,omitempty"`
}

type SubdistrictItem struct {
	SubdistrictID   string `json:"subdistrict_id,omitempty"`
	ProvinceID      string `json:"province_id,omitempty"`
	CityID          string `json:"city_id,omitempty"`
	City            string `json:"city,omitempty"`
	Type            string `json:"type,omitempty"`
	SubdistrictName string `json:"subdistrict_name,omitempty"`
}

type CostEstimationParams struct {
	Origin          string  `json:"origin"`
	OriginType      string  `json:"originType"`
	Destination     string  `json:"destination"`
	DestinationType string  `json:"destinationType"`
	Weight          float64 `json:"weight"`
	Courier         string  `json:"courier"`
}

type CostEstimationBase struct {
	RajaOngkir CostEstimationBody `json:"rajaongkir"`
}

type CostEstimationBody struct {
	Query              CostEstimationParams    `json:"query"`
	Status             ResponseStatus          `json:"status"`
	OriginDetails      SubdistrictItem         `json:"origin_details"`
	DestinationDetails SubdistrictItem         `json:"destination_details"`
	Results            []CourierEstimationItem `json:"results"`
}

type CourierEstimationItem struct {
	Code  string               `json:"code"`
	Name  string               `json:"name"`
	Costs []CourierServiceItem `json:"costs"`
}

type CourierServiceItem struct {
	Service     string            `json:"service"`
	Description string            `json:"description,omitempty"`
	Cost        []ServiceCostItem `json:"cost"`
}

type ServiceCostItem struct {
	Value float64 `json:"value,omitempty"`
	Etd   string  `json:"etd,omitempty"`
	Note  string  `json:"note,omitempty"`
}

type WayBillParams struct {
	WayBill string `json:"waybill"`
	Courier string `json:"courier"`
}

type WayBillBase struct {
	RajaOngkir WayBillBody `json:"rajaongkir"`
}

type WayBillQuery struct {
	WayBill string `json:"waybill,omitempty"`
	Courier string `json:"courier,omitempty"`
}

type WayBillSummary struct {
	CourierCode   string `json:"courier_code"`
	CourierName   string `json:"courier_name"`
	WayBillNumber string `json:"waybill_number"`
	ServiceCode   string `json:"service_code"`
	WayBillDate   string `json:"waybill_date"`
	ShipperName   string `json:"shipper_name"`
	ReceiverName  string `json:"receiver_name"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Status        string `json:"status"`
}

type WayBillDetail struct {
	WayBillNumber    string `json:"waybill_number"`
	WayBillDate      string `json:"waybill_date"`
	Weight           string `json:"weight"`
	Origin           string `json:"origin"`
	Destination      string `json:"destination"`
	ShipperName      string `json:"shipper_name"`
	ShipperAddress1  string `json:"shipper_address1"`
	ShipperAddress2  string `json:"shipper_address2"`
	ShipperAddress3  string `json:"shipper_address3"`
	ReceiverName     string `json:"receiver_name"`
	ReceiverAddress1 string `json:"receiver_address1"`
	ReceiverAddress2 string `json:"receiver_address2"`
	ReceiverAddress3 string `json:"receiver_address3"`
	ReceiverCity     string `json:"receiver_city"`
}

type WayBillDeliveryStatus struct {
	Status      string `json:"status"`
	PodReceiver string `json:"pod_receiver"`
	PodDate     string `json:"pod_date"`
	PodTime     string `json:"pod_time"`
}

type WayBillManifestItem struct {
	ManifestCode        string `json:"manifest_code"`
	ManifestDescription string `json:"manifest_description"`
	ManifestDate        string `json:"manifest_date"`
	ManifestTime        string `json:"manifest_time"`
	CityName            string `json:"city_name"`
}

type WayBillItem struct {
	Delivered      bool                  `json:"delivered"`
	Summary        WayBillSummary        `json:"summary"`
	Details        WayBillDetail         `json:"details"`
	DeliveryStatus WayBillDeliveryStatus `json:"delivery_status"`
	Manifest       []WayBillManifestItem `json:"manifest"`
}

type WayBillBody struct {
	Query  WayBillQuery   `json:"query"`
	Status ResponseStatus `json:"status"`
	Result WayBillItem    `json:"result"`
}
