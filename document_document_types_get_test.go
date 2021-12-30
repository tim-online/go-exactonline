package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDocumentTypesGet(t *testing.T) {
	params := client.Document.NewDocumentTypesGetParams()
	resp, err := client.Document.DocumentTypesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
