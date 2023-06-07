package kaspi_merchant

import (
	"context"
	"fmt"
	"kaspi-merchant/vo"
)

type CitiesResponse struct {
	Data []struct {
		Id            string            `json:"id"`
		Type          string            `json:"type"`
		Attributes    vo.CityAttributes `json:"attributes"`
		Relationships struct {
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetCities(ctx context.Context) (*CitiesResponse, error) {
	var response CitiesResponse
	if err := a.do(ctx, GET, "cities", nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type CityResponse struct {
	Data struct {
		Id            string            `json:"id"`
		Type          string            `json:"type"`
		Attributes    vo.CityAttributes `json:"attributes"`
		Relationships struct {
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetCityByPointOfServiceId(ctx context.Context, pointOfServiceId string) (*CityResponse, error) {
	var response CityResponse
	if err := a.do(ctx, GET, fmt.Sprintf("pointofservices/%s/city", pointOfServiceId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
