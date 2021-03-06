// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"api/gen/entv1/status"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Status is the model entity for the Status schema.
type Status struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Status) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case status.FieldID:
			values[i] = new(sql.NullInt64)
		case status.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Status", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Status fields.
func (s *Status) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case status.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case status.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Status.
// Note that you need to call Status.Unwrap() before calling this method if this Status
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Status) Update() *StatusUpdateOne {
	return (&StatusClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Status entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Status) Unwrap() *Status {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("entv1: Status is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Status) String() string {
	var builder strings.Builder
	builder.WriteString("Status(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteByte(')')
	return builder.String()
}

// StatusSlice is a parsable slice of Status.
type StatusSlice []*Status

func (s StatusSlice) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
