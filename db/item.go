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
	RegexpMatch
)

func stringToComparison(str string) (cmp Comparison, ok bool) {
	switch str {
	case "=":
		return Equal, true
	case "!=":
		return NotEqual, true
	case "<":
		return Less, true
	case ">":
		return More, true
	case "<=":
		return LessOrEqual, true
	case ">=":
		return MoreOrEqual, true
	case "~":
		return RegexpMatch, true
	}

	return -1, false
}

// ErrorType says what the cause of an error is
type ErrorType string

const (
	// ErrNOOP means that the operation cannot be performed on the type
	ErrNOOP ErrorType = "no operation"

	// ErrIndex means that an invalid index or key was requested
	ErrIndex ErrorType = "index error"

	// ErrUnknown means that an unknown error has occurred
	ErrUnknown ErrorType = "error"

	// ErrType means that a type error has occurred
	ErrType ErrorType = "invalid type"

	// ErrNoType means that an invalid type was specified in the schema
	ErrNoType ErrorType = "undefined type"
)

// An Item is any object in the database, such as a primitive number object or
// something more complicated like a hashmap.
type Item interface {
	// Type returns a string representing the type of the Item
	Type() Type

	String() string
	JSON() string

	Set(val interface{}) (err error)
	GetKey(key Item) (result Item, err error)
	GetField(key string) (result Item, err error)
	SetKey(key Item, to Item) (err error)
	SetKeyJSON(key interface{}, to interface{}) (err error)
	UnsetKey(key Item) (err error)
	UnsetKeyJSON(key interface{}) (err error)
	SetField(key string, to Item) (err error)
	Compare(kind Comparison, other Item) (result bool, err error)
	Filter(field string, kind Comparison, other Item) (result Item, err error)
	Append(items ...Item) (err error)
	AppendJSON(json interface{}) (err error)
	Prepend(items ...Item) (err error)
	PrependJSON(json interface{}) (err error)
	Empty() (err error)
}

type itemDefaults struct{}

func (i *itemDefaults) Set(val interface{}) (err error) {
	return newError(ErrNOOP, "set not supported")
}

func (i *itemDefaults) GetKey(key Item) (result Item, err error) {
	return nil, newError(ErrNOOP, "getkey not supported")
}

func (i *itemDefaults) GetField(key string) (result Item, err error) {
	return nil, newError(ErrNOOP, "getfield not supported")
}

func (i *itemDefaults) SetKey(key Item, to Item) (err error) {
	return newError(ErrNOOP, "setkey not supported")
}

func (i *itemDefaults) SetKeyJSON(key interface{}, to interface{}) (err error) {
	return newError(ErrNOOP, "setkey json not supported")
}

func (i *itemDefaults) SetField(key string, to Item) (err error) {
	return newError(ErrNOOP, "setfield not supported")
}

func (i *itemDefaults) Compare(kind Comparison, other Item) (result bool, err error) {
	return false, newError(ErrNOOP, "compare not supported")
}

func (i *itemDefaults) Filter(field string, kind Comparison, other Item) (result Item, err error) {
	return nil, newError(ErrNOOP, "filter not supported")
}

func (i *itemDefaults) Append(items ...Item) (err error) {
	return newError(ErrNOOP, "append not supported")
}

func (i *itemDefaults) AppendJSON(json interface{}) (err error) {
	return newError(ErrNOOP, "append json not supported")
}

func (i *itemDefaults) Prepend(items ...Item) (err error) {
	return newError(ErrNOOP, "prepend not supported")
}

func (i *itemDefaults) PrependJSON(json interface{}) (err error) {
	return newError(ErrNOOP, "prepend json not supported")
}

func (i *itemDefaults) UnsetKey(key Item) (err error) {
	return newError(ErrNOOP, "unsetkey not supported")
}

func (i *itemDefaults) UnsetKeyJSON(key interface{}) (err error) {
	return newError(ErrNOOP, "unsetkey json not supported")
}

func (i *itemDefaults) Empty() (err error) {
	return newError(ErrNOOP, "empty not supported")
}
