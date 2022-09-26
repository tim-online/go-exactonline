package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProjectTimeTransactionsGet(t *testing.T) {
	params := client.Project.NewTimeTransactionsGetParams()
	resp, err := client.Project.TimeTransactionsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
