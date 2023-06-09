// Code generated by ent, DO NOT EDIT.

package access

import (
	"time"
)

const (
	// Label holds the string label denoting the access type in the database.
	Label = "access"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCheckIn holds the string denoting the check_in field in the database.
	FieldCheckIn = "check_in"
	// FieldCheckOut holds the string denoting the check_out field in the database.
	FieldCheckOut = "check_out"
	// Table holds the table name of the access in the database.
	Table = "accesses"
)

// Columns holds all SQL columns for access fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldCheckIn,
	FieldCheckOut,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCheckIn holds the default value on creation for the "check_in" field.
	DefaultCheckIn time.Time
)
