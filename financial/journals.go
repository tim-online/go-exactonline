package financial

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tim-online/go-exactonline/odata"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	JournalsEndpoint = "/v1/%d/financial/Journals"
)

// Journals endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialJournals

func (s *Service) JournalsGet(divisionID int, requestParams *JournalsGetParams, ctx context.Context) (*JournalsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewJournalsGetResponse()
	path := fmt.Sprintf(JournalsEndpoint, divisionID)

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

func (s *Service) NewJournalsGetResponse() *JournalsGetResponse {
	return &JournalsGetResponse{}
}

type JournalsGetResponse struct {
	Results Journals `json:"results"`
}

func (s *Service) NewJournalsGetParams() *JournalsGetParams {
	selectFields, _ := utils.Fields(&Journal{})
	return &JournalsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type JournalsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
