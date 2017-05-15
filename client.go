package exact

import (
	"net/http"
	"net/url"

	"github.com/tim-online/go-exactonline/financial"
	"github.com/tim-online/go-exactonline/financialtransaction"
	"github.com/tim-online/go-exactonline/logistics"
	"github.com/tim-online/go-exactonline/rest"
	"github.com/tim-online/go-exactonline/system"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-exactonline/" + libraryVersion
	// apiPrefix      = "/api"
)

// Client manages communication with Exact Online API
type Client struct {
	// REST client used to communicate with the API.
	rest.Client

	// Services
	// Accountancy          *Accountancy
	// Activities           *Activities
	// Assets               *Assets
	// Budget               *Budget
	// Bulk                 *Bulk
	// Cashflow             *Cashflow
	// ContinuousMonitoring *ContinuousMonitoring
	// CRM                  *CRM
	// Documents            *Documents
	Financial            *financial.Service
	FinancialTransaction *financialtransaction.Service
	// General              *General
	// GeneralJournalEntry  *GeneralJournalEntry
	// HRM                  *HRM
	// Inventory            *Inventory
	Logistics *logistics.Service
	// Mailbox              *Mailbox
	// Manufacturing        *Manufacturing
	// OpeningBalance       *OpeningBalance
	// Payroll              *Payroll
	// Project              *Project
	// PurchaseEntry        *PurchaseEntry
	// PurchaseOrder        *PurchaseOrder
	// Sales                *Sales
	// SalesEntry           *SalesEntry
	// SalesInvoice         *SalesInvoice
	// SalesOrder           *SalesOrder
	// Subscription         *Subscription
	System *system.Service
	// Users                *Users
	// VAT                  *VAT
	// Workflow             *Workflow
}

// NewClient returns a new Exact Online API client
func NewClient(httpClient *http.Client, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		Client: *rest.New(httpClient),
	}

	// set default options
	c.SetBaseURL(baseURL)
	c.SetUserAgent(userAgent)
	c.SetDebug(false)

	c.Financial = financial.NewService(&c.Client)
	c.FinancialTransaction = financialtransaction.NewService(&c.Client)
	c.Logistics = logistics.NewService(&c.Client)
	c.System = system.NewService(&c.Client)

	return c
}

func (c *Client) SetDebug(debug bool) {
	c.Client.SetDebug(debug)
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	// set base url for use in http client
	c.Client.SetBaseURL(baseURL)
}
