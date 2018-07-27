package financialtransaction

import (
	"context"
	"net/http"

	"github.com/tim-online/go-exactonline/edm"
)

// POST

func (s *Service) CashEntriesPost(body *CashEntriesPostBody, ctx context.Context) (*CashEntriesPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewCashEntriesPostResponse()
	path := s.rest.SubPath(CashEntriesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type CashEntriesPostBody NewCashEntry

func (s *Service) NewCashEntriesPostBody() *CashEntriesPostBody {
	return &CashEntriesPostBody{
		CashEntryLines: NewCashEntryLines{},
	}
}

func (s *Service) NewCashEntriesPostResponse() *CashEntriesPostResponse {
	return &CashEntriesPostResponse{}
}

type CashEntriesPostResponse struct{}

type NewCashEntry struct {
	CashEntryLines   NewCashEntryLines `json:"CashEntryLines"`   // Collection of lines
	ClosingBalanceFC edm.Double        `json:"ClosingBalanceFC"` // Closing balance in the currency of the transaction
	Currency         edm.String        `json:"Currency"`         // Currency code
	EntryNumber      edm.Int32         `json:"EntryNumber"`      // Entry number
	FinancialPeriod  edm.Int16         `json:"FinancialPeriod"`  // Fiancial period
	FinancialYear    edm.Int16         `json:"FinancialYear"`    // Fiancial year
	JournalCode      edm.String        `json:"JournalCode"`      // Code of Journal
	OpeningBalanceFC edm.Double        `json:"OpeningBalanceFC"` // Opening balance in the currency of the transaction
}

type NewCashEntryLines []NewCashEntry

type NewCashEntryLine struct {
	Account       edm.GUID     `json:"Account"`       // Reference to Account
	AmountFC      edm.Double   `json:"AmountFC"`      // Amount in the currency of the transaction
	AmountVATFC   edm.Double   `json:"AmountVATFC"`   // Vat amount in the currency of the transaction
	Asset         edm.GUID     `json:"Asset"`         // Reference to an asset
	CostCenter    edm.String   `json:"CostCenter"`    // Reference to a cost center
	CostUnit      edm.String   `json:"CostUnit"`      // Reference to a cost unit
	Date          edm.DateTime `json:"Date"`          // Date
	Description   edm.String   `json:"Description"`   // Description
	Document      edm.GUID     `json:"Document"`      // Reference to a document
	EntryID       edm.GUID     `json:"EntryID"`       // Reference to the header
	ExchangeRate  edm.Double   `json:"ExchangeRate"`  // Exchange rate
	GLAccount     edm.GUID     `json:"GLAccount"`     // General ledger account
	Notes         edm.String   `json:"Notes"`         // Extra remarks
	OurRef        edm.Int32    `json:"OurRef"`        // Invoice number
	Project       edm.GUID     `json:"Project"`       // Reference to a project
	Quantity      edm.Double   `json:"Quantity"`      // Quantity
	VATCode       edm.String   `json:"VATCode"`       // Reference to vat code
	VATPercentage edm.Double   `json:"VATPercentage"` // Vat code percentage
	VATType       edm.String   `json:"VATType"`       // Type of vat code
}
