// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/confus1on/mbeledos/ent/transaction"
	"github.com/facebookincubator/ent/dialect/sql"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	orderphone  *string
	namabengkel *string
}

// SetOrderphone sets the orderphone field.
func (tc *TransactionCreate) SetOrderphone(s string) *TransactionCreate {
	tc.orderphone = &s
	return tc
}

// SetNamabengkel sets the namabengkel field.
func (tc *TransactionCreate) SetNamabengkel(s string) *TransactionCreate {
	tc.namabengkel = &s
	return tc
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	if tc.orderphone == nil {
		return nil, errors.New("ent: missing required field \"orderphone\"")
	}
	if err := transaction.OrderphoneValidator(*tc.orderphone); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"orderphone\": %v", err)
	}
	if tc.namabengkel == nil {
		return nil, errors.New("ent: missing required field \"namabengkel\"")
	}
	if err := transaction.NamabengkelValidator(*tc.namabengkel); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"namabengkel\": %v", err)
	}
	return tc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	var (
		builder = sql.Dialect(tc.driver.Dialect())
		t       = &Transaction{config: tc.config}
	)
	tx, err := tc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(transaction.Table).Default()
	if value := tc.orderphone; value != nil {
		insert.Set(transaction.FieldOrderphone, *value)
		t.Orderphone = *value
	}
	if value := tc.namabengkel; value != nil {
		insert.Set(transaction.FieldNamabengkel, *value)
		t.Namabengkel = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(transaction.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	t.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return t, nil
}
