package edm

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

// type GUID uuid.UUID
type GUID struct {
	uuid.UUID
}

func (g *GUID) Empty() bool {
	return g.UUID == uuid.Nil
}

func (g *GUID) String() string {
	if g.Empty() {
		return ""
	}
	return g.UUID.String()
}

func (g *GUID) MarshalJSON() ([]byte, error) {
	if g.Empty() {
		return json.Marshal(nil)
	}

	return json.Marshal(g.UUID)
}

func (g *GUID) FromString(input string) {
	g.UUID, _ = uuid.FromString(input)
}
