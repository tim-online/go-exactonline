package crm

import (
	"context"
	"net/http"

	"github.com/tim-online/go-exactonline/edm"
)

// POST

func (s *Service) AccountsPut(body *AccountsPutBody, ctx context.Context) error {
	method := http.MethodPut
	path := s.rest.SubPathWithID(AccountsEndpoint, body.ID.String())

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, nil)
	return err
}

type AccountsPutBody struct {
	ID                         edm.GUID     `json:"ID"`                         // Primary key
	Accountant                 edm.GUID     `json:"Accountant"`                 // Reference to the accountant of the customer. Conditions: The referred accountant must have value > 0 in the field IsAccountant
	AccountManager             edm.GUID     `json:"AccountManager"`             // ID of the account manager
	ActivitySector             edm.GUID     `json:"ActivitySector"`             // Reference to Activity sector of the account
	ActivitySubSector          edm.GUID     `json:"ActivitySubSector"`          // Reference to Activity sub-sector of the account
	AddressLine1               edm.String   `json:"AddressLine1"`               // Visit address first line
	AddressLine2               edm.String   `json:"AddressLine2"`               // Visit address second line
	AddressLine3               edm.String   `json:"AddressLine3"`               // Visit address third line
	Blocked                    edm.Boolean  `json:"Blocked"`                    // Indicates if the account is blocked
	BusinessType               edm.GUID     `json:"BusinessType"`               // Reference to the business type of the account
	CanDropShip                edm.Boolean  `json:"CanDropShip"`                // Indicates the default for the possibility to drop ship when an item is linked to a supplier
	ChamberOfCommerce          edm.String   `json:"ChamberOfCommerce"`          // Chamber of commerce number
	City                       edm.String   `json:"City"`                       // Visit address City
	Classification1            edm.GUID     `json:"Classification1"`            // Account classification 1
	Classification2            edm.GUID     `json:"Classification2"`            // Account classification 2
	Classification3            edm.GUID     `json:"Classification3"`            // Account classification 3
	Classification4            edm.GUID     `json:"Classification4"`            // Account classification 4
	Classification5            edm.GUID     `json:"Classification5"`            // Account classification 5
	Classification6            edm.GUID     `json:"Classification6"`            // Account classification 6
	Classification7            edm.GUID     `json:"Classification7"`            // Account classification 7
	Classification8            edm.GUID     `json:"Classification8"`            // Account classification 8
	Code                       edm.String   `json:"Code"`                       // Unique key, fixed length numeric string with leading spaces, length 18. IMPORTANT: When you use OData $filter on this field you have to make sure the filter parameter contains the leading spaces
	CodeAtSupplier             edm.String   `json:"CodeAtSupplier"`             // Code under which your own company is known at the account
	CompanySize                edm.GUID     `json:"CompanySize"`                // Reference to Company size of the account
	ConsolidationScenario      edm.Byte     `json:"ConsolidationScenario"`      // Consolidation scenario (Time & Billing). Values: 0 = No consolidation, 1 = Item, 2 = Item + Project, 3 = Item + Employee, 4 = Item + Employee + Project, 5 = Project + WBS + Item, 6 = Project + WBS + Item + Employee. Item means in this case including Unit and Price, these also have to be the same to consolidate
	ControlledDate             edm.DateTime `json:"ControlledDate"`             // Date of the latest control of account data with external web service
	Country                    edm.String   `json:"Country"`                    // Country code
	CreditLinePurchase         edm.Double   `json:"CreditLinePurchase"`         // Maximum amount of credit for Purchase. If no value has been defined, there is no credit limit
	CreditLineSales            edm.Double   `json:"CreditLineSales"`            // Maximum amount of credit for sales. If no value has been defined, there is no credit limit
	DatevCreditorCode          edm.String   `json:"DatevCreditorCode"`          // DATEV creditor code for Germany legislation
	DatevDebtorCode            edm.String   `json:"DatevDebtorCode"`            // DATEV debtor code for Germany legislation
	DiscountPurchase           edm.Double   `json:"DiscountPurchase"`           // Default discount percentage for purchase. This is stored as a fraction. ie 5.5% is stored as .055
	DiscountSales              edm.Double   `json:"DiscountSales"`              // Default discount percentage for sales. This is stored as a fraction. ie 5.5% is stored as .055
	Email                      edm.String   `json:"Email"`                      // E-Mail address of the account
	EndDate                    edm.DateTime `json:"EndDate"`                    // Determines in combination with the start date if the account is active. If the current date is > end date the account is inactive
	EstablishedDate            edm.DateTime `json:"EstablishedDate"`            // RegistrationDate
	Fax                        edm.String   `json:"Fax"`                        // Fax number
	GLAccountPurchase          edm.GUID     `json:"GLAccountPurchase"`          // Default (corporate) GL offset account for purchase (cost)
	GLAccountSales             edm.GUID     `json:"GLAccountSales"`             // Default (corporate) GL offset account for sales (revenue)
	GLAP                       edm.GUID     `json:"GLAP"`                       // Default GL account for Accounts Payable
	GLAR                       edm.GUID     `json:"GLAR"`                       // Default GL account for Accounts Receivable
	HasWithholdingTaxSales     edm.Boolean  `json:"HasWithholdingTaxSales"`     // Indicates whether a customer has withholding tax on sales
	IgnoreDatevWarningMessage  edm.Boolean  `json:"IgnoreDatevWarningMessage"`  // Suppressed warning message when there is duplication on the DATEV code
	IntraStatArea              edm.String   `json:"IntraStatArea"`              // Intrastat Area
	IntraStatDeliveryTerm      edm.String   `json:"IntraStatDeliveryTerm"`      // Intrastat delivery method
	IntraStatSystem            edm.String   `json:"IntraStatSystem"`            // System for Intrastat
	IntraStatTransactionA      edm.String   `json:"IntraStatTransactionA"`      // Transaction type A for Intrastat
	IntraStatTransactionB      edm.String   `json:"IntraStatTransactionB"`      // Transaction type B for Intrastat
	IntraStatTransportMethod   edm.String   `json:"IntraStatTransportMethod"`   // Transport method for Intrastat
	InvoiceAccount             edm.GUID     `json:"InvoiceAccount"`             // ID of account to be invoiced instead of this account
	InvoiceAttachmentType      edm.Int32    `json:"InvoiceAttachmentType"`      // Indicates which attachment types should be sent when a sales invoice is printed. Only values in related table with Invoice=1 are allowed
	InvoicingMethod            edm.Int32    `json:"InvoicingMethod"`            // Method of sending for sales invoices. Values: 1: Paper, 2: EMail, 4: Mailbox (electronic exchange)
	IsAccountant               edm.Byte     `json:"IsAccountant"`               // Indicates whether the account is an accountant. Values: 0 = No accountant, 1 = True, but accountant doesn't want his name to be published in the list of accountants, 2 = True, and accountant is published in the list of accountants
	IsAgency                   edm.Byte     `json:"IsAgency"`                   // Indicates whether the accounti is an agency
	IsCompetitor               edm.Byte     `json:"IsCompetitor"`               // Indicates whether the account is a competitor
	IsExtraDuty                edm.Boolean  `json:"IsExtraDuty"`                // Indicates whether a customer is eligible for extra duty
	IsMailing                  edm.Byte     `json:"IsMailing"`                  // Indicates if the account is excluded from mailing marketing information
	IsPilot                    edm.Boolean  `json:"IsPilot"`                    // Indicates whether the account is a pilot account
	IsReseller                 edm.Boolean  `json:"IsReseller"`                 // Indicates whether the account is a reseller
	IsSupplier                 edm.Boolean  `json:"IsSupplier"`                 // Indicates whether the account is a supplier
	Language                   edm.String   `json:"Language"`                   // Language code
	Latitude                   edm.Double   `json:"Latitude"`                   // Latitude (used by Google maps)
	LeadSource                 edm.GUID     `json:"LeadSource"`                 // Reference to Lead source of an account
	Logo                       edm.Binary   `json:"Logo"`                       // Bytes of the logo image
	LogoFileName               edm.String   `json:"LogoFileName"`               // The file name (without path, but with extension) of the image
	Longitude                  edm.Double   `json:"Longitude"`                  // Longitude (used by Google maps)
	MainContact                edm.GUID     `json:"MainContact"`                // Reference to main contact person
	Name                       edm.String   `json:"Name"`                       // Account name
	Parent                     edm.GUID     `json:"Parent"`                     // ID of the parent account
	PayAsYouEarn               edm.String   `json:"PayAsYouEarn"`               // Indicates the loan repayment plan for UK legislation
	PaymentConditionPurchase   edm.String   `json:"PaymentConditionPurchase"`   // Code of default payment condition for purchase
	PaymentConditionSales      edm.String   `json:"PaymentConditionSales"`      // Code of default payment condition for sales
	Phone                      edm.String   `json:"Phone"`                      // Phone number
	PhoneExtension             edm.String   `json:"PhoneExtension"`             // Phone number extention
	Postcode                   edm.String   `json:"Postcode"`                   // Visit address postcode
	PriceList                  edm.GUID     `json:"PriceList"`                  // Default sales price list for account
	PurchaseCurrency           edm.String   `json:"PurchaseCurrency"`           // Currency of purchase
	PurchaseLeadDays           edm.Int32    `json:"PurchaseLeadDays"`           // Indicates number of days required to receive a purchase. Acts as a default
	PurchaseVATCode            edm.String   `json:"PurchaseVATCode"`            // Default VAT code used for purchase entries
	RecepientOfCommissions     edm.Boolean  `json:"RecepientOfCommissions"`     // Define the relation that should be taken in the official document of the rewarding fiscal fiches Belcotax
	Remarks                    edm.String   `json:"Remarks"`                    // Remarks
	Reseller                   edm.GUID     `json:"Reseller"`                   // ID of the reseller account. Conditions: the target account must have the property IsReseller turned on
	RSIN                       edm.String   `json:"RSIN"`                       // Fiscal number for NL legislation
	SalesCurrency              edm.String   `json:"SalesCurrency"`              // Currency of Sales used for Time & Billing
	SalesTaxSchedule           edm.GUID     `json:"SalesTaxSchedule"`           // Obsolete
	SalesVATCode               edm.String   `json:"SalesVATCode"`               // Default VAT code for a sales entry
	SearchCode                 edm.String   `json:"SearchCode"`                 // Search code
	SecurityLevel              edm.Int32    `json:"SecurityLevel"`              // Security level (0 - 100)
	SeparateInvPerProject      edm.Byte     `json:"SeparateInvPerProject"`      // Separate invoice per project (Time & Billing)
	SeparateInvPerSubscription edm.Byte     `json:"SeparateInvPerSubscription"` // Indicates how invoices are generated from subscriptions. 0 = subscriptions belonging to the same customer are combined in a single invoice. 1 = each subscription results in one invoice. In both cases, each individual subscription line results in one invoice line
	ShippingLeadDays           edm.Int32    `json:"ShippingLeadDays"`           // Indicates the number of days it takes to send goods to the customer. Acts as a default
	ShippingMethod             edm.GUID     `json:"ShippingMethod"`             // Default shipping method
	StartDate                  edm.DateTime `json:"StartDate"`                  // Indicates in combination with the end date if the account is active
	State                      edm.String   `json:"State"`                      // State/Province/County code When changing the Country and the State is filled, the State must be assigned with a valid value from the selected country or set to empty
	Status                     Status       `json:"Status"`                     // If the status field is filled this means the account is a customer. The value indicates the customer status. Possible values: A=None, S=Suspect, P=Prospect, C=Customer
	TradeName                  edm.String   `json:"TradeName"`                  // Trade name can be registered and shown with the client (for all legislations)
	UniqueTaxpayerReference    edm.String   `json:"UniqueTaxpayerReference"`    // Unique taxpayer reference for UK legislation
	VATLiability               edm.String   `json:"VATLiability"`               // Indicates the VAT status of an account to be able to identify the relation that should be selected in the VAT debtor listing in Belgium
	VATNumber                  edm.String   `json:"VATNumber"`                  // The number under which the account is known at the Value Added Tax collection agency
	Website                    edm.String   `json:"Website"`                    // Website of the account}
}

func (s *Service) NewAccountsPutBody() *AccountsPutBody {
	return &AccountsPutBody{}
}
