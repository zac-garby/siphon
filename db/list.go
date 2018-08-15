package db

import (
	"strings"
)

// A List stores an ordered sequence of items.
type List struct {
	*itemDefaults

	value   []Item
	valType Type
}

// NewList makes a new list with the given values.
func NewList(valType Type, vals ...Item) *List {
	return &List{
		value:   vals,
		valType: valType,
	}
}

// Type returns the type of an item.
func (l *List) Type() Type {
	return &ListType{
		ElemType: l.valType,
	}
}

func (l *List) String() string {
	str := &strings.Builder{}

	str.WriteByte('[')

	for i, item := range l.value {
		if i > 0 {
			str.WriteString(", ")
		}

		if i > 10 {
			str.WriteString("...")
			break
		}

		str.WriteString(item.String())
	}

	str.WriteByte(']')

	return str.String()
}

// JSON returns a JSON representation of an item
func (l *List) JSON() string {
	str := &strings.Builder{}

	str.WriteByte('[')

	for i, item := range l.value {
		if i > 0 {
			str.WriteString(", ")
		}

		str.WriteString(item.JSON())
	}

	str.WriteByte(']')

	return str.String()
}

// Set sets the value of the item to the given value
func (l *List) Set(val interface{}) (err error) {
	slice, ok := val.([]interface{})
	if !ok {
		return newError(ErrType, "expected a list value")
	}

	newList := make([]Item, len(slice))

	for i, item := range slice {
		newItem := MakeZeroValue(l.valType)
		if err := newItem.Set(item); err != nil {
			return err
		}

		newList[i] = newItem
	}

	l.value = newList

	return nil
}

// GetKey returns the item at the given key, provided the key is an integer.
func (l *List) GetKey(key Item) (result Item, err error) {
	val, ok := castNumeric(key)
	if !ok {
		return nil, newError(ErrType, "can only index a list with a numeric type")
	}

	index := int(val)

	if index < 0 || index >= len(l.value) {
		return nil, newError(ErrIndex, "index out of bounds")
	}

	return l.value[index], nil
}

// SetKey sets the item at the given key to something, provided the key is
// an integer.
func (l *List) SetKey(key Item, to Item) (err error) {
	val, ok := castNumeric(key)
	if !ok {
		return newError(ErrType, "can only index a list with a numeric type")
	}

	index := int(val)

	if index < 0 || index >= len(l.value) {
		return newError(ErrIndex, "index out of bounds")
	}

	l.value[index] = to
	return nil
}

// Filter returns a new list with all members of l which pass through the
// filter.
func (l *List) Filter(field string, kind Comparison, other Item) (result Item, err error) {
	result = &List{
		valType: l.valType,
		value:   make([]Item, 0, len(l.value)/2), // initialise with capacity as len()/2
	}

	for _, i := range l.value {
		var predicate bool

		if field == "" {
			pred, err := i.Compare(kind, other)
			if err != nil {
				return nil, err
			}

			predicate = pred
		} else {
			val, err := i.GetField(field)
			if err != nil {
				return nil, err
			}

			pred, err := val.Compare(kind, other)
			if err != nil {
				return nil, err
			}

			predicate = pred
		}

		if predicate {
			result.Append(i)
		}
	}

	return result, nil
}

// Append appends an item to the list.
func (l *List) Append(items ...Item) (err error) {
	l.value = append(l.value, items...)
	return nil
}

// Prepend pushes an item to the beginning of the list. They will remain in
// the same order, so [1, 2, 3] prepend [4, 5, 6] will result in [4, 5, 6, 1, 2, 3].
func (l *List) Prepend(items ...Item) (err error) {
	l.value = append(items, l.value...)
	return nil
}

// AppendJSON appends an item encoded as JSON to the list.
func (l *List) AppendJSON(json interface{}) (err error) {
	item := MakeZeroValue(l.valType)

	if err := item.Set(json); err != nil {
		return err
	}

	return l.Append(item)
}

// PrependJSON prepends an item encoded as JSON to the list.
func (l *List) PrependJSON(json interface{}) (err error) {
	item := MakeZeroValue(l.valType)

	if err := item.Set(json); err != nil {
		return err
	}

	return l.Prepend(item)
}
