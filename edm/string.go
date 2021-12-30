package edm

type String string

func (s String) String() string {
	return string(s)
}

func (s String) IsEmpty() bool {
	return string(s) == ""
}
