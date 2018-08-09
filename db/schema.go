package db

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

var schemaLexer = lexer.Must(lexer.Regexp(`(?P<Newline>\n)` +
	`|(?m)(\s+)` +
	`|(#.*$)` +
	`|(?P<Keyword>struct)` +
	`|(?P<Ident>[\p{L}\p{M}_-][\p{L}\p{M}\d_-]*)` +
	`|(?P<Punctuation>[:{}\[\]<>])`,
))

// SchemaParser parses schemas.
var SchemaParser *participle.Parser

func init() {
	parser, err := participle.Build(&Schema{}, participle.Lexer(schemaLexer))
	if err != nil {
		panic(err)
	}

	SchemaParser = parser
}

// A Schema is used to specify the structure and field types of a database.
type Schema struct {
	Sections []*SchemaSection `{ { Newline } @@ }`
}

// A SchemaSection is a field or struct definition in a schema.
type SchemaSection struct {
	Field  *SchemaField  `  @@`
	Struct *SchemaStruct `| @@`
}

// A SchemaField defines a field in the schema or in a struct.
type SchemaField struct {
	Name string      `@Ident`
	Type *SchemaType `":" @@ Newline`
}

// A SchemaType specifies the type of a field.
type SchemaType struct {
	Ident   string         `  @Ident`
	List    *SchemaType    `| "[" @@ "]"`
	Hashmap *SchemaMapType `| @@`
}

// A SchemaMapType represents a hashmap field.
type SchemaMapType struct {
	KeyType   *SchemaType `"<" @@ ":"`
	ValueType *SchemaType `@@ ">"`
}

// A SchemaStruct defines a new type, similar to a Go struct.
type SchemaStruct struct {
	Name   string         `"struct" @Ident`
	Fields []*SchemaField `"{" { { Newline } @@ { Newline } } "}"`
}

// MakeZeroValue makes a new Item which has the zero value of the
// given type.
func MakeZeroValue(t Type) Item {
	switch ty := t.(type) {
	case *ListType:
		return NewList(ty.ElemType)

	case *HashmapType:
		return NewHashmap(ty.KeyType, ty.ValType)

	case *FloatType:
		return NewFloat(0)
	case *Float32Type:
		return NewFloat32(0)

	case *IntType:
		return NewInt(0)
	case *Int32Type:
		return NewInt32(0)
	case *Int16Type:
		return NewInt16(0)
	case *Int8Type:
		return NewInt8(0)

	case *UintType:
		return NewUint(0)
	case *Uint32Type:
		return NewUint32(0)
	case *Uint16Type:
		return NewUint16(0)
	case *Uint8Type:
		return NewUint8(0)

	case *StringType:
		return NewString("")
	case *BoolType:
		return NewBool(false)
	case *RegexpType:
		return NewRegexp("")

	default:
		return nil
	}
}

// GetActualType takes a parsed schema type and returns an actual
// Type instance.
func GetActualType(st *SchemaType) Type {
	if li := st.List; li != nil {
		return &ListType{
			ElemType: GetActualType(li),
		}
	} else if hm := st.Hashmap; hm != nil {
		return &HashmapType{
			KeyType: GetActualType(hm.KeyType),
			ValType: GetActualType(hm.ValueType),
		}
	}

	switch id := st.Ident; id {
	case "float":
		return &FloatType{}
	case "float32":
		return &Float32Type{}

	case "int":
		return &IntType{}
	case "int32":
		return &Int32Type{}
	case "int16":
		return &Int16Type{}
	case "int8":
		return &Int8Type{}

	case "uint":
		return &UintType{}
	case "uint32":
		return &Uint32Type{}
	case "uint16":
		return &Uint16Type{}
	case "uint8":
		return &Uint8Type{}

	case "string":
		return &StringType{}
	case "bool":
		return &BoolType{}
	case "regexp":
		return &RegexpType{}
	}

	return nil
}
