package salesentry

import "log"

func (s *Service) NewEntryNumber() (int, error) {
	// get highest JournalEntry.EntryNumber
	biggest, err := s.GetHighestEntryNumber()
	if err != nil {
		return 0, err
	}

	return biggest + 1, nil
}

func (s *Service) GetHighestEntryNumber() (int, error) {
	params := s.NewSalesEntriesGetParams()
	params.Top.Set(1)
	params.Select.Add("EntryNumber")
	params.OrderBy.Add("EntryNumber", "DESC")

	resp, err := s.SalesEntriesGet(params, nil)
	if err != nil {
		return 0, err
	}

	if len(resp.Results) == 0 {
		return 0, nil
	}

	log.Printf("%+v", resp.Results)
	return int(resp.Results[0].EntryNumber), nil
}
