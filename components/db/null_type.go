package db

// Nullable types are described in this file. They are used to populate SQL models by standard way.
// Use case:
//
// type Something struct {
//     Name string `db: "name"`
// }
//
// Field 'Name' has type 'string', but how to populate it when field 'name' in database has 'null' value?
// We cannot just set empty string, because we need to distinguish empty strings and null values.
// That's why 'null types' are introduced: they have additional 'Valid' field which is boolean.
// Basically, these types are defined in "database/sql" package, but here some extensions are introduced.

import (
	"database/sql"
	"encoding/json"
)

// NullString represent "nullable string" value
type NullString sql.NullString

// MarshalJSON serializes "nullable string" into json string
func (n NullString) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.String)
	}

	return []byte("null"), nil
}

// UnmarshalJSON parses data into "nullable string"
func (n *NullString) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &n.String)
	n.Valid = err == nil

	return err
}

// NullInt64 represent "nullable int64" value
type NullInt64 struct {
	sql.NullInt64
}

// NewNullInt64 create valid NullInt64 object
func NewNullInt64(Int64 int64) NullInt64 {
	return NullInt64{
		sql.NullInt64{
			Valid: true,
			Int64: Int64,
		},
	}
}

// MarshalJSON serializes "nullable int64" into json string
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int64)
	}

	return []byte("null"), nil
}

// UnmarshalJSON parses data into "nullable int64"
func (n *NullInt64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &n.Int64)
	n.Valid = err == nil

	return err
}

// NullFloat64 represent "nullable float64" value
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON serializes "nullable float64" into json string
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Float64)
	}

	return []byte("null"), nil
}

// UnmarshalJSON parses data into "nullable float64"
func (n *NullFloat64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &n.Float64)
	n.Valid = err == nil

	return err
}

// NullBool represent "nullable bool" value
type NullBool struct {
	sql.NullBool
}

// MarshalJSON serializes "nullable bool" into json string
func (n NullBool) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Bool)
	}

	return []byte("null"), nil
}

// UnmarshalJSON parses data into "nullable bool"
func (n *NullBool) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &n.Bool)
	n.Valid = err == nil

	return err
}
