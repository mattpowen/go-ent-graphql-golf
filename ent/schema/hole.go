package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hole holds the schema definition for the Hole entity.
type Hole struct {
	ent.Schema
}

// Fields of the Hole.
func (Hole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("hole_number"),
		field.String("hole_name"),
	}
}

// Edges of the Hole.
func (Hole) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("Course", Course.Type).
			Ref("holes").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),

		edge.To("Tee", Hole.Type),
	}
}
