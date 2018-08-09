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
func (s *String) Type() Type {
	return &StringType{}
}

func (s *String) String() string {
	return "\"" + s.value + "\""
}

// Compare compares an item with another item
func (s *String) Compare(kind Comparison, other Item) (result bool, status string) {
	if kind == RegexpMatch {
		or, ok := other.(*Regexp)
		if !ok {
			return false, StatusNOOP
		}

		reg, err := regexp.Compile(or.value)
		if err != nil {
			return false, StatusError
		}

		return reg.MatchString(s.value), StatusOK
	}

	os, ok := other.(*String)
	if !ok {
		return false, StatusOK
	}

	switch kind {
	case Equal:
		return s.value == os.value, StatusOK

	case NotEqual:
		return s.value != os.value, StatusOK

	case Less:
		if other.Type().Equals(&StringType{}) {
			return false, StatusOK
		}

		return s.value < os.value, StatusOK

	case More:
		if other.Type().Equals(&StringType{}) {
			return false, StatusOK
		}

		return s.value > os.value, StatusOK

	case LessOrEqual:
		if other.Type().Equals(&StringType{}) {
			return false, StatusOK
		}

		return s.value <= os.value, StatusOK

	case MoreOrEqual:
		if other.Type().Equals(&StringType{}) {
			return false, StatusOK
		}

		return s.value >= os.value, StatusOK

	default:
		return false, StatusNOOP
	}
}
