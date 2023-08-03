// Code generated by ent, DO NOT EDIT.

package tee

import (
	"golfapp/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tee {
	return predicate.Tee(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tee {
	return predicate.Tee(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tee {
	return predicate.Tee(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tee {
	return predicate.Tee(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tee {
	return predicate.Tee(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tee {
	return predicate.Tee(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tee {
	return predicate.Tee(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tee {
	return predicate.Tee(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tee {
	return predicate.Tee(sql.FieldLTE(FieldID, id))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.Tee {
	return predicate.Tee(sql.FieldEQ(FieldColor, v))
}

// HoleName applies equality check predicate on the "hole_name" field. It's identical to HoleNameEQ.
func HoleName(v int) predicate.Tee {
	return predicate.Tee(sql.FieldEQ(FieldHoleName, v))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.Tee {
	return predicate.Tee(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.Tee {
	return predicate.Tee(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.Tee {
	return predicate.Tee(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.Tee {
	return predicate.Tee(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.Tee {
	return predicate.Tee(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.Tee {
	return predicate.Tee(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.Tee {
	return predicate.Tee(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.Tee {
	return predicate.Tee(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.Tee {
	return predicate.Tee(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.Tee {
	return predicate.Tee(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.Tee {
	return predicate.Tee(sql.FieldHasSuffix(FieldColor, v))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.Tee {
	return predicate.Tee(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.Tee {
	return predicate.Tee(sql.FieldContainsFold(FieldColor, v))
}

// HoleNameEQ applies the EQ predicate on the "hole_name" field.
func HoleNameEQ(v int) predicate.Tee {
	return predicate.Tee(sql.FieldEQ(FieldHoleName, v))
}

// HoleNameNEQ applies the NEQ predicate on the "hole_name" field.
func HoleNameNEQ(v int) predicate.Tee {
	return predicate.Tee(sql.FieldNEQ(FieldHoleName, v))
}

// HoleNameIn applies the In predicate on the "hole_name" field.
func HoleNameIn(vs ...int) predicate.Tee {
	return predicate.Tee(sql.FieldIn(FieldHoleName, vs...))
}

// HoleNameNotIn applies the NotIn predicate on the "hole_name" field.
func HoleNameNotIn(vs ...int) predicate.Tee {
	return predicate.Tee(sql.FieldNotIn(FieldHoleName, vs...))
}

// HoleNameGT applies the GT predicate on the "hole_name" field.
func HoleNameGT(v int) predicate.Tee {
	return predicate.Tee(sql.FieldGT(FieldHoleName, v))
}

// HoleNameGTE applies the GTE predicate on the "hole_name" field.
func HoleNameGTE(v int) predicate.Tee {
	return predicate.Tee(sql.FieldGTE(FieldHoleName, v))
}

// HoleNameLT applies the LT predicate on the "hole_name" field.
func HoleNameLT(v int) predicate.Tee {
	return predicate.Tee(sql.FieldLT(FieldHoleName, v))
}

// HoleNameLTE applies the LTE predicate on the "hole_name" field.
func HoleNameLTE(v int) predicate.Tee {
	return predicate.Tee(sql.FieldLTE(FieldHoleName, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tee) predicate.Tee {
	return predicate.Tee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tee) predicate.Tee {
	return predicate.Tee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tee) predicate.Tee {
	return predicate.Tee(func(s *sql.Selector) {
		p(s.Not())
	})
}
