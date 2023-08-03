package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("course_name"),
		field.String("course_location"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Holes", Hole.Type),
		edge.From("Round", Round.Type).
			Ref("Rounds").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
