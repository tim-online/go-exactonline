package financialtransaction

import (
	"context"
	"fmt"

	"github.com/tim-online/go-exactonline/edm"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	TransactionsEndpoint = "/v1/%d/financialtransaction/Transactions"
)

// Transactions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionTransactions

func (s *Service) TransactionsGet(divisionID int, requestParams *TransactionsGetParams, ctx context.Context) (*TransactionsGetResponse, error) {
	method := "GET"
	responseBody := s.NewTransactionsGetResponse()
	path := fmt.Sprintf(TransactionsEndpoint, divisionID)

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

type TransactionsGetResponse []Transaction

type Transaction struct {
	EntryID                     edm.GUID         `json:"EntryID"`
	ClosingBalanceFC            edm.Double       `json:"ClosingBalanceFC"`
	Created                     edm.DateTime     `json:"Created"`
	Date                        edm.DateTime     `json:"Date"`
	Description                 edm.String       `json:"Description"`
	Division                    edm.Int32        `json:"Division"`
	Document                    edm.GUID         `json:"Document"`
	DocumentNumber              edm.Int32        `json:"DocumentNumber"`
	DocumentSubject             edm.String       `json:"DocumentSubject"`
	EntryNumber                 edm.Int32        `json:"EntryNumber"`
	ExternalLinkDescription     edm.String       `json:"ExternalLinkDescription"`
	ExternalLinkReference       edm.String       `json:"ExternalLinkReference"`
	FinancialPeriod             edm.Int16        `json:"FinancialPeriod"`
	FinancialYear               edm.Int16        `json:"FinancialYear"`
	IsExtraDuty                 edm.Boolean      `json:"IsExtraDuty"`
	JournalCode                 edm.String       `json:"JournalCode"`
	JournalDescription          edm.String       `json:"JournalDescription"`
	Modified                    edm.DateTime     `json:"Modified"`
	OpeningBalanceFC            edm.Double       `json:"OpeningBalanceFC"`
	PaymentConditionCode        edm.String       `json:"PaymentConditionCode"`
	PaymentConditionDescription edm.String       `json:"PaymentConditionDescription"`
	PaymentReference            edm.String       `json:"PaymentReference"`
	Status                      edm.Int16        `json:"Status"`
	StatusDescription           edm.String       `json:"StatusDescription"`
	TransactionLines            TransactionLines `json:"TransactionLines"`
	Type                        edm.Int32        `json:"Type"`
	TypeDescription             edm.String       `json:"TypeDescription"`
}

type TransactionLines []TransactionLine
type TransactionLine struct {
}

func (s *Service) NewTransactionsGetParams() *TransactionsGetParams {
	return &TransactionsGetParams{}
}

type TransactionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select string `schema:"$select"`
	Filter string `schema:"$filter"`
	Top    int    `schema:"$top"`
}
