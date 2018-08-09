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

// GetIndex returns the item at the given index.
func (l *List) GetIndex(index int) (result Item, status string) {
	if index < 0 || index >= len(l.value) {
		return nil, StatusIndex
	}

	return l.value[index], StatusOK
}

// SetIndex sets the item at the given index to something.
func (l *List) SetIndex(index int, to Item) (status string) {
	if index < 0 || index >= len(l.value) {
		return StatusIndex
	}

	l.value[index] = to
	return StatusOK
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
