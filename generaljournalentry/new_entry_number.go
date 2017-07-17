package generaljournalentry

func (s *Service) NewEntryNumber() (int, error) {
	// get highest JournalEntry.EntryNumber
	biggest, err := s.GetHighestEntryNumber()
	if err != nil {
		return 0, err
	}

	return biggest + 1, nil
}

func (s *Service) GetHighestEntryNumber() (int, error) {
	params := s.NewGeneralJournalEntriesGetParams()
	params.Top.Set(1)
	params.Select.Add("EntryNumber")
	params.OrderBy.Add("EntryNumber", "DESC")

	resp, err := s.GeneralJournalEntriesGet(params, nil)
	if err != nil {
		return 0, err
	}

	if len(resp.Results) == 0 {
		return 1, nil
	}

	return int(resp.Results[0].EntryNumber), nil
}
