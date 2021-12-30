package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDocumentTypeCategoriesGet(t *testing.T) {
	params := client.Document.NewDocumentTypeCategoriesGetParams()
	resp, err := client.Document.DocumentTypeCategoriesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
