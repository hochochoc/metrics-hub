// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"metric-hub/ent/predicate"
	"metric-hub/ent/report"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReportUpdate is the builder for updating Report entities.
type ReportUpdate struct {
	config
	hooks    []Hook
	mutation *ReportMutation
}

// Where appends a list predicates to the ReportUpdate builder.
func (ru *ReportUpdate) Where(ps ...predicate.Report) *ReportUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetServerName sets the "server_name" field.
func (ru *ReportUpdate) SetServerName(s string) *ReportUpdate {
	ru.mutation.SetServerName(s)
	return ru
}

// SetStartTime sets the "start_time" field.
func (ru *ReportUpdate) SetStartTime(t time.Time) *ReportUpdate {
	ru.mutation.SetStartTime(t)
	return ru
}

// SetEndTime sets the "end_time" field.
func (ru *ReportUpdate) SetEndTime(t time.Time) *ReportUpdate {
	ru.mutation.SetEndTime(t)
	return ru
}

// SetDurationInSeconds sets the "duration_in_seconds" field.
func (ru *ReportUpdate) SetDurationInSeconds(f float64) *ReportUpdate {
	ru.mutation.ResetDurationInSeconds()
	ru.mutation.SetDurationInSeconds(f)
	return ru
}

// AddDurationInSeconds adds f to the "duration_in_seconds" field.
func (ru *ReportUpdate) AddDurationInSeconds(f float64) *ReportUpdate {
	ru.mutation.AddDurationInSeconds(f)
	return ru
}

// Mutation returns the ReportMutation object of the builder.
func (ru *ReportUpdate) Mutation() *ReportMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReportUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReportUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReportUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReportUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   report.Table,
			Columns: report.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: report.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ServerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: report.FieldServerName,
		})
	}
	if value, ok := ru.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: report.FieldStartTime,
		})
	}
	if value, ok := ru.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: report.FieldEndTime,
		})
	}
	if value, ok := ru.mutation.DurationInSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: report.FieldDurationInSeconds,
		})
	}
	if value, ok := ru.mutation.AddedDurationInSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: report.FieldDurationInSeconds,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{report.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReportUpdateOne is the builder for updating a single Report entity.
type ReportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReportMutation
}

// SetServerName sets the "server_name" field.
func (ruo *ReportUpdateOne) SetServerName(s string) *ReportUpdateOne {
	ruo.mutation.SetServerName(s)
	return ruo
}

// SetStartTime sets the "start_time" field.
func (ruo *ReportUpdateOne) SetStartTime(t time.Time) *ReportUpdateOne {
	ruo.mutation.SetStartTime(t)
	return ruo
}

// SetEndTime sets the "end_time" field.
func (ruo *ReportUpdateOne) SetEndTime(t time.Time) *ReportUpdateOne {
	ruo.mutation.SetEndTime(t)
	return ruo
}

// SetDurationInSeconds sets the "duration_in_seconds" field.
func (ruo *ReportUpdateOne) SetDurationInSeconds(f float64) *ReportUpdateOne {
	ruo.mutation.ResetDurationInSeconds()
	ruo.mutation.SetDurationInSeconds(f)
	return ruo
}

// AddDurationInSeconds adds f to the "duration_in_seconds" field.
func (ruo *ReportUpdateOne) AddDurationInSeconds(f float64) *ReportUpdateOne {
	ruo.mutation.AddDurationInSeconds(f)
	return ruo
}

// Mutation returns the ReportMutation object of the builder.
func (ruo *ReportUpdateOne) Mutation() *ReportMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReportUpdateOne) Select(field string, fields ...string) *ReportUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Report entity.
func (ruo *ReportUpdateOne) Save(ctx context.Context) (*Report, error) {
	var (
		err  error
		node *Report
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReportUpdateOne) SaveX(ctx context.Context) *Report {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReportUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReportUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReportUpdateOne) sqlSave(ctx context.Context) (_node *Report, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   report.Table,
			Columns: report.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: report.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Report.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, report.FieldID)
		for _, f := range fields {
			if !report.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != report.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ServerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: report.FieldServerName,
		})
	}
	if value, ok := ruo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: report.FieldStartTime,
		})
	}
	if value, ok := ruo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: report.FieldEndTime,
		})
	}
	if value, ok := ruo.mutation.DurationInSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: report.FieldDurationInSeconds,
		})
	}
	if value, ok := ruo.mutation.AddedDurationInSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: report.FieldDurationInSeconds,
		})
	}
	_node = &Report{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{report.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
