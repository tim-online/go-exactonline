package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLogisticsSalesItemPricesGet(t *testing.T) {
	params := client.Logistics.NewSalesItemPricesGetParams()
	params.Select.Values = []string{
		"ID",
		"Account",
		"AccountName",
		"Created",
		"Creator",
		"CreatorFullName",
		"Currency",
		"DefaultItemUnit",
		"DefaultItemUnitDescription",
		"Division",
		"EndDate",
		"Item",
		"ItemCode",
		"ItemDescription",
		"Modified",
		"Modifier",
		"ModifierFullName",
		"NumberOfItemsPerUnit",
		"Price",
		"Quantity",
		"StartDate",
		"Unit",
		"UnitDescription",
	}
	resp, err := client.Logistics.SalesItemPricesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
