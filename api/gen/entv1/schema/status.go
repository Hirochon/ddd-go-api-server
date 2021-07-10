package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(256).NotEmpty().Unique(),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return nil
}
