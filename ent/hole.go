// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golfapp/ent/hole"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Hole is the model entity for the Hole schema.
type Hole struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HoleNumber holds the value of the "hole_number" field.
	HoleNumber int `json:"hole_number,omitempty"`
	// HoleName holds the value of the "hole_name" field.
	HoleName     string `json:"hole_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hole.FieldID, hole.FieldHoleNumber:
			values[i] = new(sql.NullInt64)
		case hole.FieldHoleName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hole fields.
func (h *Hole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case hole.FieldHoleNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hole_number", values[i])
			} else if value.Valid {
				h.HoleNumber = int(value.Int64)
			}
		case hole.FieldHoleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hole_name", values[i])
			} else if value.Valid {
				h.HoleName = value.String
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Hole.
// This includes values selected through modifiers, order, etc.
func (h *Hole) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// Update returns a builder for updating this Hole.
// Note that you need to call Hole.Unwrap() before calling this method if this Hole
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hole) Update() *HoleUpdateOne {
	return NewHoleClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hole) Unwrap() *Hole {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hole is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hole) String() string {
	var builder strings.Builder
	builder.WriteString("Hole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("hole_number=")
	builder.WriteString(fmt.Sprintf("%v", h.HoleNumber))
	builder.WriteString(", ")
	builder.WriteString("hole_name=")
	builder.WriteString(h.HoleName)
	builder.WriteByte(')')
	return builder.String()
}

// Holes is a parsable slice of Hole.
type Holes []*Hole
