package db

import "fmt"

// An Int is just a basic 64-bit integer.
type Int struct {
	*itemDefaults

	value int64
}

// NewInt makes a new integer item.
func NewInt(val int64) *Int {
	return &Int{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Int) Type() Type {
	return &IntType{}
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Int) JSON() string {
	return i.String()
}

// Set sets the value of the item to the given value
func (i *Int) Set(val interface{}) (err error) {
	ival, ok := val.(float64)
	if !ok {
		return newError(ErrType, "expected an integer value")
	}

	i.value = int64(ival)

	return nil
}

// Compare compares two items
func (i *Int) Compare(kind Comparison, other Item) (result bool, err error) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, newError(ErrNOOP, "can only compare ints with numeric types (ints, floats, uints, ...)")
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, nil

	case NotEqual:
		return sval != oval, nil

	case Less:
		return sval < oval, nil

	case More:
		return sval > oval, nil

	case LessOrEqual:
		return sval <= oval, nil

	case MoreOrEqual:
		return sval >= oval, nil

	default:
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >= comparisons are supported on ints")
	}
}

/////////////////////////////////////////////////

// An Int32 is just a basic 32-bit integer.
type Int32 struct {
	*itemDefaults

	value int32
}

// NewInt32 makes a new integer item.
func NewInt32(val int32) *Int32 {
	return &Int32{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Int32) Type() Type {
	return &Int32Type{}
}

func (i *Int32) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Int32) JSON() string {
	return i.String()
}

// Set sets the value of the item to the given value
func (i *Int32) Set(val interface{}) (err error) {
	ival, ok := val.(float64)
	if !ok {
		return newError(ErrType, "expected an integer value")
	}

	i.value = int32(ival)

	return nil
}

// Compare compares two items
func (i *Int32) Compare(kind Comparison, other Item) (result bool, err error) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, newError(ErrNOOP, "can only compare ints with numeric types (ints, floats, uints, ...)")
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, nil

	case NotEqual:
		return sval != oval, nil

	case Less:
		return sval < oval, nil

	case More:
		return sval > oval, nil

	case LessOrEqual:
		return sval <= oval, nil

	case MoreOrEqual:
		return sval >= oval, nil

	default:
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >= comparisons are supported on ints")
	}
}

/////////////////////////////////////////////////

// An Int16 is just a basic 16-bit integer.
type Int16 struct {
	*itemDefaults

	value int16
}

// NewInt16 makes a new integer item.
func NewInt16(val int16) *Int16 {
	return &Int16{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Int16) Type() Type {
	return &Int16Type{}
}

func (i *Int16) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Int16) JSON() string {
	return i.String()
}

// Set sets the value of the item to the given value
func (i *Int16) Set(val interface{}) (err error) {
	ival, ok := val.(float64)
	if !ok {
		return newError(ErrType, "expected an integer value")
	}

	i.value = int16(ival)

	return nil
}

// Compare compares two items
func (i *Int16) Compare(kind Comparison, other Item) (result bool, err error) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, newError(ErrNOOP, "can only compare ints with numeric types (ints, floats, uints, ...)")
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, nil

	case NotEqual:
		return sval != oval, nil

	case Less:
		return sval < oval, nil

	case More:
		return sval > oval, nil

	case LessOrEqual:
		return sval <= oval, nil

	case MoreOrEqual:
		return sval >= oval, nil

	default:
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >= comparisons are supported on ints")
	}
}

/////////////////////////////////////////////////

// An Int8 is just a basic 8-bit integer.
type Int8 struct {
	*itemDefaults

	value int8
}

// NewInt8 makes a new integer item.
func NewInt8(val int8) *Int8 {
	return &Int8{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Int8) Type() Type {
	return &Int8Type{}
}

func (i *Int8) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Int8) JSON() string {
	return i.String()
}

// Set sets the value of the item to the given value
func (i *Int8) Set(val interface{}) (err error) {
	ival, ok := val.(float64)
	if !ok {
		return newError(ErrType, "expected an integer value")
	}

	i.value = int8(ival)

	return nil
}

// Compare compares two items
func (i *Int8) Compare(kind Comparison, other Item) (result bool, err error) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, newError(ErrNOOP, "can only compare ints with numeric types (ints, floats, uints, ...)")
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, nil

	case NotEqual:
		return sval != oval, nil

	case Less:
		return sval < oval, nil

	case More:
		return sval > oval, nil

	case LessOrEqual:
		return sval <= oval, nil

	case MoreOrEqual:
		return sval >= oval, nil

	default:
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >= comparisons are supported on ints")
	}
}
