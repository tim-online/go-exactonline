package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProjectsGet(t *testing.T) {
	params := client.Project.NewProjectsGetParams()
	resp, err := client.Project.ProjectsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
