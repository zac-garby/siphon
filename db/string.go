package db

import (
	"regexp"
)

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

// Compare compares an item with another item
func (s *String) Compare(kind Comparison, other Item) (result bool, status string) {
	switch kind {
	case Equal:
		return s.value == other.Raw(), StatusOK

	case NotEqual:
		return s.value != other.Raw(), StatusOK

	case Less:
		if other.Type() != TypeString {
			return false, StatusOK
		}

		return s.value < other.Raw().(string), StatusOK

	case More:
		if other.Type() != TypeString {
			return false, StatusOK
		}

		return s.value > other.Raw().(string), StatusOK

	case LessOrEqual:
		if other.Type() != TypeString {
			return false, StatusOK
		}

		return s.value <= other.Raw().(string), StatusOK

	case MoreOrEqual:
		if other.Type() != TypeString {
			return false, StatusOK
		}

		return s.value >= other.Raw().(string), StatusOK

	case RegexpMatch:
		if other.Type() != TypeRegexp {
			return false, StatusNOOP
		}

		reg, err := regexp.Compile(other.Raw().(string))
		if err != nil {
			return false, StatusError
		}

		return reg.MatchString(s.value), StatusOK

	default:
		return false, StatusNOOP
	}
}
