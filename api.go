package kaspi_merchant

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://kaspi.kz/shop/api/v2"

var httpClient = &http.Client{}

type method string

const (
	GET  method = "GET"
	POST method = "POST"
)

type API interface {
	GetCities(ctx context.Context) (*CitiesResponse, error)
	GetCityByPointOfServiceId(ctx context.Context, pointOfServiceId string) (*CityResponse, error)

	GetOrders(ctx context.Context, req GetOrdersRequest) (*OrdersResponse, error)
	GetOrderByCode(ctx context.Context, code string) (*OrderByCodeResponse, error)

	GetOrderEntries(ctx context.Context, orderId string) (*OrderEntriesResponse, error)
	GetOrderEntry(ctx context.Context, entryId string) (*OrderEntryResponse, error)

	GetPointOfServiceByOrderEntryId(ctx context.Context, orderEntryId string) (*PointOfServiceResponse, error)
	GetPointOfService(ctx context.Context, pointOfServiceId string) (*PointOfServiceResponse, error)

	GetProductByOrderEntryId(ctx context.Context, orderEntryId string) (*ProductResponse, error)
	GetProductByMasterProductId(ctx context.Context, masterProductId string) (*ProductResponse, error)
	GetProduct(ctx context.Context, productId string) (*ProductResponse, error)

	GetReviews(ctx context.Context, req GetReviewsRequest) (*ReviewsResponse, error)
	GetReviewByOrderId(ctx context.Context, orderId string) (*ReviewResponse, error)
}

type api struct {
	authToken string
}

func New(authToken string) API {
	return &api{authToken: authToken}
}

func (a *api) do(ctx context.Context, method method, path string, params url.Values, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, string(method), fmt.Sprintf("%s/%s", baseURL, path), nil)
	if err != nil {
		return err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("X-Auth-Token", a.authToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	if _, err = io.Copy(io.Discard, resp.Body); err != nil {
		return err
	}

	return nil
}
