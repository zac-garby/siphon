package db

// A Regexp is used to check whether strings follow a particular pattern
type Regexp struct {
	*itemDefaults

	value string
}

// NewRegexp makes a new regexp item.
func NewRegexp(val string) *Regexp {
	return &Regexp{
		value: val,
	}
}

// Type returns the type of an item
func (r *Regexp) Type() string {
	return TypeRegexp
}

// Raw returns a Go value to represent the item
func (r *Regexp) Raw() interface{} {
	return r.value
}
