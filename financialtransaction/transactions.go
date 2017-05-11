package financialtransaction

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tim-online/go-exactonline/edm"
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
	Results []Transaction `json:"results"`
}

type Transaction struct {
	EntryID                     edm.GUID          `json:"EntryID"`
	ClosingBalanceFC            edm.Double        `json:"ClosingBalanceFC"`
	Created                     edm.DateTime      `json:"Created"`
	Date                        edm.DateTime      `json:"Date"`
	Description                 edm.String        `json:"Description"`
	Division                    edm.Int32         `json:"Division"`
	Document                    edm.GUID          `json:"Document"`
	DocumentNumber              edm.Int32         `json:"DocumentNumber"`
	DocumentSubject             edm.String        `json:"DocumentSubject"`
	EntryNumber                 edm.Int32         `json:"EntryNumber"`
	ExternalLinkDescription     edm.String        `json:"ExternalLinkDescription"`
	ExternalLinkReference       edm.String        `json:"ExternalLinkReference"`
	FinancialPeriod             edm.Int16         `json:"FinancialPeriod"`
	FinancialYear               edm.Int16         `json:"FinancialYear"`
	IsExtraDuty                 edm.Boolean       `json:"IsExtraDuty"`
	JournalCode                 edm.String        `json:"JournalCode"`
	JournalDescription          edm.String        `json:"JournalDescription"`
	Modified                    edm.DateTime      `json:"Modified"`
	OpeningBalanceFC            edm.Double        `json:"OpeningBalanceFC"`
	PaymentConditionCode        edm.String        `json:"PaymentConditionCode"`
	PaymentConditionDescription edm.String        `json:"PaymentConditionDescription"`
	PaymentReference            edm.String        `json:"PaymentReference"`
	Status                      TransactionStatus `json:"Status"`
	StatusDescription           edm.String        `json:"StatusDescription"`
	TransactionLines            TransactionLines  `json:"TransactionLines"`
	Type                        TransactionType   `json:"Type"`
	TypeDescription             edm.String        `json:"TypeDescription"`
}

type TransactionLines []TransactionLine

func (t *TransactionLines) UnmarshalJSON(data []byte) (err error) {
	type Results TransactionLines

	type Envelope struct {
		Results Results `json:"results"`
	}

	envelope := &Envelope{Results: Results(*t)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*t = TransactionLines(envelope.Results)
	return nil
}

type TransactionLine struct {
	ID                        edm.GUID     `json:"ID"`
	Account                   edm.GUID     `json:"Account"`
	AccountCode               edm.String   `json:"AccountCode"`
	AccountName               edm.String   `json:"AccountName"`
	AmountDC                  edm.Double   `json:"AmountDC"`
	AmountFC                  edm.Double   `json:"AmountFC"`
	AmountVATBaseFC           edm.Double   `json:"AmountVATBaseFC"`
	AmountVATFC               edm.Double   `json:"AmountVATFC"`
	Asset                     edm.GUID     `json:"Asset"`
	AssetCode                 edm.String   `json:"AssetCode"`
	AssetDescription          edm.String   `json:"AssetDescription"`
	CostCenter                edm.String   `json:"CostCenter"`
	CostCenterDescription     edm.String   `json:"CostCenterDescription"`
	CostUnit                  edm.String   `json:"CostUnit"`
	CostUnitDescription       edm.String   `json:"CostUnitDescription"`
	Created                   edm.DateTime `json:"Created"`
	Creator                   edm.GUID     `json:"Creator"`
	CreatorFullName           edm.String   `json:"CreatorFullName"`
	Currency                  edm.String   `json:"Currency"`
	Date                      edm.DateTime `json:"Date"`
	Description               edm.String   `json:"Description"`
	Division                  edm.Int32    `json:"Division"`
	Document                  edm.GUID     `json:"Document"`
	DocumentNumber            edm.Int32    `json:"DocumentNumber"`
	DocumentSubject           edm.String   `json:"DocumentSubject"`
	DueDate                   edm.DateTime `json:"DueDate"`
	EntryID                   edm.GUID     `json:"EntryID"`
	EntryNumber               edm.Int32    `json:"EntryNumber"`
	ExchangeRate              edm.Double   `json:"ExchangeRate"`
	ExtraDutyAmountFC         edm.Double   `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage       edm.Double   `json:"ExtraDutyPercentage"`
	FinancialPeriod           edm.Int16    `json:"FinancialPeriod"`
	FinancialYear             edm.Int16    `json:"FinancialYear"`
	GLAccount                 edm.GUID     `json:"GLAccount"`
	GLAccountCode             edm.String   `json:"GLAccountCode"`
	GLAccountDescription      edm.String   `json:"GLAccountDescription"`
	InvoiceNumber             edm.Int32    `json:"InvoiceNumber"`
	Item                      edm.GUID     `json:"Item"`
	ItemCode                  edm.String   `json:"ItemCode"`
	ItemDescription           edm.String   `json:"ItemDescription"`
	JournalCode               edm.String   `json:"JournalCode"`
	JournalDescription        edm.String   `json:"JournalDescription"`
	LineNumber                edm.Int32    `json:"LineNumber"`
	LineType                  edm.Int16    `json:"LineType"`
	Modified                  edm.DateTime `json:"Modified"`
	Modifier                  edm.GUID     `json:"Modifier"`
	ModifierFullName          edm.String   `json:"ModifierFullName"`
	Notes                     edm.String   `json:"Notes"`
	OffsetID                  edm.GUID     `json:"OffsetID"`
	OrderNumber               edm.Int32    `json:"OrderNumber"`
	PaymentDiscountAmount     edm.Double   `json:"PaymentDiscountAmount"`
	PaymentReference          edm.String   `json:"PaymentReference"`
	Project                   edm.GUID     `json:"Project"`
	ProjectCode               edm.String   `json:"ProjectCode"`
	ProjectDescription        edm.String   `json:"ProjectDescription"`
	Quantity                  edm.Double   `json:"Quantity"`
	SerialNumber              edm.String   `json:"SerialNumber"`
	Status                    edm.Int16    `json:"Status"`
	Subscription              edm.GUID     `json:"Subscription"`
	SubscriptionDescription   edm.String   `json:"SubscriptionDescription"`
	TrackingNumber            edm.String   `json:"TrackingNumber"`
	TrackingNumberDescription edm.String   `json:"TrackingNumberDescription"`
	Type                      edm.Int32    `json:"Type"`
	VATCode                   edm.String   `json:"VATCode"`
	VATCodeDescription        edm.String   `json:"VATCodeDescription"`
	VATPercentage             edm.Double   `json:"VATPercentage"`
	VATType                   edm.String   `json:"VATType"`
	YourRef                   edm.String   `json:"YourRef"`
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

type TransactionStatus edm.Int16
type TransactionType edm.Int32

func (t *TransactionType) String() string {
	switch int(*t) {
	case 10:
		return "Opening balance"
	case 20:
		return "Sales entry"
	case 21:
		return "Sales credit note"
	case 30:
		return "Purchase entry"
	case 31:
		return "Purchase credit note"
	case 40:
		return "Cash flow"
	case 50:
		return "VAT return"
	case 70:
		return "Asset - Depreciation"
	case 71:
		return "Asset - Investment"
	case 72:
		return "Asset - Revaluation"
	case 73:
		return "Asset - Transfer"
	case 74:
		return "Asset - Split"
	case 75:
		return "Asset - Discontinue"
	case 76:
		return "Asset - Sales"
	case 80:
		return "Revaluation"
	case 82:
		return "Exchange rate difference"
	case 83:
		return "Payment difference"
	case 84:
		return "Deferred revenue"
	case 85:
		return "Tracking number:Revaluation"
	case 86:
		return "Deferred cost"
	case 87:
		return "VAT on prepayment"
	case 90:
		return "Other"
	case 120:
		return "Delivery"
	case 121:
		return "Sales Return"
	case 130:
		return "Receipt"
	case 131:
		return "Purchase return"
	case 140:
		return "Shop order stock receipt"
	case 141:
		return "Shop order stock reversal"
	case 142:
		return "Issue to parent"
	case 145:
		return "Shop order time entry"
	case 146:
		return "Shop order time entry reversal"
	case 147:
		return "Shop order by-product receipt"
	case 148:
		return "Shop order by-product reversal"
	case 150:
		return "Requirement issue"
	case 151:
		return "Requirement reversal"
	case 152:
		return "Returned from parent"
	case 155:
		return "Subcontract Issue"
	case 156:
		return "Subcontract reversal"
	case 158:
		return "Shop order completed"
	case 162:
		return "Finish assembly"
	case 170:
		return "Payroll"
	case 180:
		return "Stock revaluation"
	case 181:
		return "Financial revaluation"
	case 195:
		return "Stock count"
	case 290:
		return "Correction entry"
	case 310:
		return "Period closing"
	case 320:
		return "Year end reflection"
	case 321:
		return "Year end costing"
	case 322:
		return "Year end profits to gross profit"
	case 323:
		return "Year end costs to gross profit"
	case 324:
		return "Year end tax"
	case 325:
		return "Year end gross profit to net p/l"
	case 326:
		return "Year end net p/l to balance sheet"
	case 327:
		return "Year end closing balance"
	case 328:
		return "Year start opening balance"
	case 3000:
		return "Budget"
	}
	return ""
}
