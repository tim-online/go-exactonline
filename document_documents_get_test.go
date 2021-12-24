package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDocumentsGet(t *testing.T) {
	params := client.Document.NewDocumentsGetParams()
	resp, err := client.Document.DocumentsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
