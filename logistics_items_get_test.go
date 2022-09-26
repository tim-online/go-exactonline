package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLogisticsItemsGet(t *testing.T) {
	params := client.Logistics.NewItemsGetParams()
	resp, err := client.Logistics.ItemsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
