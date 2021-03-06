// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/answerlangs"
	"github.com/vmkevv/duiztapi/ent/predicate"
)

// AnswerLangsDelete is the builder for deleting a AnswerLangs entity.
type AnswerLangsDelete struct {
	config
	hooks    []Hook
	mutation *AnswerLangsMutation
}

// Where adds a new predicate to the AnswerLangsDelete builder.
func (ald *AnswerLangsDelete) Where(ps ...predicate.AnswerLangs) *AnswerLangsDelete {
	ald.mutation.predicates = append(ald.mutation.predicates, ps...)
	return ald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ald *AnswerLangsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ald.hooks) == 0 {
		affected, err = ald.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnswerLangsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ald.mutation = mutation
			affected, err = ald.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ald.hooks) - 1; i >= 0; i-- {
			mut = ald.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ald.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ald *AnswerLangsDelete) ExecX(ctx context.Context) int {
	n, err := ald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ald *AnswerLangsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: answerlangs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: answerlangs.FieldID,
			},
		},
	}
	if ps := ald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ald.driver, _spec)
}

// AnswerLangsDeleteOne is the builder for deleting a single AnswerLangs entity.
type AnswerLangsDeleteOne struct {
	ald *AnswerLangsDelete
}

// Exec executes the deletion query.
func (aldo *AnswerLangsDeleteOne) Exec(ctx context.Context) error {
	n, err := aldo.ald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{answerlangs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aldo *AnswerLangsDeleteOne) ExecX(ctx context.Context) {
	aldo.ald.ExecX(ctx)
}
