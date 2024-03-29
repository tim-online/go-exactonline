package edm

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// type GUID uuid.UUID
type GUID struct {
	uuid.UUID
}

func (g GUID) IsEmpty() bool {
	return g.UUID == uuid.Nil
}

func (g GUID) String() string {
	if g.IsEmpty() {
		return ""
	}
	return g.UUID.String()
}

func (g *GUID) UnmarshalJSON(text []byte) (err error) {
	var s string
	err = json.Unmarshal(text, &s)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	uuid, err := uuid.FromString(s)
	*g = GUID{uuid}
	return err
}

func (g GUID) MarshalJSON() ([]byte, error) {
	if g.IsEmpty() {
		return json.Marshal(nil)
	}

	return json.Marshal(g.UUID)
}

func (g *GUID) FromString(input string) {
	g.UUID, _ = uuid.FromString(input)
}
