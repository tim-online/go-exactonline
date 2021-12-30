package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDocumentCategoriesGet(t *testing.T) {
	params := client.Document.NewDocumentCategoriesGetParams()
	resp, err := client.Document.DocumentCategoriesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
