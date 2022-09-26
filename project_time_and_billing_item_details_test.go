package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProjectTimeAndBillingItemDetailsGet(t *testing.T) {
	params := client.Project.NewTimeAndBillingItemDetailsGetParams()
	resp, err := client.Project.TimeAndBillingItemDetailsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
