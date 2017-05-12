package financial

import "github.com/tim-online/go-exactonline/edm"

type Journals []Journal

type Journal struct {
	ID                              edm.GUID     `json:"ID"`
	AllowVariableCurrency           edm.Boolean  `json:"AllowVariableCurrency"`
	AllowVariableExchangeRate       edm.Boolean  `json:"AllowVariableExchangeRate"`
	AllowVAT                        edm.Boolean  `json:"AllowVAT"`
	AutoSave                        edm.Boolean  `json:"AutoSave"`
	Bank                            edm.GUID     `json:"Bank"`
	BankAccountBICCode              edm.String   `json:"BankAccountBICCode"`
	BankAccountCountry              edm.String   `json:"BankAccountCountry"`
	BankAccountDescription          edm.String   `json:"BankAccountDescription"`
	BankAccountIBAN                 edm.String   `json:"BankAccountIBAN"`
	BankAccountID                   edm.GUID     `json:"BankAccountID"`
	BankAccountIncludingMask        edm.String   `json:"BankAccountIncludingMask"`
	BankAccountUseSEPA              edm.Boolean  `json:"BankAccountUseSEPA"`
	BankAccountUseSepaDirectDebit   edm.Boolean  `json:"BankAccountUseSepaDirectDebit"`
	BankName                        edm.String   `json:"BankName"`
	Code                            edm.String   `json:"Code"`
	Created                         edm.DateTime `json:"Created"`
	Creator                         edm.GUID     `json:"Creator"`
	CreatorFullName                 edm.String   `json:"CreatorFullName"`
	Currency                        edm.String   `json:"Currency"`
	CurrencyDescription             edm.String   `json:"CurrencyDescription"`
	Description                     edm.String   `json:"Description"`
	Division                        edm.Int32    `json:"Division"`
	GLAccount                       edm.GUID     `json:"GLAccount"`
	GLAccountCode                   edm.String   `json:"GLAccountCode"`
	GLAccountDescription            edm.String   `json:"GLAccountDescription"`
	GLAccountType                   edm.Int32    `json:"GLAccountType"`
	Modified                        edm.DateTime `json:"Modified"`
	Modifier                        edm.GUID     `json:"Modifier"`
	ModifierFullName                edm.String   `json:"ModifierFullName"`
	PaymentInTransitAccount         edm.GUID     `json:"PaymentInTransitAccount"`
	PaymentServiceAccountIdentifier edm.String   `json:"PaymentServiceAccountIdentifier"`
	PaymentServiceProvider          edm.Int32    `json:"PaymentServiceProvider"`
	PaymentServiceProviderName      edm.String   `json:"PaymentServiceProviderName"`
	Type                            JournalType  `json:"Type"`
}

const (
	TypeCash                  JournalType = 10
	TypeBank                  JournalType = 12
	TypePaymentService        JournalType = 16
	TypeSales                 JournalType = 20
	TypeReturnInvoce          JournalType = 21
	TypePurchase              JournalType = 22
	TypeReceivedReturnInvoice JournalType = 23
	TypeGeneralJournal        JournalType = 90
)

type JournalType edm.Int32

func (j *JournalType) String() string {
	switch int(*j) {
	case 10:
		return "Cash"
	case 12:
		return "Bank"
	case 16:
		return "Payment service"
	case 20:
		return "Sales"
	case 21:
		return "Return invoice"
	case 22:
		return "Purchase"
	case 23:
		return "Received return invoice"
	case 90:
		return "General journal"
	}
	return ""
}

type GLSchemes []GLScheme

type GLScheme struct {
    ID                 edm.GUID       `json:"ID"`
    Code               edm.String     `json:"Code"`
    Created            edm.DateTime   `json:"Created"`
    Creator            edm.GUID       `json:"Creator"`
    CreatorFullName    edm.String     `json:"CreatorFullName"`
    Description        edm.String     `json:"Description"`
    Division           edm.Int32      `json:"Division"`
    Main               edm.Byte       `json:"Main"`
    Modified           edm.DateTime   `json:"Modified"`
    Modifier           edm.GUID       `json:"Modifier"`
    ModifierFullName   edm.String     `json:"ModifierFullName"`
    TargetNamespace    edm.String     `json:"TargetNamespace"`
}
