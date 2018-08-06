package db

// A String is just a basic string, with unicode support.
type String struct {
	*itemDefaults

	value string
}

// NewString makes a new string with the given initial value.
func NewString(val string) *String {
	return &String{
		value: val,
	}
}

// Type returns the type of the Item
func (s *String) Type() string {
	return TypeString
}

// Raw returns a Go value to represent the Item
func (s *String) Raw() interface{} {
	return s.value
}
