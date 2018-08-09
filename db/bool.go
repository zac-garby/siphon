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

// Compare compares two items
func (b *Bool) Compare(kind Comparison, other Item) (result bool, status string) {
	ob, ok := other.(*Bool)
	if !ok {
		return false, StatusOK
	}

	switch kind {
	case Equal:
		return b.value == ob.value, StatusOK
	case NotEqual:
		return b.value != ob.value, StatusOK
	default:
		return false, StatusNOOP
	}
}
