package db

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
func (b *Bool) Type() string {
	return TypeBool
}

// Raw returns a Go value representing the item
func (b *Bool) Raw() interface{} {
	return b.value
}

// Compare compares two items
func (b *Bool) Compare(kind Comparison, other Item) (result bool, status string) {
	switch kind {
	case Equal:
		return b.value == other.Raw(), StatusOK
	case NotEqual:
		return b.value != other.Raw(), StatusOK
	default:
		return false, StatusNOOP
	}
}
