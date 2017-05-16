package edm

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

// type GUID uuid.UUID
type GUID struct {
	uuid.UUID
}

func (g *GUID) MarshalJSON() ([]byte, error) {
	empty := uuid.UUID{}
	if g.UUID == empty {
		return json.Marshal(nil)
	}

	return json.Marshal(g.UUID)
}
