package financialtransaction

import (
	"context"
	"fmt"

	"github.com/tim-online/go-exactonline/odata"
	"github.com/tim-online/go-exactonline/utils"
)

const (
	TransactionsEndpoint = "/v1/%d/financialtransaction/Transactions"

	TypeOpeningBalance              TransactionType = 10
	TypeSalesEntry                  TransactionType = 20
	TypeSalesCreditNote             TransactionType = 21
	TypePurchaseEntry               TransactionType = 30
	TypePurchaseCreditNote          TransactionType = 31
	TypeCashFlow                    TransactionType = 40
	TypeVatReturn                   TransactionType = 50
	TypeAssetDepreciation           TransactionType = 70
	TypeAssetInvestment             TransactionType = 71
	TypeAssetRevaluation            TransactionType = 72
	TypeAssetTransfer               TransactionType = 73
	TypeAssetSplit                  TransactionType = 74
	TypeAssetDiscontinue            TransactionType = 75
	TypeAssetSales                  TransactionType = 76
	TypeRevaluation                 TransactionType = 80
	TypeExchangeRateDifference      TransactionType = 82
	TypePaymentDifference           TransactionType = 83
	TypeDeferredRevenue             TransactionType = 84
	TypeTrackingNumberRevaluation   TransactionType = 85
	TypeDeferredCost                TransactionType = 86
	TypeVatOnPrepayment             TransactionType = 87
	TypeOther                       TransactionType = 90
	TypeDelivery                    TransactionType = 120
	TypeSalesReturn                 TransactionType = 121
	TypeReceipt                     TransactionType = 130
	TypePurchaseReturn              TransactionType = 131
	TypeShopOrderStockReceipt       TransactionType = 140
	TypeShopOrderStockReversal      TransactionType = 141
	TypeIssueToParent               TransactionType = 142
	TypeShopOrderTimeEntry          TransactionType = 145
	TypeShopOrderTimeEntryReversal  TransactionType = 146
	TypeShopOrderByProductReceipt   TransactionType = 147
	TypeShopOrderByProductReversal  TransactionType = 148
	TypeRequirementIssue            TransactionType = 150
	TypeRequirementReversal         TransactionType = 151
	TypeReturnedFromParent          TransactionType = 152
	TypeSubcontractIssue            TransactionType = 155
	TypeSubcontractReversal         TransactionType = 156
	TypeShopOrderCompleted          TransactionType = 158
	TypeFinishAssembly              TransactionType = 162
	TypePayroll                     TransactionType = 170
	TypeStockRevaluation            TransactionType = 180
	TypeFinancialRevaluation        TransactionType = 181
	TypeStockCount                  TransactionType = 195
	TypeCorrectionEntry             TransactionType = 290
	TypePeriodClosing               TransactionType = 310
	TypeYearEndReflection           TransactionType = 320
	TypeYearEndCosting              TransactionType = 321
	TypeYearEndProfitsToGrossProfit TransactionType = 322
	TypeYearEndCostsToGrossProfit   TransactionType = 323
	TypeYearEndTax                  TransactionType = 324
	TypeYearEndGrossProfitToNetPL   TransactionType = 325
	TypeYearEndNetPLToBalanceSheet  TransactionType = 326
	TypeYearEndClosingBalance       TransactionType = 327
	TypeYearStartOpeningBalance     TransactionType = 328
	TypeBudget                      TransactionType = 3000

	StatusRejected  TransactionStatus = 5
	StatusOpen      TransactionStatus = 20
	StatusProcessed TransactionStatus = 50
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

type TransactionsGetResponse struct {
	Results Transactions `json:"results"`
}

func (s *Service) NewTransactionsGetParams() *TransactionsGetParams {
	selectFields, _ := utils.Fields(&Transaction{})
	expandFields := []string{"TransactionLines"}
	return &TransactionsGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type TransactionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
