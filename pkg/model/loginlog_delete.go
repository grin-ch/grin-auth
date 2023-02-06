// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/grin-ch/grin-auth/pkg/model/loginlog"
	"github.com/grin-ch/grin-auth/pkg/model/predicate"
)

// LoginLogDelete is the builder for deleting a LoginLog entity.
type LoginLogDelete struct {
	config
	hooks    []Hook
	mutation *LoginLogMutation
}

// Where appends a list predicates to the LoginLogDelete builder.
func (lld *LoginLogDelete) Where(ps ...predicate.LoginLog) *LoginLogDelete {
	lld.mutation.Where(ps...)
	return lld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lld *LoginLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LoginLogMutation](ctx, lld.sqlExec, lld.mutation, lld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lld *LoginLogDelete) ExecX(ctx context.Context) int {
	n, err := lld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lld *LoginLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: loginlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: loginlog.FieldID,
			},
		},
	}
	if ps := lld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lld.mutation.done = true
	return affected, err
}

// LoginLogDeleteOne is the builder for deleting a single LoginLog entity.
type LoginLogDeleteOne struct {
	lld *LoginLogDelete
}

// Where appends a list predicates to the LoginLogDelete builder.
func (lldo *LoginLogDeleteOne) Where(ps ...predicate.LoginLog) *LoginLogDeleteOne {
	lldo.lld.mutation.Where(ps...)
	return lldo
}

// Exec executes the deletion query.
func (lldo *LoginLogDeleteOne) Exec(ctx context.Context) error {
	n, err := lldo.lld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loginlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lldo *LoginLogDeleteOne) ExecX(ctx context.Context) {
	if err := lldo.Exec(ctx); err != nil {
		panic(err)
	}
}
