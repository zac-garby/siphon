package db

// A DB stores all the information about a database, and the data inside
// it. A DB is created from a Schema.
type DB struct {
	data    *Struct
	structs map[string]*StructType
}

// MakeDB makes a new database from a Schema, with all data set to its
// initial zero value.
func MakeDB(schema *Schema) (db *DB, status string) {
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
				return nil, StatusNoType
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
			return nil, StatusNoType
		}
		fields[field.Name] = ty
	}

	structType := &StructType{
		Name:   "db",
		Fields: fields,
	}

	return &DB{
		data: NewStruct(structType),
	}, StatusOK
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
