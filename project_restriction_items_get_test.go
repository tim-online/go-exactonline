package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProjectProjectRestrictionItemsGet(t *testing.T) {
	params := client.Project.NewProjectRestrictionItemsGetParams()
	params.Select.Values = []string{
		"ID",
		"Created",
		"Creator",
		"CreatorFullName",
		"Division",
		"Item",
		"ItemCode",
		"ItemDescription",
		"ItemIsTime",
		"Modified",
		"Modifier",
		"ModifierFullName",
		"Project",
		"ProjectCode",
		"ProjectDescription",
	}
	resp, err := client.Project.ProjectRestrictionItemsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
