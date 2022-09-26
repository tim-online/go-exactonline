package project

import (
	"context"
	"net/http"
)

// POST

func (s *Service) TimeTransactionsPost(body *TimeTransactionsPostBody, ctx context.Context) (*TimeTransactionsPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewTimeTransactionsPostResponse()
	path := s.rest.SubPath(TimeTransactionsEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type TimeTransactionsPostBody NewTimeTransaction

func (s *Service) NewTimeTransactionsPostBody() *TimeTransactionsPostBody {
	return &TimeTransactionsPostBody{}
}

func (s *Service) NewTimeTransactionsPostResponse() *TimeTransactionsPostResponse {
	return &TimeTransactionsPostResponse{}
}

type TimeTransactionsPostResponse TimeTransaction
