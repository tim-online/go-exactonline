package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestVatCodesGet(t *testing.T) {
	params := client.VAT.NewVatCodesGetParams()
	resp, err := client.VAT.VatCodesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
