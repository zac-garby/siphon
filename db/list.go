package db

import "strings"

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

// GetKey returns the item at the given key, provided the key is an integer.
func (l *List) GetKey(key Item) (result Item, status string) {
	val, ok := castNumeric(key)
	if !ok {
		return nil, StatusType
	}

	index := int(val)

	if index < 0 || index >= len(l.value) {
		return nil, StatusIndex
	}

	return l.value[index], StatusOK
}

// SetKey sets the item at the given key to something, provided the key is
// an integer.
func (l *List) SetKey(key Item, to Item) (status string) {
	val, ok := castNumeric(key)
	if !ok {
		return StatusType
	}

	index := int(val)

	if index < 0 || index >= len(l.value) {
		return StatusIndex
	}

	l.value[index] = to
	return StatusOK
}

// Filter returns a new list with all members of l which pass through the
// filter.
func (l *List) Filter(field string, kind Comparison, other Item) (result Item, status string) {
	result = &List{
		valType: l.valType,
		value:   make([]Item, 0, len(l.value)/2), // initialise with capacity as len()/2
	}

	for _, i := range l.value {
		var predicate bool

		if field == "" {
			pred, status := i.Compare(kind, other)
			if status != StatusOK {
				return nil, status
			}

			predicate = pred
		} else {
			val, status := i.GetField(field)
			if status != StatusOK {
				return nil, status
			}

			pred, status := val.Compare(kind, other)
			if status != StatusOK {
				return nil, status
			}

			predicate = pred
		}

		if predicate {
			result.Append(i)
		}
	}

	return result, StatusOK
}

// Append appends an item to the list.
func (l *List) Append(items ...Item) (status string) {
	l.value = append(l.value, items...)
	return StatusOK
}

// Prepend pushes an item to the beginning of the list. They will remain in
// the same order, so [1, 2, 3] prepend [4, 5, 6] will result in [4, 5, 6, 1, 2, 3].
func (l *List) Prepend(items ...Item) (status string) {
	l.value = append(items, l.value...)
	return StatusOK
}
