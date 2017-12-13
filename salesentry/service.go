package salesentry

import "github.com/tim-online/go-exactonline/rest"

const (
	SalesEntriesEndpoint = "/v1/{division}/salesentry/SalesEntries"
)

func NewService(rest *rest.Client) *Service {
	return &Service{rest: rest}
}

type Service struct {
	rest *rest.Client
}
