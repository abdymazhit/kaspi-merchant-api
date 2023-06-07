package kaspi_merchant

import (
	"context"
	"fmt"
)

type ProductResponse struct {
	Data struct {
		Id            string            `json:"id"`
		Type          string            `json:"type"`
		Attributes    ProductAttributes `json:"attributes"`
		Relationships struct {
			MerchantProduct Relationship `json:"merchantProduct"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetProductByOrderEntryId(ctx context.Context, orderEntryId string) (*ProductResponse, error) {
	var response ProductResponse
	if err := a.do(ctx, GET, fmt.Sprintf("orderentries/%s/product", orderEntryId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *api) GetProductByMasterProductId(ctx context.Context, masterProductId string) (*ProductResponse, error) {
	var response ProductResponse
	if err := a.do(ctx, GET, fmt.Sprintf("masterproducts/%s/merchantProduct", masterProductId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *api) GetProduct(ctx context.Context, productId string) (*ProductResponse, error) {
	var response ProductResponse
	if err := a.do(ctx, GET, fmt.Sprintf("orderentries/%s/product", productId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
