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

// JSON returns a JSON representation of an item
func (s *String) JSON() string {
	return s.String()
}

// Set sets the value of the item to the given value
func (s *String) Set(val interface{}) (err error) {
	sval, ok := val.(string)
	if !ok {
		return newError(ErrType, "expected a string value")
	}

	s.value = sval

	return nil
}

// Compare compares an item with another item
func (s *String) Compare(kind Comparison, other Item) (result bool, err error) {
	if kind == RegexpMatch {
		or, ok := other.(*Regexp)
		if !ok {
			return false, newError(ErrNOOP, "the right hand argument to ~ must be a regexp")
		}

		reg, err := regexp.Compile(or.value)
		if err != nil {
			return false, newError(ErrUnknown, "regexp could not be compiled")
		}

		return reg.MatchString(s.value), nil
	}

	os, ok := other.(*String)
	if !ok {
		return false, nil
	}

	switch kind {
	case Equal:
		return s.value == os.value, nil

	case NotEqual:
		return s.value != os.value, nil

	case Less:
		if other.Type().Equals(&StringType{}) {
			return false, nil
		}

		return s.value < os.value, nil

	case More:
		if other.Type().Equals(&StringType{}) {
			return false, nil
		}

		return s.value > os.value, nil

	case LessOrEqual:
		if other.Type().Equals(&StringType{}) {
			return false, nil
		}

		return s.value <= os.value, nil

	case MoreOrEqual:
		if other.Type().Equals(&StringType{}) {
			return false, nil
		}

		return s.value >= os.value, nil

	default:
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >=, and ~ comparisons are supported on strings")
	}
}
