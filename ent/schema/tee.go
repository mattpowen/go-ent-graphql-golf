package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tee holds the schema definition for the Tee entity.
type Tee struct {
	ent.Schema
}

// Fields of the Tee.
func (Tee) Fields() []ent.Field {
	return []ent.Field{
		field.String("color"),
		field.Int("hole_name"),
	}
}

// Edges of the Tee.
func (Tee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hole", Hole.Type).
			Ref("Tees").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
