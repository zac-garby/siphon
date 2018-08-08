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
}
