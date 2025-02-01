package utils

import "database/sql"

// GetNullInt returns a sql.NullInt32 object from an int32 pointer
func GetNullInt(i *int32) sql.NullInt32 {
	if i != nil {
		return sql.NullInt32{Int32: *i, Valid: true}
	} else {
		return sql.NullInt32{Valid: false}
	}
}

// GetNullString returns a sql.NullString object from a string
func GetNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{String: s, Valid: true}
	} else {
		return sql.NullString{Valid: false}
	}
}
