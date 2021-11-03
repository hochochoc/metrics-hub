// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"metric-hub/ent/report"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReportCreate is the builder for creating a Report entity.
type ReportCreate struct {
	config
	mutation *ReportMutation
	hooks    []Hook
}

// SetServerName sets the "server_name" field.
func (rc *ReportCreate) SetServerName(s string) *ReportCreate {
	rc.mutation.SetServerName(s)
	return rc
}

// SetStartTime sets the "start_time" field.
func (rc *ReportCreate) SetStartTime(t time.Time) *ReportCreate {
	rc.mutation.SetStartTime(t)
	return rc
}

// SetEndTime sets the "end_time" field.
func (rc *ReportCreate) SetEndTime(t time.Time) *ReportCreate {
	rc.mutation.SetEndTime(t)
	return rc
}

// SetDurationInSeconds sets the "duration_in_seconds" field.
func (rc *ReportCreate) SetDurationInSeconds(f float64) *ReportCreate {
	rc.mutation.SetDurationInSeconds(f)
	return rc
}

// Mutation returns the ReportMutation object of the builder.
func (rc *ReportCreate) Mutation() *ReportMutation {
	return rc.mutation
}

// Save creates the Report in the database.
func (rc *ReportCreate) Save(ctx context.Context) (*Report, error) {
	var (
		err  error
		node *Report
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReportCreate) SaveX(ctx context.Context) *Report {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReportCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReportCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReportCreate) check() error {
	if _, ok := rc.mutation.ServerName(); !ok {
		return &ValidationError{Name: "server_name", err: errors.New(`ent: missing required field "server_name"`)}
	}
	if _, ok := rc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "start_time"`)}
	}
	if _, ok := rc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "end_time"`)}
	}
	if _, ok := rc.mutation.DurationInSeconds(); !ok {
		return &ValidationError{Name: "duration_in_seconds", err: errors.New(`ent: missing required field "duration_in_seconds"`)}
	}
	return nil
}

func (rc *ReportCreate) sqlSave(ctx context.Context) (*Report, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *ReportCreate) createSpec() (*Report, *sqlgraph.CreateSpec) {
	var (
		_node = &Report{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: report.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: report.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.ServerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: report.FieldServerName,
		})
		_node.ServerName = value
	}
	if value, ok := rc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: report.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := rc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: report.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := rc.mutation.DurationInSeconds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: report.FieldDurationInSeconds,
		})
		_node.DurationInSeconds = value
	}
	return _node, _spec
}

// ReportCreateBulk is the builder for creating many Report entities in bulk.
type ReportCreateBulk struct {
	config
	builders []*ReportCreate
}

// Save creates the Report entities in the database.
func (rcb *ReportCreateBulk) Save(ctx context.Context) ([]*Report, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Report, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReportMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReportCreateBulk) SaveX(ctx context.Context) []*Report {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReportCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReportCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
