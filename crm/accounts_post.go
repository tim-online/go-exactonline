package crm

import (
	"context"
	"net/http"
)

// POST

func (s *Service) AccountsPost(body *AccountsPostBody, ctx context.Context) (*AccountsPostResponse, error) {
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

type AccountsPostBody NewAccount

func (s *Service) NewAccountsPostBody() *AccountsPostBody {
	return &AccountsPostBody{}
}

func (s *Service) NewAccountsPostResponse() *AccountsPostResponse {
	return &AccountsPostResponse{}
}

type AccountsPostResponse Account
