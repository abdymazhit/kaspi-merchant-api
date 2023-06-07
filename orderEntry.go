package kaspi_merchant

import (
	"context"
	"fmt"
)

type OrderEntriesResponse struct {
	Data []struct {
		Id            string               `json:"id"`
		Type          string               `json:"type"`
		Attributes    OrderEntryAttributes `json:"attributes"`
		Relationships struct {
			Product                Relationship `json:"product"`
			DeliveryPointOfService Relationship `json:"deliveryPointOfService"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetOrderEntries(ctx context.Context, orderId string) (*OrderEntriesResponse, error) {
	var response OrderEntriesResponse
	if err := a.do(ctx, GET, fmt.Sprintf("orders/%s/entries", orderId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type OrderEntryResponse struct {
	Data struct {
		Id            string               `json:"id"`
		Type          string               `json:"type"`
		Attributes    OrderEntryAttributes `json:"attributes"`
		Relationships struct {
			Product                Relationship `json:"product"`
			DeliveryPointOfService Relationship `json:"deliveryPointOfService"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetOrderEntry(ctx context.Context, entryId string) (*OrderEntryResponse, error) {
	var response OrderEntryResponse
	if err := a.do(ctx, GET, fmt.Sprintf("orderentries/%s", entryId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
