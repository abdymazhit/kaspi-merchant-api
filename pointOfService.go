package kaspi_merchant

import (
	"context"
	"fmt"
	"kaspi-merchant/vo"
)

type PointOfServiceResponse struct {
	Data struct {
		Id            string                      `json:"id"`
		Type          string                      `json:"type"`
		Attributes    vo.PointOfServiceAttributes `json:"attributes"`
		Relationships struct {
			City vo.Relationship `json:"city"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetPointOfServiceByOrderEntryId(ctx context.Context, orderEntryId string) (*PointOfServiceResponse, error) {
	var response PointOfServiceResponse
	if err := a.do(ctx, GET, fmt.Sprintf("orderentries/%s/deliveryPointOfService", orderEntryId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *api) GetPointOfService(ctx context.Context, pointOfServiceId string) (*PointOfServiceResponse, error) {
	var response PointOfServiceResponse
	if err := a.do(ctx, GET, fmt.Sprintf("pointofservices/%s", pointOfServiceId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
