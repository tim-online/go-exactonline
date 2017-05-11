package odata

type Select []string

func (s *Select) Add(key string) {
	*s = append(*s, key)
}

// req.URL.RawQuery = "$select=EntryID,ClosingBalanceFC,Date,Description,Division,Document,DocumentNumber,DocumentSubject,EntryNumber,ExternalLinkDescription,ExternalLinkReference,FinancialPeriod,FinancialYear,IsExtraDuty,JournalCode,JournalDescription,Modified,OpeningBalanceFC,PaymentConditionCode,PaymentConditionDescription,PaymentReference,Status,StatusDescription,Type,TransactionLines"
