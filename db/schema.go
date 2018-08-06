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
