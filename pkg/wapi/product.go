package wapi

import (
	"context"
	"kaspi-merchant/pkg/wapi/valueobject"
)

type ProductResponse struct {
	Data struct {
		Id            string                        `json:"id"`
		Type          string                        `json:"type"`
		Attributes    valueobject.ProductAttributes `json:"attributes"`
		Relationships struct {
			MerchantProduct valueobject.Relationship `json:"merchantProduct"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetProductByOrderEntryId(ctx context.Context, authToken, orderEntryId string) (string, error) {
	return w.do(ctx, "GET", "/orderentries/"+orderEntryId+"/product", authToken, nil)
}

func (w *wApi) GetProductByMasterProductId(ctx context.Context, authToken, masterProductId string) (string, error) {
	return w.do(ctx, "GET", "/masterproducts/"+masterProductId+"/merchantProduct", authToken, nil)
}

func (w *wApi) GetProduct(ctx context.Context, authToken, productId string) (string, error) {
	return w.do(ctx, "GET", "/orderentries/"+productId+"/product", authToken, nil)
}
