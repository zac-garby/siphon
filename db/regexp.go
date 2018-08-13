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
func (r *Regexp) Type() Type {
	return &RegexpType{}
}

func (r *Regexp) String() string {
	return "/" + r.value + "/"
}

// JSON returns a JSON representation of an item
func (r *Regexp) JSON() string {
	return "\"" + r.value + "\""
}

// Set sets the value of the item to the given value
func (r *Regexp) Set(val interface{}) (err error) {
	sval, ok := val.(string)
	if !ok {
		return newError(ErrType, "expected a string value")
	}

	r.value = sval

	return nil
}
