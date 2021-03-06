// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/confus1on/mbeledos/ent/predicate"
	"github.com/confus1on/mbeledos/ent/transaction"
	"github.com/facebookincubator/ent/dialect/sql"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	orderphone  *string
	namabengkel *string
	predicates  []predicate.Transaction
}

// Where adds a new predicate for the builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetOrderphone sets the orderphone field.
func (tu *TransactionUpdate) SetOrderphone(s string) *TransactionUpdate {
	tu.orderphone = &s
	return tu
}

// SetNamabengkel sets the namabengkel field.
func (tu *TransactionUpdate) SetNamabengkel(s string) *TransactionUpdate {
	tu.namabengkel = &s
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	if tu.orderphone != nil {
		if err := transaction.OrderphoneValidator(*tu.orderphone); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"orderphone\": %v", err)
		}
	}
	if tu.namabengkel != nil {
		if err := transaction.NamabengkelValidator(*tu.namabengkel); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"namabengkel\": %v", err)
		}
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(tu.driver.Dialect())
		selector = builder.Select(transaction.FieldID).From(builder.Table(transaction.Table))
	)
	for _, p := range tu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = tu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := tu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(transaction.Table)
	)
	updater = updater.Where(sql.InInts(transaction.FieldID, ids...))
	if value := tu.orderphone; value != nil {
		updater.Set(transaction.FieldOrderphone, *value)
	}
	if value := tu.namabengkel; value != nil {
		updater.Set(transaction.FieldNamabengkel, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	id          int
	orderphone  *string
	namabengkel *string
}

// SetOrderphone sets the orderphone field.
func (tuo *TransactionUpdateOne) SetOrderphone(s string) *TransactionUpdateOne {
	tuo.orderphone = &s
	return tuo
}

// SetNamabengkel sets the namabengkel field.
func (tuo *TransactionUpdateOne) SetNamabengkel(s string) *TransactionUpdateOne {
	tuo.namabengkel = &s
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	if tuo.orderphone != nil {
		if err := transaction.OrderphoneValidator(*tuo.orderphone); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"orderphone\": %v", err)
		}
	}
	if tuo.namabengkel != nil {
		if err := transaction.NamabengkelValidator(*tuo.namabengkel); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"namabengkel\": %v", err)
		}
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (t *Transaction, err error) {
	var (
		builder  = sql.Dialect(tuo.driver.Dialect())
		selector = builder.Select(transaction.Columns...).From(builder.Table(transaction.Table))
	)
	transaction.ID(tuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = tuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		t = &Transaction{config: tuo.config}
		if err := t.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Transaction: %v", err)
		}
		id = t.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Transaction with id: %v", tuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Transaction with the same id: %v", tuo.id)
	}

	tx, err := tuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(transaction.Table)
	)
	updater = updater.Where(sql.InInts(transaction.FieldID, ids...))
	if value := tuo.orderphone; value != nil {
		updater.Set(transaction.FieldOrderphone, *value)
		t.Orderphone = *value
	}
	if value := tuo.namabengkel; value != nil {
		updater.Set(transaction.FieldNamabengkel, *value)
		t.Namabengkel = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return t, nil
}
