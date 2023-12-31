package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_of_round"),
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Course", Course.Type),
	}
}
