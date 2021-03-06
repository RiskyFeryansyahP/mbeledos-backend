// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/confus1on/mbeledos/ent/predicate"
	"github.com/confus1on/mbeledos/ent/transaction"
	"github.com/facebookincubator/ent/dialect/sql"
)

// TransactionDelete is the builder for deleting a Transaction entity.
type TransactionDelete struct {
	config
	predicates []predicate.Transaction
}

// Where adds a new predicate to the delete builder.
func (td *TransactionDelete) Where(ps ...predicate.Transaction) *TransactionDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TransactionDelete) Exec(ctx context.Context) (int, error) {
	return td.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TransactionDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TransactionDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(td.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(transaction.Table))
	for _, p := range td.predicates {
		p(selector)
	}
	query, args := builder.Delete(transaction.Table).FromSelect(selector).Query()
	if err := td.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// TransactionDeleteOne is the builder for deleting a single Transaction entity.
type TransactionDeleteOne struct {
	td *TransactionDelete
}

// Exec executes the deletion query.
func (tdo *TransactionDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{transaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TransactionDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
