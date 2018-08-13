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

// JSON returns a JSON representation of an item
func (s *Struct) JSON() string {
	str := &strings.Builder{}
	str.WriteByte('{')

	i := 0
	for name, val := range s.value {
		if i > 0 {
			str.WriteString(", ")
		}
		str.WriteString("\"" + name + "\"")
		str.WriteString(": ")
		str.WriteString(val.JSON())
		i++
	}

	str.WriteByte('}')
	return str.String()
}

// Set sets the value of the item to the given value
func (s *Struct) Set(val interface{}) (err error) {
	hval, ok := val.(map[string]interface{})
	if !ok {
		return newError(ErrType, "expected a hashmap value whose keys are strings")
	}

	if len(hval) != len(s.value) {
		return newError(ErrType, "all fields must be present to set a struct's value")
	}

	newMap := make(map[string]Item, len(hval))

	for k, ty := range s.ty.Fields {
		newVal := MakeZeroValue(ty)
		newInterVal, ok := hval[k]
		if !ok {
			return newError(ErrType, "all fields must be present to set a struct's value")
		}

		if err := newVal.Set(newInterVal); err != nil {
			return err
		}

		newMap[k] = newVal
	}

	s.value = newMap

	return nil
}

// GetField returns the field named 'key' in the struct.
func (s *Struct) GetField(key string) (result Item, err error) {
	val, ok := s.value[key]
	if !ok {
		return nil, newError(ErrIndex, "cannot retrieve undefined field %s", key)
	}

	return val, nil
}

// SetField sets the field named 'key' to the given value.
func (s *Struct) SetField(key string, to Item) (err error) {
	reqType, ok := s.ty.Fields[key]
	if !ok {
		return newError(ErrIndex, "cannot retrieve undefined field %s")
	}
	if !to.Type().Equals(reqType) {
		return newError(
			ErrType,
			"field %s of type %s cannot be assigned to a %s",
			key,
			reqType,
			to.Type(),
		)
	}

	s.value[key] = to

	return nil
}
