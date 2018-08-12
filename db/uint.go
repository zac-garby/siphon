package db

import "fmt"

// An Uint is just a basic 64-bit unsigned integer.
type Uint struct {
	*itemDefaults

	value uint64
}

// NewUint makes a new unsigned integer item.
func NewUint(val uint64) *Uint {
	return &Uint{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Uint) Type() Type {
	return &UintType{}
}

func (i *Uint) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Uint) JSON() string {
	return fmt.Sprintf("%d", i.value)
}

// Set sets the value of the item to the given value
func (i *Uint) Set(val interface{}) (status string) {
	fval, ok := val.(float64)
	if !ok {
		return StatusType
	}

	i.value = uint64(fval)

	return StatusOK
}

// Compare compares two items
func (i *Uint) Compare(kind Comparison, other Item) (result bool, status string) {
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

// An Uint32 is just a basic 32-bit unsigned integer.
type Uint32 struct {
	*itemDefaults

	value uint32
}

// NewUint32 makes a new unsigned integer item.
func NewUint32(val uint32) *Uint32 {
	return &Uint32{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Uint32) Type() Type {
	return &Uint32Type{}
}

func (i *Uint32) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Uint32) JSON() string {
	return fmt.Sprintf("%d", i.value)
}

// Set sets the value of the item to the given value
func (i *Uint32) Set(val interface{}) (status string) {
	fval, ok := val.(float64)
	if !ok {
		return StatusType
	}

	i.value = uint32(fval)

	return StatusOK
}

// Compare compares two items
func (i *Uint32) Compare(kind Comparison, other Item) (result bool, status string) {
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

// An Uint16 is just a basic 16-bit unsigned integer.
type Uint16 struct {
	*itemDefaults

	value uint16
}

// NewUint16 makes a new unsigned integer item.
func NewUint16(val uint16) *Uint16 {
	return &Uint16{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Uint16) Type() Type {
	return &Uint16Type{}
}

func (i *Uint16) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Uint16) JSON() string {
	return fmt.Sprintf("%d", i.value)
}

// Set sets the value of the item to the given value
func (i *Uint16) Set(val interface{}) (status string) {
	fval, ok := val.(float64)
	if !ok {
		return StatusType
	}

	i.value = uint16(fval)

	return StatusOK
}

// Compare compares two items
func (i *Uint16) Compare(kind Comparison, other Item) (result bool, status string) {
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

// An Uint8 is just a basic 8-bit unsigned integer.
type Uint8 struct {
	*itemDefaults

	value uint8
}

// NewUint8 makes a new unsigned integer item.
func NewUint8(val uint8) *Uint8 {
	return &Uint8{
		value: val,
	}
}

// Type returns the type of an item.
func (i *Uint8) Type() Type {
	return &Uint8Type{}
}

func (i *Uint8) String() string {
	return fmt.Sprintf("%d", i.value)
}

// JSON returns a JSON representation of the item
func (i *Uint8) JSON() string {
	return fmt.Sprintf("%d", i.value)
}

// Set sets the value of the item to the given value
func (i *Uint8) Set(val interface{}) (status string) {
	fval, ok := val.(float64)
	if !ok {
		return StatusType
	}

	i.value = uint8(fval)

	return StatusOK
}

// Compare compares two items
func (i *Uint8) Compare(kind Comparison, other Item) (result bool, status string) {
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
