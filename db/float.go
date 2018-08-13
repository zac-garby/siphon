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
func (f *Float) Set(val interface{}) (err error) {
	fval, ok := val.(float64)
	if !ok {
		return newError(ErrType, "expected a float value")
	}

	f.value = fval

	return nil
}

// Compare compares two items
func (f *Float) Compare(kind Comparison, other Item) (result bool, err error) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, newError(ErrNOOP, "can only compare floats with numeric types (ints, floats, uints, ...)")
	}

	switch kind {
	case Equal:
		return f.value == oval, nil

	case NotEqual:
		return f.value != oval, nil

	case Less:
		return f.value < oval, nil

	case More:
		return f.value > oval, nil

	case LessOrEqual:
		return f.value <= oval, nil

	case MoreOrEqual:
		return f.value >= oval, nil

	default:
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >= comparisons are supported on floats")
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
func (f *Float32) Set(val interface{}) (err error) {
	fval, ok := val.(float64)
	if !ok {
		return newError(ErrType, "expected a float value")
	}

	f.value = float32(fval)

	return nil
}

// Compare compares two items
func (f *Float32) Compare(kind Comparison, other Item) (result bool, err error) {
	oval, ok := castNumeric(other)
	if !ok {
		return false, newError(ErrNOOP, "can only compare floats with numeric types (ints, floats, uints, ...)")
	}

	sval, _ := castNumeric(f)

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
		return false, newError(ErrNOOP, "only =, !=, <, >, <=, >= comparisons are supported on floats")
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
