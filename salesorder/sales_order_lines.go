package salesorder

import (
	"context"
	"net/http"

	"github.com/tim-online/go-exactonline/odata"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	SalesOrderLinesEndpoint = "/v1/{division}/salesorder/SalesOrderLines"
)

// SalesOrderLines endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderSalesOrderLines

func (s *Service) SalesOrderLinesGet(requestParams *SalesOrderLinesGetParams, ctx context.Context) (*SalesOrderLinesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewSalesOrderLinesGetResponse()
	path := s.rest.SubPath(SalesOrderLinesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	utils.AddQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewSalesOrderLinesGetResponse() *SalesOrderLinesGetResponse {
	return &SalesOrderLinesGetResponse{}
}

type SalesOrderLinesGetResponse struct {
	Results SalesOrderLines `json:"results"`
}

func (s *Service) NewSalesOrderLinesGetParams() *SalesOrderLinesGetParams {
	selectFields, _ := utils.Fields(&SalesOrder{})
	expandFields := []string{"SalesOrderLines"}
	return &SalesOrderLinesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type SalesOrderLinesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
