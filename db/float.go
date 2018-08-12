package db

import (
	"fmt"
	"strconv"
)

// A Float is just a basic 64-bit floating point number.
type Float struct {
	*itemDefaults

	value float64
}

// NewFloat makes a new float with the given initial value.
func NewFloat(val float64) *Float {
	return &Float{
		value: val,
	}
}

// Type returns the type of the Item
func (f *Float) Type() Type {
	return &FloatType{}
}

func (f *Float) String() string {
	return fmt.Sprintf("%v", f.value)
}

// JSON returns a JSON representation of the item
func (f *Float) JSON() string {
	return f.String()
}

// Set sets the value of the item to the given value
func (f *Float) Set(val interface{}) (status string) {
	fval, ok := val.(float64)
	if !ok {
		return StatusType
	}

	f.value = fval

	return StatusOK
}

// Compare compares two items
func (f *Float) Compare(kind Comparison, other Item) (result bool, status string) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, StatusNOOP
	}

	switch kind {
	case Equal:
		return f.value == oval, StatusOK

	case NotEqual:
		return f.value != oval, StatusOK

	case Less:
		return f.value < oval, StatusOK

	case More:
		return f.value > oval, StatusOK

	case LessOrEqual:
		return f.value <= oval, StatusOK

	case MoreOrEqual:
		return f.value >= oval, StatusOK

	default:
		return false, StatusNOOP
	}
}

//////////////////////////////////////////////

// A Float32 is just a basic 32-bit floating point number.
type Float32 struct {
	*itemDefaults

	value float32
}

// NewFloat32 makes a new Float32 with the given initial value.
func NewFloat32(val float32) *Float32 {
	return &Float32{
		value: val,
	}
}

// Type returns the type of the Item
func (f *Float32) Type() Type {
	return &Float32Type{}
}

func (f *Float32) String() string {
	return fmt.Sprintf("%v", f.value)
}

// JSON returns a JSON representation of the item
func (f *Float32) JSON() string {
	return f.String()
}

// Set sets the value of the item to the given value
func (f *Float32) Set(val interface{}) (status string) {
	fval, ok := val.(float64)
	if !ok {
		return StatusType
	}

	f.value = float32(fval)

	return StatusOK
}

// Compare compares two items
func (f *Float32) Compare(kind Comparison, other Item) (result bool, status string) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, StatusNOOP
	}

	sval, _ := castNumeric(f)

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

//////////////////////////////////////////////

func castNumeric(item Item) (val float64, ok bool) {
	val, err := strconv.ParseFloat(item.String(), 64)
	if err != nil {
		return 0, false
	}

	return val, true
}
