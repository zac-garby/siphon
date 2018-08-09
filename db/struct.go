package db

import "strings"

// A Struct stores pairs of corresponding names and values. Each field has
// a type, and its value can only be that type.
type Struct struct {
	*itemDefaults

	ty    *StructType
	value map[string]Item
}

// NewStruct makes a new struct according to the given struct type.
func NewStruct(ty *StructType) *Struct {
	s := &Struct{
		ty:    ty,
		value: make(map[string]Item),
	}

	for k, t := range ty.Fields {
		s.value[k] = MakeZeroValue(t)
	}

	return s
}

// Type returns the type of an item
func (s *Struct) Type() Type {
	return s.ty
}

func (s *Struct) String() string {
	str := &strings.Builder{}
	str.WriteString(s.ty.Name)
	str.WriteByte('{')

	i := 0
	for name, val := range s.value {
		if i > 0 {
			str.WriteString(", ")
		}
		str.WriteString(name)
		str.WriteString(": ")
		str.WriteString(val.String())
		i++
	}

	str.WriteByte('}')
	return str.String()
}

// GetField returns the field named 'key' in the struct.
func (s *Struct) GetField(key string) (result Item, status string) {
	val, ok := s.value[key]
	if !ok {
		return nil, StatusIndex
	}

	return val, StatusOK
}

// SetField sets the field named 'key' to the given value.
func (s *Struct) SetField(key string, to Item) (status string) {
	reqType, ok := s.ty.Fields[key]
	if !ok {
		return StatusIndex
	}
	if !to.Type().Equals(reqType) {
		return StatusType
	}

	s.value[key] = to

	return StatusOK
}
