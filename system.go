package exact

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/tim-online/go-exactonline/edm"
)

const (
	SystemMeEndpoint        = "/v1/current/Me"
	SystemDivisionsEndpoint = "/v1/%d/system/Divisions"
)

func NewSystemService(client *Client) *SystemService {
	return &SystemService{Client: client}
}

type SystemService struct {
	Client *Client
}

// Me endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemDivisions

func (s *SystemService) MeGet(requestParams *SystemMeGetParams, ctx context.Context) (*SystemMeGetResponse, error) {
	method := "GET"
	responseBody := s.NewSystemMeGetResponse()
	path := fmt.Sprintf(SystemMeEndpoint)

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

func NewSystemMeGetResponse() *SystemMeGetResponse {
	return &SystemMeGetResponse{}
}

type SystemMeGetResponse []SystemMeUser

type SystemMeUser struct {
	MetaData               MetaData   `json:"__metadata"`
	UserID                 edm.GUID   `json:"UserID"`
	CurrentDivision        edm.Int32  `json:"CurrentDivision"`
	DivisionCustomer       edm.GUID   `json:"DivisionCustomer"`
	DivisionCustomerCode   edm.String `json:"DivisionCustomerCode"`
	Email                  edm.String `json:"Email"`
	EmployeeID             edm.GUID   `json:"EmployeeID"`
	FirstName              edm.String `json:"FirstName"`
	FullName               edm.String `json:"FullName"`
	Gender                 Gender     `json:"Gender"`
	Initials               edm.String `json:"Initials"`
	Language               edm.String `json:"Language"`
	LanguageCode           edm.String `json:"LanguageCode"`
	LastName               edm.String `json:"LastName"`
	Legislation            edm.Int64  `json:"Legislation,string"`
	MiddleName             edm.String `json:"MiddleName"`
	Mobile                 edm.String `json:"Mobile"`
	Nationality            edm.String `json:"Nationality"`
	Phone                  edm.String `json:"Phone"`
	PhoneExtension         edm.String `json:"PhoneExtension"`
	PictureURL             URL        `json:"PictureUrl"`
	ServerTime             edm.String `json:"ServerTime"`
	ServerUtcOffset        edm.Double `json:"ServerUtcOffset"`
	ThumbnailPicture       edm.Binary `json:"ThumbnailPicture"`
	ThumbnailPictureFormat edm.String `json:"ThumbnailPictureFormat"`
	Title                  edm.String `json:"Title"`
	UserName               edm.String `json:"UserName"`
}

func (s *SystemService) NewMeGetParams() *SystemMeGetParams {
	return &SystemMeGetParams{}
}

type SystemMeGetParams struct {
	// @TODO: check if this an OData struct or something
	Select string `schema:"$select"`
	Filter string `schema:"$filter"`
}

func (p *SystemMeGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}

// Divisions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemDivisions

func (s *SystemService) DivisionsGet(divisionID int, requestParams *SystemDivisionsGetParams, ctx context.Context) (*SystemDivisionsGetResponse, error) {
	method := "GET"
	responseBody := s.NewDivisionsGetResponse()
	path := fmt.Sprintf(SystemDivisionsEndpoint, divisionID)

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

func (s *SystemService) NewDivisionsGetResponse() *SystemDivisionsGetResponse {
	return &SystemDivisionsGetResponse{}
}

type SystemDivisionsGetResponse []SystemDivisionsUser

type SystemDivisionsUser struct {
	MetaData                MetaData       `json:"__metadata"`
	Code                    edm.Int32      `json:"Code"`
	AddressLine1            edm.String     `json:"AddressLine1"`
	AddressLine2            edm.String     `json:"AddressLine2"`
	AddressLine3            edm.String     `json:"AddressLine3"`
	ChamberOfCommerceNumber edm.String     `json:"ChamberOfCommerceNumber"`
	City                    edm.String     `json:"City"`
	Country                 edm.String     `json:"Country"`
	Created                 edm.DateTime   `json:"Created"`
	Currency                edm.String     `json:"Currency"`
	Current                 edm.Boolean    `json:"Current"`
	Customer                edm.GUID       `json:"Customer"`
	CustomerCode            edm.String     `json:"CustomerCode"`
	CustomerName            edm.String     `json:"CustomerName"`
	Description             edm.String     `json:"Description"`
	Email                   edm.String     `json:"Email"`
	Hid                     edm.Int64      `json:"Hid,string"`
	IsMainDivision          edm.Boolean    `json:"IsMainDivision"`
	Modified                edm.DateTime   `json:"Modified"`
	Phone                   edm.String     `json:"Phone"`
	Postcode                edm.String     `json:"Postcode"`
	State                   edm.String     `json:"State"`
	Status                  DivisionStatus `json:"Status"`
}

func (s *SystemService) NewDivisionsGetParams() *SystemDivisionsGetParams {
	return &SystemDivisionsGetParams{}
}

type SystemDivisionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select string `schema:"$select,omitempty"`
	Filter string `schema:"$filter,omitempty"`
}

func (p *SystemDivisionsGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}
