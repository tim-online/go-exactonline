package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestFinancialGLAccountsGet(t *testing.T) {
	params := client.Financial.NewGLAccountsGetParams()
	resp, err := client.Financial.GLAccountsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
