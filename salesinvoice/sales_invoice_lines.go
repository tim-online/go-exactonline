package salesinvoice

import (
	"context"
	"net/http"

	"github.com/tim-online/go-exactonline/odata"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	SalesInvoiceLinesEndpoint = "/v1/{division}/salesinvoice/SalesInvoiceLines"
)

// SalesInvoiceLines endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesInvoiceSalesInvoiceLines

func (s *Service) SalesInvoiceLinesGet(requestParams *SalesInvoiceLinesGetParams, ctx context.Context) (*SalesInvoiceLinesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewSalesInvoiceLinesGetResponse()
	path := s.rest.SubPath(SalesInvoiceLinesEndpoint)

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

func (s *Service) NewSalesInvoiceLinesGetResponse() *SalesInvoiceLinesGetResponse {
	return &SalesInvoiceLinesGetResponse{}
}

type SalesInvoiceLinesGetResponse struct {
	Results SalesInvoiceLines `json:"results"`
}

func (s *Service) NewSalesInvoiceLinesGetParams() *SalesInvoiceLinesGetParams {
	selectFields, _ := utils.Fields(&SalesInvoice{})
	return &SalesInvoiceLinesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type SalesInvoiceLinesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
