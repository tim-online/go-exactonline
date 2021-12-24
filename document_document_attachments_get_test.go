package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDocumentAttachmentsGet(t *testing.T) {
	params := client.Document.NewDocumentAttachmentsGetParams()
	resp, err := client.Document.DocumentAttachmentsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
