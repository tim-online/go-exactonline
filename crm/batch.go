package crm

import (
	"context"
	"log"
	"net/http"
)

// Batch

func (s *Service) AccountsBatch(body *AccountsBatchBody, ctx context.Context) (*AccountsPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewAccountsPostResponse()
	path := s.rest.SubPath(AccountsEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type AccountsBatchBody struct {
	// Gets    []AccountsGetRequest
	// Puts  []AccountsPutRequest
	Puts []AccountsPutBody
	// Posts []AccountsPostRequest
	Posts []AccountsPostBody
	// Deletes []AccountsDeleteRequest
}

func (a *AccountsBatchBody) Read(p []byte) (n int, err error) {
	log.Fatal("@TODO")
	return 0, nil
}

func (s *Service) NewAccountsBatchBody() *AccountsBatchBody {
	return &AccountsBatchBody{}
}

func (s *Service) NewAccountsBatchResponse() *AccountsBatchResponse {
	return &AccountsBatchResponse{}
}

type AccountsBatchResponse Account
