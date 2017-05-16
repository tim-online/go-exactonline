package salesinvoice

import (
	"encoding/json"

	"github.com/tim-online/go-exactonline/edm"
	"github.com/tim-online/go-exactonline/utils"
)

type SalesInvoices []SalesInvoice

type SalesInvoice struct {
	InvoiceID                            edm.GUID          `json:"InvoiceID"`
	AmountDC                             edm.Double        `json:"AmountDC"`
	AmountFC                             edm.Double        `json:"AmountFC"`
	Created                              edm.DateTime      `json:"Created"`
	Creator                              edm.GUID          `json:"Creator"`
	CreatorFullName                      edm.String        `json:"CreatorFullName"`
	Currency                             edm.String        `json:"Currency"`
	DeliverTo                            edm.GUID          `json:"DeliverTo"`
	DeliverToAddress                     edm.GUID          `json:"DeliverToAddress"`
	DeliverToContactPerson               edm.GUID          `json:"DeliverToContactPerson"`
	DeliverToContactPersonFullName       edm.String        `json:"DeliverToContactPersonFullName"`
	DeliverToName                        edm.String        `json:"DeliverToName"`
	Description                          edm.String        `json:"Description"`
	Division                             edm.Int32         `json:"Division"`
	Document                             edm.GUID          `json:"Document"`
	DocumentNumber                       edm.Int32         `json:"DocumentNumber"`
	DocumentSubject                      edm.String        `json:"DocumentSubject"`
	DueDate                              edm.DateTime      `json:"DueDate"`
	ExtraDutyAmountFC                    edm.Double        `json:"ExtraDutyAmountFC"`
	InvoiceDate                          edm.DateTime      `json:"InvoiceDate"`
	InvoiceNumber                        edm.Int32         `json:"InvoiceNumber"`
	InvoiceTo                            edm.GUID          `json:"InvoiceTo"`
	InvoiceToContactPerson               edm.GUID          `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName       edm.String        `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                        edm.String        `json:"InvoiceToName"`
	IsExtraDuty                          edm.Boolean       `json:"IsExtraDuty"`
	Journal                              edm.String        `json:"Journal"`
	JournalDescription                   edm.String        `json:"JournalDescription"`
	Modified                             edm.DateTime      `json:"Modified"`
	Modifier                             edm.GUID          `json:"Modifier"`
	ModifierFullName                     edm.String        `json:"ModifierFullName"`
	OrderDate                            edm.DateTime      `json:"OrderDate"`
	OrderedBy                            edm.GUID          `json:"OrderedBy"`
	OrderedByContactPerson               edm.GUID          `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName       edm.String        `json:"OrderedByContactPersonFullName"`
	OrderedByName                        edm.String        `json:"OrderedByName"`
	OrderNumber                          edm.Int32         `json:"OrderNumber"`
	PaymentCondition                     edm.String        `json:"PaymentCondition"`
	PaymentConditionDescription          edm.String        `json:"PaymentConditionDescription"`
	PaymentReference                     edm.String        `json:"PaymentReference"`
	Remarks                              edm.String        `json:"Remarks"`
	SalesInvoiceLines                    SalesInvoiceLines `json:"SalesInvoiceLines"`
	Salesperson                          edm.GUID          `json:"Salesperson"`
	SalespersonFullName                  edm.String        `json:"SalespersonFullName"`
	StarterSalesInvoiceStatus            edm.Int16         `json:"StarterSalesInvoiceStatus"`
	StarterSalesInvoiceStatusDescription edm.String        `json:"StarterSalesInvoiceStatusDescription"`
	Status                               edm.Int16         `json:"Status"`
	StatusDescription                    edm.String        `json:"StatusDescription"`
	TaxSchedule                          edm.GUID          `json:"TaxSchedule"`
	TaxScheduleCode                      edm.String        `json:"TaxScheduleCode"`
	TaxScheduleDescription               edm.String        `json:"TaxScheduleDescription"`
	Type                                 edm.Int32         `json:"Type"`
	TypeDescription                      edm.String        `json:"TypeDescription"`
	VATAmountDC                          edm.Double        `json:"VATAmountDC"`
	VATAmountFC                          edm.Double        `json:"VATAmountFC"`
	WithholdingTaxAmountFC               edm.Double        `json:"WithholdingTaxAmountFC"`
	WithholdingTaxBaseAmount             edm.Double        `json:"WithholdingTaxBaseAmount"`
	WithholdingTaxPercentage             edm.Double        `json:"WithholdingTaxPercentage"`
	YourRef                              edm.String        `json:"YourRef"`
}

type SalesInvoiceLines []SalesInvoiceLine

// standalone: "SalesInvoiceLines": []
// deferred: "SalesInvoiceLines": {"__deferred": {}}
// embedded: "SalesInvoiceLines": {"results": []}
func (i *SalesInvoiceLines) UnmarshalJSON(data []byte) error {
	type Results SalesInvoiceLines

	type Envelope struct {
		Results  Results         `json:"results"`
		Deferred json.RawMessage `json:"__deferred"`
	}

	// create the json tester
	tester := &utils.JsonTester{}
	err := json.Unmarshal(data, tester)
	if err != nil {
		return err
	}

	// test if json is array (standalone)
	if tester.IsArray() {
		results := &Results{}
		err := json.Unmarshal(data, results)
		if err != nil {
			return err
		}

		*i = SalesInvoiceLines(*results)
		return nil
	}

	envelope := &Envelope{Results: Results(*i)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*i = SalesInvoiceLines(envelope.Results)
	return nil
}

type SalesInvoiceLine struct {
	ID                      edm.GUID     `json:"ID"`
	AmountDC                edm.Double   `json:"AmountDC"`
	AmountFC                edm.Double   `json:"AmountFC"`
	CostCenter              edm.String   `json:"CostCenter"`
	CostCenterDescription   edm.String   `json:"CostCenterDescription"`
	CostUnit                edm.String   `json:"CostUnit"`
	CostUnitDescription     edm.String   `json:"CostUnitDescription"`
	DeliveryDate            edm.DateTime `json:"DeliveryDate"`
	Description             edm.String   `json:"Description"`
	Discount                edm.Double   `json:"Discount"`
	Division                edm.Int32    `json:"Division"`
	Employee                edm.GUID     `json:"Employee"`
	EmployeeFullName        edm.String   `json:"EmployeeFullName"`
	EndTime                 edm.DateTime `json:"EndTime"`
	ExtraDutyAmountFC       edm.Double   `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage     edm.Double   `json:"ExtraDutyPercentage"`
	GLAccount               edm.GUID     `json:"GLAccount"`
	GLAccountDescription    edm.String   `json:"GLAccountDescription"`
	InvoiceID               edm.GUID     `json:"InvoiceID"`
	Item                    edm.GUID     `json:"Item"`
	ItemCode                edm.String   `json:"ItemCode"`
	ItemDescription         edm.String   `json:"ItemDescription"`
	LineNumber              edm.Int32    `json:"LineNumber"`
	NetPrice                edm.Double   `json:"NetPrice"`
	Notes                   edm.String   `json:"Notes"`
	Pricelist               edm.GUID     `json:"Pricelist"`
	PricelistDescription    edm.String   `json:"PricelistDescription"`
	Project                 edm.GUID     `json:"Project"`
	ProjectDescription      edm.String   `json:"ProjectDescription"`
	ProjectWBS              edm.GUID     `json:"ProjectWBS"`
	ProjectWBSDescription   edm.String   `json:"ProjectWBSDescription"`
	Quantity                edm.Double   `json:"Quantity"`
	SalesOrder              edm.GUID     `json:"SalesOrder"`
	SalesOrderLine          edm.GUID     `json:"SalesOrderLine"`
	SalesOrderLineNumber    edm.Int32    `json:"SalesOrderLineNumber"`
	SalesOrderNumber        edm.Int32    `json:"SalesOrderNumber"`
	StartTime               edm.DateTime `json:"StartTime"`
	Subscription            edm.GUID     `json:"Subscription"`
	SubscriptionDescription edm.String   `json:"SubscriptionDescription"`
	TaxSchedule             edm.GUID     `json:"TaxSchedule"`
	TaxScheduleCode         edm.String   `json:"TaxScheduleCode"`
	TaxScheduleDescription  edm.String   `json:"TaxScheduleDescription"`
	UnitCode                edm.String   `json:"UnitCode"`
	UnitDescription         edm.String   `json:"UnitDescription"`
	UnitPrice               edm.Double   `json:"UnitPrice"`
	VATAmountDC             edm.Double   `json:"VATAmountDC"`
	VATAmountFC             edm.Double   `json:"VATAmountFC"`
	VATCode                 edm.String   `json:"VATCode"`
	VATCodeDescription      edm.String   `json:"VATCodeDescription"`
	VATPercentage           edm.Double   `json:"VATPercentage"`
}
