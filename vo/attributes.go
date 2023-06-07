package vo

type CityAttributes struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type ProductAttributes struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
}

type PointOfServiceAttributes struct {
	DisplayName string `json:"displayName"`
	Address     struct {
		StreetName       string      `json:"streetName"`
		StreetNumber     string      `json:"streetNumber"`
		Town             string      `json:"town"`
		District         interface{} `json:"district"`
		Building         interface{} `json:"building"`
		FormattedAddress string      `json:"formattedAddress"`
	} `json:"address"`
}

type ReviewAttributes struct {
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
	UserName string `json:"userName"`
}

type OrderEntryAttributes struct {
	EntryNumber  int     `json:"entryNumber"`
	Quantity     int     `json:"quantity"`
	TotalPrice   float64 `json:"totalPrice"`
	DeliveryCost float64 `json:"deliveryCost"`
	BasePrice    float64 `json:"basePrice"`
}

type OrderAttributes struct {
	State      OrdersState  `json:"state"`
	Status     OrdersStatus `json:"status"`
	Code       string       `json:"code"`
	TotalPrice int          `json:"totalPrice"`
	Customer   struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		CellPhone string `json:"cellPhone"`
	} `json:"customer"`

	SignatureRequired   bool `json:"signatureRequired"`
	PreOrder            bool `json:"preOrder"`
	Express             bool `json:"express"`
	ReturnedToWarehouse bool `json:"returnedToWarehouse"`

	DeliveryMode    OrderDeliveryMode `json:"deliveryMode"`
	IsKaspiDelivery bool              `json:"isKaspiDelivery"`
	Latitude        float64           `json:"latitude"`
	Longitude       float64           `json:"longitude"`
	Apartment       string            `json:"apartment"`
	DeliveryAddress struct {
		StreetName       string      `json:"streetName"`
		StreetNumber     string      `json:"streetNumber"`
		Town             string      `json:"town"`
		District         interface{} `json:"district"`
		Building         interface{} `json:"building"`
		FormattedAddress string      `json:"formattedAddress"`
		Latitude         float64     `json:"latitude"`
		Longitude        float64     `json:"longitude"`
	} `json:"deliveryAddress"`
	KaspiDelivery struct {
		Waybill                         string `json:"waybill"`
		CourierTransmissionDate         int64  `json:"courierTransmissionDate"`
		CourierTransmissionPlanningDate int64  `json:"courierTransmissionPlanningDate"`
	} `json:"kaspiDelivery"`

	CreditTerm int `json:"creditTerm"`

	PaymentMode OrderPaymentMode `json:"paymentMode"`

	CreationDate                    int64 `json:"creationDate"`
	ApprovedByBankDate              int64 `json:"approvedByBankDate"`
	PlannedDeliveryDate             int64 `json:"plannedDeliveryDate"`
	ReservationDate                 int64 `json:"reservationDate"`
	CourierTransmissionPlanningDate int64 `json:"courierTransmissionPlanningDate"`
	CourierTransmissionDate         int64 `json:"courierTransmissionDate"`

	DeliveryCost          int `json:"deliveryCost"`
	DeliveryCostForSeller int `json:"deliveryCostForSeller"`

	Waybill       string `json:"waybill"`
	WaybillNumber string `json:"waybillNumber"`

	Category string `json:"category"`
}
