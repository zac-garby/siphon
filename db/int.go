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
func (i *Int) Type() string {
	return TypeInt
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.value)
}

// Compare compares two items
func (i *Int) Compare(kind Comparison, other Item) (result bool, status string) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, StatusNOOP
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, StatusOK

	case NotEqual:
		return sval != oval, StatusOK

	case Less:
		return sval < oval, StatusOK

	case More:
		return sval > oval, StatusOK

	case LessOrEqual:
		return sval <= oval, StatusOK

	case MoreOrEqual:
		return sval >= oval, StatusOK

	default:
		return false, StatusNOOP
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
func (i *Int32) Type() string {
	return TypeInt
}

func (i *Int32) String() string {
	return fmt.Sprintf("%d", i.value)
}

// Compare compares two items
func (i *Int32) Compare(kind Comparison, other Item) (result bool, status string) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, StatusNOOP
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, StatusOK

	case NotEqual:
		return sval != oval, StatusOK

	case Less:
		return sval < oval, StatusOK

	case More:
		return sval > oval, StatusOK

	case LessOrEqual:
		return sval <= oval, StatusOK

	case MoreOrEqual:
		return sval >= oval, StatusOK

	default:
		return false, StatusNOOP
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
func (i *Int16) Type() string {
	return TypeInt
}

func (i *Int16) String() string {
	return fmt.Sprintf("%d", i.value)
}

// Compare compares two items
func (i *Int16) Compare(kind Comparison, other Item) (result bool, status string) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, StatusNOOP
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, StatusOK

	case NotEqual:
		return sval != oval, StatusOK

	case Less:
		return sval < oval, StatusOK

	case More:
		return sval > oval, StatusOK

	case LessOrEqual:
		return sval <= oval, StatusOK

	case MoreOrEqual:
		return sval >= oval, StatusOK

	default:
		return false, StatusNOOP
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
func (i *Int8) Type() string {
	return TypeInt
}

func (i *Int8) String() string {
	return fmt.Sprintf("%d", i.value)
}

// Compare compares two items
func (i *Int8) Compare(kind Comparison, other Item) (result bool, status string) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, StatusNOOP
	}

	sval, _ := castNumeric(i)

	switch kind {
	case Equal:
		return sval == oval, StatusOK

	case NotEqual:
		return sval != oval, StatusOK

	case Less:
		return sval < oval, StatusOK

	case More:
		return sval > oval, StatusOK

	case LessOrEqual:
		return sval <= oval, StatusOK

	case MoreOrEqual:
		return sval >= oval, StatusOK

	default:
		return false, StatusNOOP
	}
}
