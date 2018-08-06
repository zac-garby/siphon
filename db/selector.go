package db

import (
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

var selectorLexer = lexer.Must(lexer.Regexp(`(?m)(\s+)` +
	`|(?P<Ident>[\p{L}\p{M}_-][\p{L}\p{M}\d_-]*)` +
	`|(?P<Number>\d+(?:\.\d+)?)` +
	`|(?P<String>"(?:\\.|[^"])*"|'(?:\\.|[^'])*')` +
	`|(?P<Regexp>/(?:\\.|[^/])+/)` +
	`|(?P<Comparison>(?:=|!=|>|>=|<|<=|~))` +
	`|(?P<Punctuation>[\.\[\]])`,
))

// SelectorParser parses query selectors.
var SelectorParser *participle.Parser

func init() {
	parser, err := participle.Build(
		&Selector{},

		participle.Lexer(selectorLexer),
		participle.Unquote(selectorLexer, "String"),

		participle.Map(func(token lexer.Token) lexer.Token {
			if token.Type == lexer.DefaultDefinition.Symbols()["Regexp"] {
				token.Value = strings.Trim(token.Value, "/")
			}
			return token
		}),
	)

	if err != nil {
		panic(err)
	}

	SelectorParser = parser
}

// A Selector is used to query the database.
type Selector struct {
	Clauses []*SelectorClause `@@ { "." @@ }`
}

// A SelectorClause is one part of a selector, for example "users[3]".
type SelectorClause struct {
	Ident   string            `@Ident`
	Filters []*SelectorFilter `{ "[" @@ "]" }`
}

// A SelectorFilter filters a clause based on either an index or a comparison.
type SelectorFilter struct {
	Comparison *SelectorFilterComparison `  @@`
	Index      *float64                  `| @Number`
}

// A SelectorFilterComparison filters a clause based on whether an attribute of
// a value is a certain thing.
type SelectorFilterComparison struct {
	Ident      string           `@Ident`
	Comparison string           `@Comparison`
	Literal    *SelectorLiteral `@@`
}

// A SelectorLiteral is a literal value, like a string, number, or regexp.
type SelectorLiteral struct {
	String *string  `  @String`
	Number *float64 `| @Number`
	Regexp *string  `| @Regexp`
}
