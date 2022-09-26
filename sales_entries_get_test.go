package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSalesEntriesGet(t *testing.T) {
	params := client.SalesEntry.NewSalesEntriesGetParams()
	params.Filter.Set("YourRef eq '21731'")
	resp, err := client.SalesEntry.SalesEntriesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
