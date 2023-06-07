package wapi

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type WApi interface {
}

type wApi struct {
	baseUrl    string
	httpClient *http.Client
}

func New() WApi {
	return &wApi{
		baseUrl:    "https://kaspi.kz/shop/api/v2",
		httpClient: &http.Client{},
	}
}

type Response struct {
	Data          []interface{} `json:"data"`
	Type          string        `json:"type"`
	Id            string        `json:"id"`
	Attributes    interface{}   `json:"attributes"`
	Relationships interface{}   `json:"relationships"`
	Links         interface{}   `json:"links"`
	Included      interface{}   `json:"included"`
	Meta          interface{}   `json:"meta"`
}

func (w *wApi) do(ctx context.Context, method, path, authToken string, params url.Values) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, w.baseUrl+"/"+path, nil)
	if err != nil {
		return "", err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("X-Auth-Token", authToken)

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
