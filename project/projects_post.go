package project

import (
	"context"
	"net/http"
)

// POST

func (s *Service) ProjectsPost(body *ProjectsPostBody, ctx context.Context) (*ProjectsPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewProjectsPostResponse()
	path := s.rest.SubPath(ProjectsEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type ProjectsPostBody NewProject

func (s *Service) NewProjectsPostBody() *ProjectsPostBody {
	return &ProjectsPostBody{}
}

func (s *Service) NewProjectsPostResponse() *ProjectsPostResponse {
	return &ProjectsPostResponse{}
}

type ProjectsPostResponse Project
