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
	ErrApprovedDateGe = errors.New("filter merchant reviews approved date ge must be greater than 0")
	ErrApprovedDateLe = errors.New("filter merchant reviews approved date le must be greater than 0")
)

type GetReviewsRequest struct {
	PageNumber int
	PageSize   int
	Filter     struct {
		MerchantReviews struct {
			ApprovedDateGe time.Duration
			ApprovedDateLe time.Duration
			Quality        valueobject.ReviewQuality
		}
	}
}

func (p *GetReviewsRequest) ToUrlValues() (url.Values, error) {
	if p.PageNumber <= 0 {
		return nil, ErrPageNumber
	}
	if p.PageSize <= 0 {
		return nil, ErrPageSize
	}
	if p.PageSize > 100 {
		return nil, ErrPageSizeLimit
	}
	if p.Filter.MerchantReviews.ApprovedDateGe <= 0 {
		return nil, ErrApprovedDateGe
	}
	if p.Filter.MerchantReviews.ApprovedDateLe <= 0 {
		return nil, ErrApprovedDateLe
	}

	params := make(url.Values)
	params.Add("page[number]", strconv.Itoa(p.PageNumber))
	params.Add("page[size]", strconv.Itoa(p.PageSize))
	if p.Filter.MerchantReviews.Quality != "" {
		params.Add("filter[merchantreviews][quality]", string(p.Filter.MerchantReviews.Quality))
	}
	params.Add("filter[merchantreviews][approvedDate][$ge]", strconv.FormatInt(int64(p.Filter.MerchantReviews.ApprovedDateGe), 10))
	params.Add("filter[merchantreviews][approvedDate][$le]", strconv.FormatInt(int64(p.Filter.MerchantReviews.ApprovedDateLe), 10))
	return params, nil
}

type ReviewsResponse struct {
	Data struct {
		Id            string                       `json:"id"`
		Type          string                       `json:"type"`
		Attributes    valueobject.ReviewAttributes `json:"attributes"`
		Relationships struct {
			Order valueobject.Relationship `json:"order"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetReviews(ctx context.Context, authToken string, req GetReviewsRequest) (string, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return "", err
	}
	return w.do(ctx, "GET", "/merchantreviews", authToken, reqParams)
}

type ReviewResponse struct {
	Data struct {
		Id            string                       `json:"id"`
		Type          string                       `json:"type"`
		Attributes    valueobject.ReviewAttributes `json:"attributes"`
		Relationships struct {
			Order valueobject.Relationship `json:"order"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (w *wApi) GetReviewByOrderId(ctx context.Context, authToken, orderId string) (string, error) {
	return w.do(ctx, "GET", "/merchantreviews/"+orderId, authToken, nil)
}
