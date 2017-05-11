package financialtransaction

import (
	"context"
	"fmt"

	"github.com/tim-online/go-exactonline/odata"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	BankEntriesEndpoint = "/v1/%d/financialtransaction/BankEntries"
)

// BankEntries endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionBankEntries

func (s *Service) BankEntriesGet(divisionID int, requestParams *BankEntriesGetParams, ctx context.Context) (*BankEntriesGetResponse, error) {
	method := "GET"
	responseBody := s.NewBankEntriesGetResponse()
	path := fmt.Sprintf(BankEntriesEndpoint, divisionID)

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

func (s *Service) NewBankEntriesGetResponse() *BankEntriesGetResponse {
	return &BankEntriesGetResponse{}
}

type BankEntriesGetResponse struct {
	Results BankEntries `json:"results"`
}

func (s *Service) NewBankEntriesGetParams() *BankEntriesGetParams {
	selectFields, _ := utils.Fields(&BankEntry{})
	expandFields := []string{"BankEntryLines"}
	return &BankEntriesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type BankEntriesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
