package db

import "fmt"

// A Bool is either true or false.
type Bool struct {
	*itemDefaults

	value bool
}

// NewBool makes a new boolean.
func NewBool(val bool) *Bool {
	return &Bool{
		value: val,
	}
}

// Type returns the type of an item
func (b *Bool) Type() Type {
	return &BoolType{}
}

func (b *Bool) String() string {
	return fmt.Sprintf("%t", b.value)
}

// JSON returns a JSON representation of the item
func (b *Bool) JSON() string {
	return b.String()
}

// Set sets the value of the item to the given value
func (b *Bool) Set(val interface{}) (err error) {
	bval, ok := val.(bool)
	if !ok {
		return newError(ErrType, "expected a boolean value")
	}

	b.value = bval

	return nil
}

// Compare compares two items
func (b *Bool) Compare(kind Comparison, other Item) (result bool, err error) {
	ob, ok := other.(*Bool)
	if !ok {
		return false, nil
	}

	switch kind {
	case Equal:
		return b.value == ob.value, nil
	case NotEqual:
		return b.value != ob.value, nil
	default:
		return false, newError(ErrNOOP, "only = and != are supported on booleans")
	}
}
