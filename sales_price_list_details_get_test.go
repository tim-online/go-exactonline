package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSalesPriceListDetailsGet(t *testing.T) {
	params := client.Sales.NewPriceListDetailsGetParams()
	resp, err := client.Sales.PriceListDetailsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
