// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/i18n"
	"github.com/vmkevv/duiztapi/ent/predicate"
	"github.com/vmkevv/duiztapi/ent/quiz"
	"github.com/vmkevv/duiztapi/ent/quizlangs"
)

// QuizLangsUpdate is the builder for updating QuizLangs entities.
type QuizLangsUpdate struct {
	config
	hooks    []Hook
	mutation *QuizLangsMutation
}

// Where adds a new predicate for the QuizLangsUpdate builder.
func (qlu *QuizLangsUpdate) Where(ps ...predicate.QuizLangs) *QuizLangsUpdate {
	qlu.mutation.predicates = append(qlu.mutation.predicates, ps...)
	return qlu
}

// SetName sets the "name" field.
func (qlu *QuizLangsUpdate) SetName(s string) *QuizLangsUpdate {
	qlu.mutation.SetName(s)
	return qlu
}

// SetDescription sets the "description" field.
func (qlu *QuizLangsUpdate) SetDescription(s string) *QuizLangsUpdate {
	qlu.mutation.SetDescription(s)
	return qlu
}

// SetQuizID sets the "quiz" edge to the Quiz entity by ID.
func (qlu *QuizLangsUpdate) SetQuizID(id int) *QuizLangsUpdate {
	qlu.mutation.SetQuizID(id)
	return qlu
}

// SetNillableQuizID sets the "quiz" edge to the Quiz entity by ID if the given value is not nil.
func (qlu *QuizLangsUpdate) SetNillableQuizID(id *int) *QuizLangsUpdate {
	if id != nil {
		qlu = qlu.SetQuizID(*id)
	}
	return qlu
}

// SetQuiz sets the "quiz" edge to the Quiz entity.
func (qlu *QuizLangsUpdate) SetQuiz(q *Quiz) *QuizLangsUpdate {
	return qlu.SetQuizID(q.ID)
}

// SetI18nID sets the "i18n" edge to the I18n entity by ID.
func (qlu *QuizLangsUpdate) SetI18nID(id int) *QuizLangsUpdate {
	qlu.mutation.SetI18nID(id)
	return qlu
}

// SetNillableI18nID sets the "i18n" edge to the I18n entity by ID if the given value is not nil.
func (qlu *QuizLangsUpdate) SetNillableI18nID(id *int) *QuizLangsUpdate {
	if id != nil {
		qlu = qlu.SetI18nID(*id)
	}
	return qlu
}

// SetI18n sets the "i18n" edge to the I18n entity.
func (qlu *QuizLangsUpdate) SetI18n(i *I18n) *QuizLangsUpdate {
	return qlu.SetI18nID(i.ID)
}

// Mutation returns the QuizLangsMutation object of the builder.
func (qlu *QuizLangsUpdate) Mutation() *QuizLangsMutation {
	return qlu.mutation
}

// ClearQuiz clears the "quiz" edge to the Quiz entity.
func (qlu *QuizLangsUpdate) ClearQuiz() *QuizLangsUpdate {
	qlu.mutation.ClearQuiz()
	return qlu
}

// ClearI18n clears the "i18n" edge to the I18n entity.
func (qlu *QuizLangsUpdate) ClearI18n() *QuizLangsUpdate {
	qlu.mutation.ClearI18n()
	return qlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qlu *QuizLangsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qlu.hooks) == 0 {
		affected, err = qlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuizLangsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qlu.mutation = mutation
			affected, err = qlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qlu.hooks) - 1; i >= 0; i-- {
			mut = qlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qlu *QuizLangsUpdate) SaveX(ctx context.Context) int {
	affected, err := qlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qlu *QuizLangsUpdate) Exec(ctx context.Context) error {
	_, err := qlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qlu *QuizLangsUpdate) ExecX(ctx context.Context) {
	if err := qlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qlu *QuizLangsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   quizlangs.Table,
			Columns: quizlangs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: quizlangs.FieldID,
			},
		},
	}
	if ps := qlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qlu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quizlangs.FieldName,
		})
	}
	if value, ok := qlu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quizlangs.FieldDescription,
		})
	}
	if qlu.mutation.QuizCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.QuizTable,
			Columns: []string{quizlangs.QuizColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quiz.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qlu.mutation.QuizIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.QuizTable,
			Columns: []string{quizlangs.QuizColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quiz.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qlu.mutation.I18nCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.I18nTable,
			Columns: []string{quizlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qlu.mutation.I18nIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.I18nTable,
			Columns: []string{quizlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{quizlangs.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// QuizLangsUpdateOne is the builder for updating a single QuizLangs entity.
type QuizLangsUpdateOne struct {
	config
	hooks    []Hook
	mutation *QuizLangsMutation
}

// SetName sets the "name" field.
func (qluo *QuizLangsUpdateOne) SetName(s string) *QuizLangsUpdateOne {
	qluo.mutation.SetName(s)
	return qluo
}

// SetDescription sets the "description" field.
func (qluo *QuizLangsUpdateOne) SetDescription(s string) *QuizLangsUpdateOne {
	qluo.mutation.SetDescription(s)
	return qluo
}

// SetQuizID sets the "quiz" edge to the Quiz entity by ID.
func (qluo *QuizLangsUpdateOne) SetQuizID(id int) *QuizLangsUpdateOne {
	qluo.mutation.SetQuizID(id)
	return qluo
}

// SetNillableQuizID sets the "quiz" edge to the Quiz entity by ID if the given value is not nil.
func (qluo *QuizLangsUpdateOne) SetNillableQuizID(id *int) *QuizLangsUpdateOne {
	if id != nil {
		qluo = qluo.SetQuizID(*id)
	}
	return qluo
}

// SetQuiz sets the "quiz" edge to the Quiz entity.
func (qluo *QuizLangsUpdateOne) SetQuiz(q *Quiz) *QuizLangsUpdateOne {
	return qluo.SetQuizID(q.ID)
}

// SetI18nID sets the "i18n" edge to the I18n entity by ID.
func (qluo *QuizLangsUpdateOne) SetI18nID(id int) *QuizLangsUpdateOne {
	qluo.mutation.SetI18nID(id)
	return qluo
}

// SetNillableI18nID sets the "i18n" edge to the I18n entity by ID if the given value is not nil.
func (qluo *QuizLangsUpdateOne) SetNillableI18nID(id *int) *QuizLangsUpdateOne {
	if id != nil {
		qluo = qluo.SetI18nID(*id)
	}
	return qluo
}

// SetI18n sets the "i18n" edge to the I18n entity.
func (qluo *QuizLangsUpdateOne) SetI18n(i *I18n) *QuizLangsUpdateOne {
	return qluo.SetI18nID(i.ID)
}

// Mutation returns the QuizLangsMutation object of the builder.
func (qluo *QuizLangsUpdateOne) Mutation() *QuizLangsMutation {
	return qluo.mutation
}

// ClearQuiz clears the "quiz" edge to the Quiz entity.
func (qluo *QuizLangsUpdateOne) ClearQuiz() *QuizLangsUpdateOne {
	qluo.mutation.ClearQuiz()
	return qluo
}

// ClearI18n clears the "i18n" edge to the I18n entity.
func (qluo *QuizLangsUpdateOne) ClearI18n() *QuizLangsUpdateOne {
	qluo.mutation.ClearI18n()
	return qluo
}

// Save executes the query and returns the updated QuizLangs entity.
func (qluo *QuizLangsUpdateOne) Save(ctx context.Context) (*QuizLangs, error) {
	var (
		err  error
		node *QuizLangs
	)
	if len(qluo.hooks) == 0 {
		node, err = qluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuizLangsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qluo.mutation = mutation
			node, err = qluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qluo.hooks) - 1; i >= 0; i-- {
			mut = qluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (qluo *QuizLangsUpdateOne) SaveX(ctx context.Context) *QuizLangs {
	node, err := qluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (qluo *QuizLangsUpdateOne) Exec(ctx context.Context) error {
	_, err := qluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qluo *QuizLangsUpdateOne) ExecX(ctx context.Context) {
	if err := qluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qluo *QuizLangsUpdateOne) sqlSave(ctx context.Context) (_node *QuizLangs, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   quizlangs.Table,
			Columns: quizlangs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: quizlangs.FieldID,
			},
		},
	}
	id, ok := qluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing QuizLangs.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := qluo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quizlangs.FieldName,
		})
	}
	if value, ok := qluo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quizlangs.FieldDescription,
		})
	}
	if qluo.mutation.QuizCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.QuizTable,
			Columns: []string{quizlangs.QuizColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quiz.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qluo.mutation.QuizIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.QuizTable,
			Columns: []string{quizlangs.QuizColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quiz.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qluo.mutation.I18nCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.I18nTable,
			Columns: []string{quizlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qluo.mutation.I18nIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   quizlangs.I18nTable,
			Columns: []string{quizlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &QuizLangs{config: qluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, qluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{quizlangs.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}