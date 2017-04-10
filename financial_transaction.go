package exact

const (
	TransactionsEndpoint = "/v1/%d/financialtransaction/Transactions"
)

func NewFinancialTransactionService(client *Client) *FinancialTransactionService {
	return &FinancialTransactionService{Client: client}
}

type FinancialTransactionService struct {
	Client *Client
}

// Transactions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionTransactions

func (s *FinancialTransactionService) Transactions(divisionID int, requestParams *TransactionsGetParams, ctx context.Context) (*TransactionsGetResponse, error) {
	method := "GET"
	responseBody := s.NewTransactionsGetResponse()
	path := fmt.Sprintf(TransactionsEndpoint)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}
