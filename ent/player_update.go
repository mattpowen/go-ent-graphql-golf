// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golfapp/ent/player"
	"golfapp/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerUpdate is the builder for updating Player entities.
type PlayerUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerMutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (pu *PlayerUpdate) Where(ps ...predicate.Player) *PlayerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// Mutation returns the PlayerMutation object of the builder.
func (pu *PlayerUpdate) Mutation() *PlayerMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlayerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlayerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlayerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlayerUpdateOne is the builder for updating a single Player entity.
type PlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerMutation
}

// Mutation returns the PlayerMutation object of the builder.
func (puo *PlayerUpdateOne) Mutation() *PlayerMutation {
	return puo.mutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (puo *PlayerUpdateOne) Where(ps ...predicate.Player) *PlayerUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlayerUpdateOne) Select(field string, fields ...string) *PlayerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Player entity.
func (puo *PlayerUpdateOne) Save(ctx context.Context) (*Player, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlayerUpdateOne) SaveX(ctx context.Context) *Player {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlayerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlayerUpdateOne) sqlSave(ctx context.Context) (_node *Player, err error) {
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Player.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for _, f := range fields {
			if !player.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != player.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Player{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
