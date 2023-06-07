package kaspi_merchant

import (
	"context"
	"errors"
	"fmt"
	"kaspi-merchant/vo"
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
			Quality        vo.ReviewQuality
		}
	}
}

func (r *GetReviewsRequest) ToUrlValues() (url.Values, error) {
	if r.PageNumber <= 0 {
		return nil, ErrPageNumber
	}
	if r.PageSize <= 0 {
		return nil, ErrPageSize
	}
	if r.PageSize > 100 {
		return nil, ErrPageSizeLimit
	}
	if r.Filter.MerchantReviews.ApprovedDateGe <= 0 {
		return nil, ErrApprovedDateGe
	}
	if r.Filter.MerchantReviews.ApprovedDateLe <= 0 {
		return nil, ErrApprovedDateLe
	}

	params := make(url.Values)
	params.Add("page[number]", strconv.Itoa(r.PageNumber))
	params.Add("page[size]", strconv.Itoa(r.PageSize))
	if r.Filter.MerchantReviews.Quality != "" {
		params.Add("filter[merchantreviews][quality]", string(r.Filter.MerchantReviews.Quality))
	}
	params.Add("filter[merchantreviews][approvedDate][$ge]", strconv.FormatInt(int64(r.Filter.MerchantReviews.ApprovedDateGe), 10))
	params.Add("filter[merchantreviews][approvedDate][$le]", strconv.FormatInt(int64(r.Filter.MerchantReviews.ApprovedDateLe), 10))
	return params, nil
}

type ReviewsResponse struct {
	Data struct {
		Id            string              `json:"id"`
		Type          string              `json:"type"`
		Attributes    vo.ReviewAttributes `json:"attributes"`
		Relationships struct {
			Order vo.Relationship `json:"order"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetReviews(ctx context.Context, req GetReviewsRequest) (*ReviewsResponse, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return nil, err
	}

	var response ReviewsResponse
	if err = a.do(ctx, GET, "merchantreviews", reqParams, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type ReviewResponse struct {
	Data struct {
		Id            string              `json:"id"`
		Type          string              `json:"type"`
		Attributes    vo.ReviewAttributes `json:"attributes"`
		Relationships struct {
			Order vo.Relationship `json:"order"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (a *api) GetReviewByOrderId(ctx context.Context, orderId string) (*ReviewResponse, error) {
	var response ReviewResponse
	if err := a.do(ctx, GET, fmt.Sprintf("merchantreviews/%s", orderId), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
