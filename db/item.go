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

// An Item is any object in the database, such as a primitive number object or
// something more complicated like a hashmap.
type Item interface {
	// Type returns a string representing the type of the Item
	Type() string

	// Raw returns a Go value to represent the Item
	Raw() interface{}

	GetIndex(index int) (result Item, status int)
	GetKey(key string) (result Item, status int)
	Compare(kind Comparison, other Item) (result bool, status int)
	Append(items ...Item) (status int)
	Prepend(items ...Item) (status int)
}
