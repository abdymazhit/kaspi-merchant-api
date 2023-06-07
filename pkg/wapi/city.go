package wapi

import (
	"context"
	"kaspi-merchant/pkg/wapi/valueobject"
)

type CitiesResponse struct {
	Data []struct {
		Id            string                     `json:"id"`
		Type          string                     `json:"type"`
		Attributes    valueobject.CityAttributes `json:"attributes"`
		Relationships struct {
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetCities(ctx context.Context, authToken string) (string, error) {
	return w.do(ctx, "GET", "cities", authToken, nil)
}

type CityResponse struct {
	Data struct {
		Id            string                     `json:"id"`
		Type          string                     `json:"type"`
		Attributes    valueobject.CityAttributes `json:"attributes"`
		Relationships struct {
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetCityByPointOfServiceId(ctx context.Context, authToken, pointOfServiceId string) (string, error) {
	return w.do(ctx, "GET", "pointofservices/"+pointOfServiceId+"/city", authToken, nil)
}
