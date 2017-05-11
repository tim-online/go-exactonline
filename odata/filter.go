package odata

func NewFilter() *Filter {
	return &Filter{}
}

type Filter struct{}

func (f *Filter) MarshalSchema() string {
	// return strings.Join(t.Values, ",")
	return ""
}
