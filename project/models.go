package project

import "github.com/tim-online/go-exactonline/edm"

type Projects []Project

type Project struct {
	ID                               edm.GUID     `json:"ID,omitempty"`                               // Primary key
	Account                          edm.GUID     `json:"Account,omitempty"`                          // The account for this project
	AccountCode                      edm.String   `json:"AccountCode,omitempty"`                      // Code of Account
	AccountContact                   edm.GUID     `json:"AccountContact,omitempty"`                   // Contact person of Account
	AccountName                      edm.String   `json:"AccountName,omitempty"`                      // Name of Account
	AllowAdditionalInvoicing         edm.Boolean  `json:"AllowAdditionalInvoicing,omitempty"`         // Indicates if additional invoice is allowed for project
	BlockEntry                       edm.Boolean  `json:"BlockEntry,omitempty"`                       // Block time and cost entries
	BlockPurchasing                  edm.Boolean  `json:"BlockPurchasing,omitempty"`                  // Block purchasing
	BlockRebilling                   edm.Boolean  `json:"BlockRebilling,omitempty"`                   // Block rebilling
	BudgetedAmount                   edm.Double   `json:"BudgetedAmount,omitempty"`                   // Budgeted amount of sales in the default currency of the company
	BudgetedCosts                    edm.Double   `json:"BudgetedCosts,omitempty"`                    // Budgeted amount of costs in the default currency of the company
	BudgetedHoursPerHourType         interface{}  `json:"BudgetedHoursPerHourType,omitempty"`         // Collection of budgeted hours
	BudgetedRevenue                  edm.Double   `json:"BudgetedRevenue,omitempty"`                  // Budgeted amount of revenue in the default currency of the company
	BudgetOverrunHours               edm.Byte     `json:"BudgetOverrunHours,omitempty"`               // BudgetOverrunHours: 10-Allowed, 20-Not Allowed
	BudgetType                       edm.Int16    `json:"BudgetType,omitempty"`                       // Budget type
	BudgetTypeDescription            edm.String   `json:"BudgetTypeDescription,omitempty"`            // Budget type description
	Classification                   edm.GUID     `json:"Classification,omitempty"`                   // Used only for PSA to link a project classification to the project
	ClassificationDescription        edm.String   `json:"ClassificationDescription,omitempty"`        // Description of Classification
	Code                             edm.String   `json:"Code"`                                       // Code
	CostsAmountFC                    edm.Double   `json:"CostsAmountFC,omitempty"`                    // Used only for PSA to store the budgetted costs of a project (except for project type Campaign and Non-billable). Positive quantities only
	Created                          edm.DateTime `json:"Created,omitempty"`                          // Creation date
	Creator                          edm.GUID     `json:"Creator,omitempty"`                          // User ID of creator
	CreatorFullName                  edm.String   `json:"CreatorFullName,omitempty"`                  // Name of creator
	CustomerPOnumber                 edm.String   `json:"CustomerPOnumber,omitempty"`                 // Used only for PSA to store the customer's PO number
	CustomField                      edm.String   `json:"CustomField,omitempty"`                      // Custom field endpoint
	Description                      edm.String   `json:"Description"`                                // Description of the project
	Division                         edm.Int32    `json:"Division,omitempty"`                         // Division code
	DivisionName                     edm.String   `json:"DivisionName,omitempty"`                     // Name of Division
	EndDate                          edm.DateTime `json:"EndDate,omitempty"`                          // End date of the project. In combination with the start date the status is determined
	FixedPriceItem                   edm.GUID     `json:"FixedPriceItem,omitempty"`                   // Item used for fixed price invoicing. To be defined per project. If empty the functionality relies on the setting
	FixedPriceItemDescription        edm.String   `json:"FixedPriceItemDescription,omitempty"`        // Description of FixedPriceItem
	HasWBSLines                      edm.Boolean  `json:"HasWBSLines,omitempty"`                      // Indicates if whether the Project has WBS
	IncludeInvoiceSpecification      edm.Int16    `json:"IncludeInvoiceSpecification,omitempty"`      // Include invoice specification. E.g: 1 = Based on account, 2 = Always, 3 = Never
	IncludeSpecificationInInvoicePdf edm.Boolean  `json:"IncludeSpecificationInInvoicePdf,omitempty"` // Indicates whether to include invoice specification in invoice PDF
	InternalNotes                    edm.String   `json:"InternalNotes,omitempty"`                    // Internal notes not to be printed in invoice
	InvoiceAddress                   edm.GUID     `json:"InvoiceAddress,omitempty"`                   // Invoice address
	InvoiceAsQuoted                  edm.Boolean  `json:"InvoiceAsQuoted,omitempty"`                  // Indicates whether the project is invoice as quoted
	InvoiceTerms                     interface{}  `json:"InvoiceTerms,omitempty"`                     // Collection of invoice terms
	Manager                          edm.GUID     `json:"Manager,omitempty"`                          // Responsible person for this project
	ManagerFullname                  edm.String   `json:"ManagerFullname,omitempty"`                  // Name of Manager
	MarkupPercentage                 edm.Double   `json:"MarkupPercentage,omitempty"`                 // Purchase markup percentage
	Modified                         edm.DateTime `json:"Modified,omitempty"`                         // Last modified date
	Modifier                         edm.GUID     `json:"Modifier,omitempty"`                         // User ID of modifier
	ModifierFullName                 edm.String   `json:"ModifierFullName,omitempty"`                 // Name of modifier
	Notes                            edm.String   `json:"Notes,omitempty"`                            // For additional information about projects
	PrepaidItem                      edm.GUID     `json:"PrepaidItem,omitempty"`                      // Used only for PSA. This item is used for prepaid invoicing. If left empty, the functionality relies on a setting
	PrepaidItemDescription           edm.String   `json:"PrepaidItemDescription,omitempty"`           // Description of PrepaidItem
	PrepaidType                      edm.Int16    `json:"PrepaidType,omitempty"`                      // PrepaidType: 1-Retainer, 2-Hour type bundle
	PrepaidTypeDescription           edm.String   `json:"PrepaidTypeDescription,omitempty"`           // Description of PrepaidType
	ProjectRestrictionEmployees      interface{}  `json:"ProjectRestrictionEmployees,omitempty"`      // Collection of employee restrictions
	ProjectRestrictionItems          interface{}  `json:"ProjectRestrictionItems,omitempty"`          // Collection of item restrictions
	ProjectRestrictionRebillings     interface{}  `json:"ProjectRestrictionRebillings,omitempty"`     // Collection of rebilling restrictions
	SalesTimeQuantity                edm.Double   `json:"SalesTimeQuantity,omitempty"`                // Budgeted time. Total number of hours estimated for the fixed price project
	SourceQuotation                  edm.GUID     `json:"SourceQuotation,omitempty"`                  // Source quotation
	StartDate                        edm.DateTime `json:"StartDate,omitempty"`                        // Start date of a project. In combination with the end date the status is determined
	TimeQuantityToAlert              edm.Double   `json:"TimeQuantityToAlert,omitempty"`              // Alert when exceeding (Hours)
	Type                             edm.Int32    `json:"Type"`                                       // Reference to ProjectTypes. E.g: 1 = Campaign , 2 = Fixed Price, 3 = Time and Material, 4 = Non billable, 5 = Prepaid
	TypeDescription                  edm.String   `json:"TypeDescription,omitempty"`                  // Description of Type
	UseBillingMilestones             edm.Boolean  `json:"UseBillingMilestones,omitempty"`             // Indicates whether the Project is using billing milestones
}

type NewProject Project
