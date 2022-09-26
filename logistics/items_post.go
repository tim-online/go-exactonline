package logistics

import (
	"context"
	"net/http"
)

// POST

func (s *Service) ItemsPost(body *ItemsPostBody, ctx context.Context) (*ItemsPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewItemsPostResponse()
	path := s.rest.SubPath(ItemsEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type ItemsPostBody NewItem

func (s *Service) NewItemsPostBody() *ItemsPostBody {
	return &ItemsPostBody{}
}

func (s *Service) NewItemsPostResponse() *ItemsPostResponse {
	return &ItemsPostResponse{}
}

type ItemsPostResponse Item
