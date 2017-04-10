package edm

import "github.com/tim-online/go-exactonline/utils"

type MetaData struct {
	URL  utils.URL `json:"uri"`
	Type string    `json:"type"`
}
