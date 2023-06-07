package wapi

import (
	"context"
	"errors"
	"kaspi-merchant/pkg/wapi/valueobject"
	"net/url"
	"strconv"
	"time"
)

var (
	ErrPageNumber                                    = errors.New("page number must be greater than 0")
	ErrPageSize                                      = errors.New("page size must be greater than 0")
	ErrPageSizeLimit                                 = errors.New("page size limit is 100")
	ErrFilterOrdersState                             = errors.New("filter orders state must be set")
	ErrFilterOrdersCreationDateGe                    = errors.New("filter orders creation date ge must be greater than 0")
	ErrFilterOrdersStatePickupDeliveryTypePickup     = errors.New("filter orders state pickup and delivery type pickup are incompatible")
	ErrFilterOrdersStateDeliveryDeliveryTypeDelivery = errors.New("filter orders state delivery and delivery type delivery are incompatible")
	ErrFilterOrdersStateSignRequired                 = errors.New("filter orders state sign required cannot be used without filter orders state sign required")
)

type GetOrdersRequest struct {
	PageNumber int
	PageSize   int
	Filter     struct {
		Orders struct {
			CreationDateGe    time.Duration
			CreationDateLe    time.Duration
			State             valueobject.OrdersState
			Status            valueobject.OrdersStatus
			DeliveryType      valueobject.OrdersDeliveryType
			SignatureRequired bool
		}
	}
	Include struct {
		Orders string
	}
}

func (p *GetOrdersRequest) ToUrlValues() (url.Values, error) {
	if p.PageNumber <= 0 {
		return nil, ErrPageNumber
	}
	if p.PageSize <= 0 {
		return nil, ErrPageSize
	}
	if p.PageSize > 100 {
		return nil, ErrPageSizeLimit
	}
	if p.Filter.Orders.CreationDateGe <= 0 {
		return nil, ErrFilterOrdersCreationDateGe
	}
	if p.Filter.Orders.State == "" {
		return nil, ErrFilterOrdersState
	}
	if p.Filter.Orders.State == valueobject.OrdersStatePickup && p.Filter.Orders.DeliveryType == valueobject.OrdersDeliveryTypePickup {
		return nil, ErrFilterOrdersStatePickupDeliveryTypePickup
	}
	if p.Filter.Orders.State == valueobject.OrdersStateDelivery && p.Filter.Orders.DeliveryType == valueobject.OrdersDeliveryTypeDelivery {
		return nil, ErrFilterOrdersStateDeliveryDeliveryTypeDelivery
	}
	if p.Filter.Orders.SignatureRequired && p.Filter.Orders.State != valueobject.OrdersStateSignRequired {
		return nil, ErrFilterOrdersStateSignRequired
	}

	params := make(url.Values)
	params.Add("page[number]", strconv.Itoa(p.PageNumber))
	params.Add("page[size]", strconv.Itoa(p.PageSize))
	params.Add("filter[orders][creationDate][$ge]", strconv.FormatInt(int64(p.Filter.Orders.CreationDateGe), 10))
	if p.Filter.Orders.CreationDateLe > 0 {
		params.Add("filter[orders][creationDate][$le]", strconv.FormatInt(int64(p.Filter.Orders.CreationDateLe), 10))
	}
	params.Add("filter[orders][state]", string(p.Filter.Orders.State))
	if p.Filter.Orders.Status != "" {
		params.Add("filter[orders][status]", string(p.Filter.Orders.Status))
	}
	if p.Filter.Orders.DeliveryType != "" {
		params.Add("filter[orders][deliveryType]", string(p.Filter.Orders.DeliveryType))
	}
	if p.Filter.Orders.SignatureRequired {
		params.Add("filter[orders][signatureRequired]", strconv.FormatBool(p.Filter.Orders.SignatureRequired))
	}
	//if p.Include.Orders != "" {
	//	params.Add("include[orders]", p.Include.Orders)
	//}
	return params, nil
}

type OrdersResponse struct {
	Data []struct {
		Id            string                      `json:"id"`
		Type          string                      `json:"type"`
		Attributes    valueobject.OrderAttributes `json:"attributes"`
		Relationships struct {
			Entries valueobject.EntriesRelationship `json:"entries"`
			User    valueobject.Relationship        `json:"user"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			CellPhone string `json:"cellPhone"`
		} `json:"attributes"`
		Relationships struct {
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"included"`
	Meta struct {
		PageCount  int `json:"pageCount"`
		TotalCount int `json:"totalCount"`
	} `json:"meta"`
}

func (w *wApi) GetOrders(ctx context.Context, authToken string, req GetOrdersRequest) (string, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return "", err
	}
	return w.do(ctx, "GET", "orders", authToken, reqParams)
}

type OrderByCodeResponse struct {
	Data []struct {
		Id            string                      `json:"id"`
		Type          string                      `json:"type"`
		Attributes    valueobject.OrderAttributes `json:"attributes"`
		Relationships struct {
			Entries valueobject.EntriesRelationship `json:"entries"`
			User    valueobject.Relationship        `json:"user"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
	Meta     struct {
		PageCount  int `json:"pageCount"`
		TotalCount int `json:"totalCount"`
	} `json:"meta"`
}

func (w *wApi) GetOrderByCode(ctx context.Context, authToken, code string) (string, error) {
	return w.do(ctx, "GET", "orders", authToken, url.Values{
		"filter[orders][code]": []string{code},
	})
}
