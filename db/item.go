package db

// A Comparison is one of the comparison operators.
type Comparison int

// The different kinds of comparison
const (
	_ Comparison = iota

	Equal
	NotEqual
	Less
	More
	LessOrEqual
	MoreOrEqual
	Regexp
)

const (
	// StatusOK means that the operation has been carried out successfully
	StatusOK = "OK"

	// StatusNOOP means that the operation cannot be performed on the type
	StatusNOOP = "NOOP"

	// StatusIndex means that an invalid index or key was requested
	StatusIndex = "IDX_ERR"

	// StatusError means that an unknown error has occurred
	StatusError = "ERR"
)

// An Item is any object in the database, such as a primitive number object or
// something more complicated like a hashmap.
type Item interface {
	// Type returns a string representing the type of the Item
	Type() string

	// Raw returns a Go value to represent the Item
	Raw() interface{}

	GetIndex(index int) (result Item, status string)
	GetKey(key Item) (result Item, status string)
	GetField(key string) (result Item, status string)
	SetIndex(index int, to Item) (status string)
	SetKey(key Item, to Item) (status string)
	SetField(key string, to Item) (status string)
	Compare(kind Comparison, other Item) (result bool, status string)
	Append(items ...Item) (status string)
	Prepend(items ...Item) (status string)
}

type itemDefaults struct{}

func (i *itemDefaults) GetIndex(index int) (result Item, status string) {
	return nil, StatusNOOP
}

func (i *itemDefaults) GetKey(key Item) (result Item, status string) {
	return nil, StatusNOOP
}

func (i *itemDefaults) GetField(key string) (result Item, status string) {
	return nil, StatusNOOP
}

func (i *itemDefaults) SetIndex(index int, to Item) (status string) {
	return StatusNOOP
}

func (i *itemDefaults) SetKey(key Item, to Item) (status string) {
	return StatusNOOP
}

func (i *itemDefaults) SetField(key string, to Item) (status string) {
	return StatusNOOP
}

func (i *itemDefaults) Compare(kind Comparison, other Item) (result bool, status string) {
	return false, StatusNOOP
}

func (i *itemDefaults) Append(items ...Item) (status string) {
	return StatusNOOP
}

func (i *itemDefaults) Prepend(items ...Item) (status string) {
	return StatusNOOP
}
