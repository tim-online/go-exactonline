package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProjectHourTypesGet(t *testing.T) {
	params := client.Project.NewHourTypesGetParams()
	resp, err := client.Project.HourTypesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
