// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/DahuiSheng/gobackend/ent/access"
)

// Access is the model entity for the Access schema.
type Access struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Access) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case access.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Access", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Access fields.
func (a *Access) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case access.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Access.
// Note that you need to call Access.Unwrap() before calling this method if this Access
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Access) Update() *AccessUpdateOne {
	return NewAccessClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Access entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Access) Unwrap() *Access {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Access is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Access) String() string {
	var builder strings.Builder
	builder.WriteString("Access(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Accesses is a parsable slice of Access.
type Accesses []*Access