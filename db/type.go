package db

import "fmt"

// Type is the type of an item in the database.
type Type interface {
	String() string
	Equals(other Type) bool
}

type (
	// ListType stores an ordered homogenous sequence of elements
	ListType struct {
		ElemType Type
	}

	// HashmapType stores a mapping of keys to values
	HashmapType struct {
		KeyType, ValType Type
	}

	// FloatType is a 64-bit float
	FloatType struct{}
	// Float32Type is a 32-bit float
	Float32Type struct{}

	// IntType is a 64-bit signed int
	IntType struct{}
	// Int32Type is a 32-bit signed int
	Int32Type struct{}
	// Int16Type is a 16-bit signed int
	Int16Type struct{}
	// Int8Type is a 8-bit signed int
	Int8Type struct{}

	// UintType is a 64-bit unsigned int
	UintType struct{}
	// Uint32Type is a 32-bit unsigned int
	Uint32Type struct{}
	// Uint16Type is a 16-bit unsigned int
	Uint16Type struct{}
	// Uint8Type is a 8-bit unsigned int
	Uint8Type struct{}

	// StringType stores a unicode string
	StringType struct{}
	// BoolType stores a boolean true/false value
	BoolType struct{}
	// RegexpType stores a regexp which can be used to match patterns in strings
	RegexpType struct{}
	// AnyType allows any type, but not item is Any
	AnyType struct{}
)

func (l *ListType) String() string { return fmt.Sprintf("[%s]", l.ElemType) }

// Equals checks whether two types are equal
func (l *ListType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	o, ok := other.(*ListType)
	return ok && l.ElemType.Equals(o.ElemType)
}

func (h *HashmapType) String() string { return fmt.Sprintf("<%s:%s>", h.KeyType, h.ValType) }

// Equals checks whether two types are equal
func (h *HashmapType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	o, ok := other.(*HashmapType)
	return ok && h.KeyType.Equals(o.KeyType) && h.ValType.Equals(o.ValType)
}

func (f *FloatType) String() string { return "float" }

// Equals checks whether two types are equal
func (f *FloatType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Float32Type) String() string { return "float32" }

// Equals checks whether two types are equal
func (f *Float32Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *IntType) String() string { return "int" }

// Equals checks whether two types are equal
func (f *IntType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Int32Type) String() string { return "int32" }

// Equals checks whether two types are equal
func (f *Int32Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Int16Type) String() string { return "int16" }

// Equals checks whether two types are equal
func (f *Int16Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Int8Type) String() string { return "int8" }

// Equals checks whether two types are equal
func (f *Int8Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *UintType) String() string { return "uint" }

// Equals checks whether two types are equal
func (f *UintType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Uint32Type) String() string { return "uint32" }

// Equals checks whether two types are equal
func (f *Uint32Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Uint16Type) String() string { return "uint16" }

// Equals checks whether two types are equal
func (f *Uint16Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *Uint8Type) String() string { return "uint8" }

// Equals checks whether two types are equal
func (f *Uint8Type) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *StringType) String() string { return "string" }

// Equals checks whether two types are equal
func (f *StringType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *BoolType) String() string { return "bool" }

// Equals checks whether two types are equal
func (f *BoolType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *RegexpType) String() string { return "regexp" }

// Equals checks whether two types are equal
func (f *RegexpType) Equals(other Type) bool {
	if other.String() == "any" {
		return true
	}
	return f.String() == other.String()
}

func (f *AnyType) String() string { return "any" }

// Equals checks whether two types are equal
func (f *AnyType) Equals(other Type) bool {
	return true
}
