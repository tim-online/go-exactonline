package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestFinancialJournalsGet(t *testing.T) {
	params := client.Financial.NewJournalsGetParams()
	resp, err := client.Financial.JournalsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
