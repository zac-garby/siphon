package db

import "fmt"

// A DB stores all the information about a database, and the data inside
// it. A DB is created from a Schema.
type DB struct {
	data    *Struct
	structs map[string]*StructType
}

// JSON represents JSON data.
type JSON interface{}

// MakeDB makes a new database from a Schema, with all data set to its
// initial zero value.
func MakeDB(schema *Schema) (db *DB, err error) {
	structs := make(map[string]*StructType)

	for _, section := range schema.Sections {
		secStruct := section.Struct
		if secStruct == nil {
			continue
		}

		str := &StructType{
			Name:   secStruct.Name,
			Fields: make(map[string]Type),
		}
		structs[secStruct.Name] = str

		for _, field := range secStruct.Fields {
			ty := GetActualType(field.Type, structs)
			if ty == nil {
				return nil, fmt.Errorf("db init: type '%s' does not exist", field.Type.Ident)
			}
			str.Fields[field.Name] = ty
		}
	}

	fields := make(map[string]Type)

	for _, section := range schema.Sections {
		field := section.Field
		if field == nil {
			continue
		}

		ty := GetActualType(field.Type, structs)
		if ty == nil {
			return nil, fmt.Errorf("db init: type '%s' does not exist", field.Type.Ident)
		}
		fields[field.Name] = ty
	}

	structType := &StructType{
		Name:   "db",
		Fields: fields,
	}

	return &DB{
		data: NewStruct(structType),
	}, nil
}

// Query queries a database with a selector.
func (d *DB) Query(selector *Selector) (result Item, status string) {
	result = d.data

	for _, clause := range selector.Clauses {
		result, status = d.QuerySelectorClause(result, clause)
		if status != StatusOK {
			return
		}
	}

	return
}

// QuerySelectorClause queries an item
func (d *DB) QuerySelectorClause(item Item, clause *SelectorClause) (result Item, status string) {
	result, status = item.GetField(clause.Ident)
	if status != StatusOK {
		return nil, status
	}

	for _, filter := range clause.Filters {
		if cmp := filter.Comparison; cmp != nil {
			field := cmp.Ident

			comparison, ok := stringToComparison(cmp.Comparison)
			if !ok {
				return nil, StatusNOOP
			}

			other := selectorLiteralToItem(cmp.Literal)

			result, status = result.Filter(field, comparison, other)
			if status != StatusOK {
				return nil, status
			}
		} else if idx := filter.Index; idx != nil {
			key := selectorLiteralToItem(idx)

			result, status = result.GetKey(key)
			if status != StatusOK {
				return nil, status
			}
		}
	}

	return result, StatusOK
}

func selectorLiteralToItem(lit *SelectorLiteral) Item {
	if num := lit.Number; num != nil {
		return NewFloat(*num)
	} else if str := lit.String; str != nil {
		return NewString(*str)
	} else if reg := lit.Regexp; reg != nil {
		return NewRegexp(*reg)
	}
	return nil
}

// QueryString queries a database, parsing the string as
// as selector first.
func (d *DB) QueryString(str string) (result Item, status string) {
	selector := &Selector{}
	if err := SelectorParser.ParseString(str, selector); err != nil {
		return nil, err.Error()
	}

	return d.Query(selector)
}

// Set sets the value of the selected item(s) to the given value.
func (d *DB) Set(selector *Selector, to JSON) (status string) {
	_, status = d.Query(selector)
	if status != StatusOK {
		return status
	}

	return StatusOK
}
