package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLogisticsUnitsGet(t *testing.T) {
	params := client.Logistics.NewUnitsGetParams()
	resp, err := client.Logistics.UnitsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
