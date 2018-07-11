package edm

import (
	"encoding/json"
	"strconv"
)

type Double float64

func (d *Double) UnmarshalJSON(text []byte) (err error) {
	var f float64
	err = json.Unmarshal(text, &f)
	if err == nil {
		*d = Double(f)
		return nil
	}

	var s string
	err = json.Unmarshal(text, &s)
	if err != nil {
		return err
	}

	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*d = Double(f)
	return nil
}
