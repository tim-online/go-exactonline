package exact

import (
	nativeurl "net/url"
)

type MetaData struct {
	URL  URL    `json:"uri"`
	Type string `json:"type"`
}

// 0 for Inactive, 1 for Active and 2 for Archived Divisions
type DivisionStatus int32

// M=Male, V=Female, O=Unknown
type Gender string

type URL struct {
	nativeurl.URL
}

// UnmarshalText calls url.Parse
func (u *URL) UnmarshalText(p []byte) error {
	nu, err := nativeurl.Parse(string(p))
	if err != nil {
		return err
	}
	// (*u) = URL(*nu)
	u.URL = nativeurl.URL(*nu)
	return nil
}

// MarshalText just calls String()
func (u *URL) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}
