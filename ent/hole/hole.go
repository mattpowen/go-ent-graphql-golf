// Code generated by ent, DO NOT EDIT.

package hole

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the hole type in the database.
	Label = "hole"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHoleNumber holds the string denoting the hole_number field in the database.
	FieldHoleNumber = "hole_number"
	// FieldHoleName holds the string denoting the hole_name field in the database.
	FieldHoleName = "hole_name"
	// Table holds the table name of the hole in the database.
	Table = "holes"
)

// Columns holds all SQL columns for hole fields.
var Columns = []string{
	FieldID,
	FieldHoleNumber,
	FieldHoleName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Hole queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHoleNumber orders the results by the hole_number field.
func ByHoleNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHoleNumber, opts...).ToFunc()
}

// ByHoleName orders the results by the hole_name field.
func ByHoleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHoleName, opts...).ToFunc()
}
