package db

// A DB stores all the information about a database, and the data inside
// it. A DB is created from a Schema.
type DB struct {
	data *Hashmap
}

// MakeDB makes a new database from a Schema, with all data set to its
// initial zero value.
func MakeDB(schema *Schema) *DB {
	db := &DB{
		data: NewHashmap(),
	}

	for _, section := range schema.Sections {
		field := section.Field
		if field == nil {
			continue
		}

		db.data.SetKey(NewString(field.Name), MakeZeroValue(field.Type))
	}

	return db
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
	hm, ok := item.(*Hashmap)
	if !ok {
		return nil, StatusNOOP
	}

	return hm.GetKey(NewString(clause.Ident))
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
