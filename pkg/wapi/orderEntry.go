package wapi

import (
	"context"
	"kaspi-merchant/pkg/wapi/valueobject"
)

type OrderEntriesResponse struct {
	Data []struct {
		Id            string                           `json:"id"`
		Type          string                           `json:"type"`
		Attributes    valueobject.OrderEntryAttributes `json:"attributes"`
		Relationships struct {
			Product                valueobject.Relationship `json:"product"`
			DeliveryPointOfService valueobject.Relationship `json:"deliveryPointOfService"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetOrderEntries(ctx context.Context, authToken, orderId string) (string, error) {
	return w.do(ctx, "GET", "orders/"+orderId+"/entries", authToken, nil)
}

type OrderEntryResponse struct {
	Data struct {
		Id            string                           `json:"id"`
		Type          string                           `json:"type"`
		Attributes    valueobject.OrderEntryAttributes `json:"attributes"`
		Relationships struct {
			Product                valueobject.Relationship `json:"product"`
			DeliveryPointOfService valueobject.Relationship `json:"deliveryPointOfService"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetOrderEntry(ctx context.Context, authToken, entryId string) (string, error) {
	return w.do(ctx, "GET", "orderentries/"+entryId, authToken, nil)
}
