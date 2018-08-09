package db

// A DB stores all the information about a database, and the data inside
// it. A DB is created from a Schema.
type DB struct {
	data *Struct
}

// MakeDB makes a new database from a Schema, with all data set to its
// initial zero value.
func MakeDB(schema *Schema) *DB {
	fields := make(map[string]Type)

	for _, section := range schema.Sections {
		field := section.Field
		if field == nil {
			continue
		}

		fields[field.Name] = GetActualType(field.Type)
	}

	structType := &StructType{
		Name:   "db",
		Fields: fields,
	}

	return &DB{
		data: NewStruct(structType),
	}
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
	return item.GetField(clause.Ident)
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
