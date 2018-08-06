package db

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
func (f *Float) Type() string {
	return TypeFloat
}

// Raw returns a Go value to represent the Item
func (f *Float) Raw() interface{} {
	return f.value
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
func (f *Float32) Type() string {
	return TypeFloat32
}

// Raw returns a Go value to represent the Item
func (f *Float32) Raw() interface{} {
	return f.value
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
	switch val := item.Raw().(type) {
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case int64:
		return float64(val), true
	case int32:
		return float64(val), true
	case int16:
		return float64(val), true
	case int8:
		return float64(val), true
	case uint64:
		return float64(val), true
	case uint32:
		return float64(val), true
	case uint16:
		return float64(val), true
	case uint8:
		return float64(val), true
	default:
		return 0, false
	}
}
