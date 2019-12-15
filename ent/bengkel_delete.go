// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/confus1on/mbeledos/ent/bengkel"
	"github.com/confus1on/mbeledos/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// BengkelDelete is the builder for deleting a Bengkel entity.
type BengkelDelete struct {
	config
	predicates []predicate.Bengkel
}

// Where adds a new predicate to the delete builder.
func (bd *BengkelDelete) Where(ps ...predicate.Bengkel) *BengkelDelete {
	bd.predicates = append(bd.predicates, ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BengkelDelete) Exec(ctx context.Context) (int, error) {
	return bd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BengkelDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BengkelDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(bd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(bengkel.Table))
	for _, p := range bd.predicates {
		p(selector)
	}
	query, args := builder.Delete(bengkel.Table).FromSelect(selector).Query()
	if err := bd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// BengkelDeleteOne is the builder for deleting a single Bengkel entity.
type BengkelDeleteOne struct {
	bd *BengkelDelete
}

// Exec executes the deletion query.
func (bdo *BengkelDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{bengkel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BengkelDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
