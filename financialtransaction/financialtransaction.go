package financialtransaction

import (
	"context"
	"fmt"

	"github.com/tim-online/go-exactonline/rest"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	TransactionsEndpoint = "/v1/%d/financialtransaction/Transactions"
)

func NewService(rest *rest.Client) *Service {
	return &Service{rest: rest}
}

type Service struct {
	rest *rest.Client
}

// Transactions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionTransactions

func (s *Service) Transactions(divisionID int, requestParams *TransactionsGetParams, ctx context.Context) (*TransactionsGetResponse, error) {
	method := "GET"
	responseBody := s.NewTransactionsGetResponse()
	path := fmt.Sprintf(TransactionsEndpoint)

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

func (s *Service) NewTransactionsGetResponse() *TransactionsGetResponse {
	return &TransactionsGetResponse{}
}

type TransactionsGetResponse struct {
}

func (s *Service) NewTransactionsGetParams() *TransactionsGetParams {
	return &TransactionsGetParams{}
}

type TransactionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select string `schema:"$select"`
	Filter string `schema:"$filter"`
}
