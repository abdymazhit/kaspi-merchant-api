package kaspi_merchant

import (
	"context"
	"errors"
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
			CreationDateGe    time.Time
			CreationDateLe    time.Time
			State             OrdersState
			Status            OrdersStatus
			DeliveryType      OrdersDeliveryType
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
	if p.Filter.Orders.CreationDateGe.Unix() <= 0 {
		return nil, ErrFilterOrdersCreationDateGe
	}
	if p.Filter.Orders.State == "" {
		return nil, ErrFilterOrdersState
	}
	if p.Filter.Orders.State == OrdersStatePickup && p.Filter.Orders.DeliveryType == OrdersDeliveryTypePickup {
		return nil, ErrFilterOrdersStatePickupDeliveryTypePickup
	}
	if p.Filter.Orders.State == OrdersStateDelivery && p.Filter.Orders.DeliveryType == OrdersDeliveryTypeDelivery {
		return nil, ErrFilterOrdersStateDeliveryDeliveryTypeDelivery
	}
	if p.Filter.Orders.SignatureRequired && p.Filter.Orders.State != OrdersStateSignRequired {
		return nil, ErrFilterOrdersStateSignRequired
	}

	params := make(url.Values)
	params.Add("page[number]", strconv.Itoa(p.PageNumber))
	params.Add("page[size]", strconv.Itoa(p.PageSize))
	params.Add("filter[orders][creationDate][$ge]", strconv.FormatInt(p.Filter.Orders.CreationDateGe.Unix(), 10))
	if p.Filter.Orders.CreationDateLe.Unix() > 0 {
		params.Add("filter[orders][creationDate][$le]", strconv.FormatInt(p.Filter.Orders.CreationDateLe.Unix(), 10))
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
		Id            string          `json:"id"`
		Type          string          `json:"type"`
		Attributes    OrderAttributes `json:"attributes"`
		Relationships struct {
			Entries EntriesRelationship `json:"entries"`
			User    Relationship        `json:"user"`
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

func (a *api) GetOrders(ctx context.Context, req GetOrdersRequest) (*OrdersResponse, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return nil, err
	}

	var response OrdersResponse
	if err = a.do(ctx, GET, "orders", reqParams, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type OrderByCodeResponse struct {
	Data []struct {
		Id            string          `json:"id"`
		Type          string          `json:"type"`
		Attributes    OrderAttributes `json:"attributes"`
		Relationships struct {
			Entries EntriesRelationship `json:"entries"`
			User    Relationship        `json:"user"`
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

func (a *api) GetOrderByCode(ctx context.Context, code string) (*OrderByCodeResponse, error) {
	var response OrderByCodeResponse
	if err := a.do(ctx, GET, "orders", url.Values{"filter[orders][code]": []string{code}}, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// --------------------------------------------------
// VO
// --------------------------------------------------

type OrderDeliveryMode string

const (
	OrderDeliveryModeLocal          OrderDeliveryMode = "DELIVERY_LOCAL"
	OrderDeliveryModePickup         OrderDeliveryMode = "DELIVERY_PICKUP"
	OrderDeliveryModeRegionalPickup OrderDeliveryMode = "DELIVERY_REGIONAL_PICKUP"
	OrderDeliveryModeRegionalTodoor OrderDeliveryMode = "DELIVERY_REGIONAL_TODOOR"
)

type OrdersDeliveryType string

const (
	OrdersDeliveryTypePickup   OrdersDeliveryType = "PICKUP"
	OrdersDeliveryTypeDelivery OrdersDeliveryType = "DELIVERY"
)

type OrderPaymentMode string

const (
	OrderPaymentModePayWithCredit OrderPaymentMode = "PAY_WITH_CREDIT"
	OrderPaymentModePrepaid       OrderPaymentMode = "PREPAID"
)

type OrdersState string

const (
	OrdersStateNew           OrdersState = "NEW"
	OrdersStateSignRequired  OrdersState = "SIGN_REQUIRED"
	OrdersStatePickup        OrdersState = "PICKUP"
	OrdersStateDelivery      OrdersState = "DELIVERY"
	OrdersStateKaspiDelivery OrdersState = "KASPI_DELIVERY"
	OrdersStateArchive       OrdersState = "ARCHIVE"
)

type OrdersStatus string

const (
	OrdersStatusApprovedByBank               OrdersStatus = "APPROVED_BY_BANK"
	OrdersStatusAcceptedByMerchant           OrdersStatus = "ACCEPTED_BY_MERCHANT"
	OrdersStatusCompleted                    OrdersStatus = "COMPLETED"
	OrdersStatusCancelled                    OrdersStatus = "CANCELLED"
	OrdersStatusCancelling                   OrdersStatus = "CANCELLING"
	OrdersStatusKaspiDeliveryReturnRequested OrdersStatus = "KASPI_DELIVERY_RETURN_REQUESTED"
	OrdersStatusReturnAcceptedByMerchant     OrdersStatus = "RETURN_ACCEPTED_BY_MERCHANT"
	OrdersStatusReturned                     OrdersStatus = "RETURNED"
)
