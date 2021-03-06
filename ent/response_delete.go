// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/predicate"
	"github.com/vmkevv/duiztapi/ent/response"
)

// ResponseDelete is the builder for deleting a Response entity.
type ResponseDelete struct {
	config
	hooks    []Hook
	mutation *ResponseMutation
}

// Where adds a new predicate to the ResponseDelete builder.
func (rd *ResponseDelete) Where(ps ...predicate.Response) *ResponseDelete {
	rd.mutation.predicates = append(rd.mutation.predicates, ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ResponseDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResponseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ResponseDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ResponseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: response.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: response.FieldID,
			},
		},
	}
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// ResponseDeleteOne is the builder for deleting a single Response entity.
type ResponseDeleteOne struct {
	rd *ResponseDelete
}

// Exec executes the deletion query.
func (rdo *ResponseDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{response.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ResponseDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}
