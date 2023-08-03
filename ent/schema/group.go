package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("group_name").
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.Time("date_of_group"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
