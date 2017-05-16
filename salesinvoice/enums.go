package salesinvoice

import "github.com/tim-online/go-exactonline/edm"

const (
	InvoiceTypeInvoice    InvoiceType = 8020
	InvoiceTypeCreditNote             = 8021
)

type InvoiceType edm.Int32
