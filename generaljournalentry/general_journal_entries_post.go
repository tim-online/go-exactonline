package generaljournalentry

import (
	"context"
	"net/http"
)

// POST

func (s *Service) GeneralJournalEntriesPost(body *GeneralJournalEntriesPostBody, ctx context.Context) (*GeneralJournalEntriesPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewGeneralJournalEntriesPostResponse()
	path := s.rest.SubPath(GeneralJournalEntriesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type GeneralJournalEntriesPostBody NewGeneralJournalEntry

func (s *Service) NewGeneralJournalEntriesPostBody() *GeneralJournalEntriesPostBody {
	return &GeneralJournalEntriesPostBody{
		GeneralJournalEntryLines: NewGeneralJournalEntryLines{},
	}
}

func (s *Service) NewGeneralJournalEntriesPostResponse() *GeneralJournalEntriesPostResponse {
	return &GeneralJournalEntriesPostResponse{}
}

type GeneralJournalEntriesPostResponse GeneralJournalEntry
