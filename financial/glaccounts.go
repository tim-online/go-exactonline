package financial

import (
	"context"
	"net/http"

	"github.com/tim-online/go-exactonline/odata"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	GLAccountsEndpoint = "/v1/{division}/financial/GLAccounts"
)

// GLAccounts endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialGLAccounts

func (s *Service) GLAccountsGet(requestParams *GLAccountsGetParams, ctx context.Context) (*GLAccountsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewGLAccountsGetResponse()
	path := s.rest.SubPath(GLAccountsEndpoint)

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

func (s *Service) NewGLAccountsGetResponse() *GLAccountsGetResponse {
	return &GLAccountsGetResponse{}
}

type GLAccountsGetResponse struct {
	Results GLAccounts `json:"results"`
}

func (s *Service) NewGLAccountsGetParams() *GLAccountsGetParams {
	selectFields, _ := utils.Fields(&GLAccount{})
	return &GLAccountsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type GLAccountsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
