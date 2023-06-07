package wapi

import (
	"context"
	"kaspi-merchant/pkg/wapi/valueobject"
)

type PointOfServiceResponse struct {
	Data struct {
		Id            string                               `json:"id"`
		Type          string                               `json:"type"`
		Attributes    valueobject.PointOfServiceAttributes `json:"attributes"`
		Relationships struct {
			City valueobject.Relationship `json:"city"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetPointOfServiceByOrderEntryId(ctx context.Context, authToken, orderEntryId string) (string, error) {
	return w.do(ctx, "GET", "/orderentries/"+orderEntryId+"/deliveryPointOfService", authToken, nil)
}

func (w *wApi) GetPointOfService(ctx context.Context, authToken, pointOfServiceId string) (string, error) {
	return w.do(ctx, "GET", "/pointofservices/"+pointOfServiceId, authToken, nil)
}
