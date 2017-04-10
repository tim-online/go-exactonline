package exact

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-exactonline/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
	// apiPrefix      = "/api"
)

// Client manages communication with Exact Online API
type Client struct {
	// SOAP client used to communicate with the API.
	client *http.Client

	// Url pointing to base Exact Online API
	BaseURL *url.URL

	// Debugging flag
	Debug bool

	// User agent for client
	UserAgent string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

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
	// Financial            *Financial
	FinancialTransaction *FinancialTransaction
	// General              *General
	// GeneralJournalEntry  *GeneralJournalEntry
	// HRM                  *HRM
	// Inventory            *Inventory
	// Logistics            *Logistics
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
	System *SystemService
	// Users                *Users
	// VAT                  *VAT
	// Workflow             *Workflow
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// NewClient returns a new Exact Online API client
func NewClient(httpClient *http.Client, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:    httpClient,
		BaseURL:   nil,
		UserAgent: userAgent,
		Debug:     false,
	}

	c.SetBaseURL(baseURL)

	c.FinancialTransaction = NewFinancialTransactionService(c)
	c.System = NewSystemService(c)

	return c
}

func (c *Client) SetDebug(debug bool) {
	c.Debug = debug
}

func (c *Client) SetSandbox(sandbox bool) {
	if sandbox == true {
		// u, _ := url.ParseRequestURI(companies.SandboxEndpoint)
		// c.Companies.Endpoint = u
	} else {
		// u, _ := url.ParseRequestURI(companies.Endpoint)
		// c.Companies.Endpoint = u
	}
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	// set base url for use in http client
	c.BaseURL = baseURL
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	u := c.GetEndpoint(path)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", mediaType, charset))
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) GetEndpoint(path string) *url.URL {
	basePath := strings.TrimSuffix(c.BaseURL.Path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	u := *c.BaseURL
	u.Path = basePath + path
	return &u
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.Debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.Debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// {
	// 	"d" : {
	// 		"results" : [
	// 		{}
	// 		]
	// 	}
	// }

	type D struct {
		Results interface{} `json:"results"`
	}

	type Envelope struct {
		D D `json:"d"`
	}

	envelope := &Envelope{D: D{Results: responseBody}}

	// try to decode body into interface parameter
	err = json.NewDecoder(httpResp.Body).Decode(envelope)
	return httpResp, err
}
