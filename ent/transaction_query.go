// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/confus1on/mbeledos/ent/predicate"
	"github.com/confus1on/mbeledos/ent/transaction"
	"github.com/facebookincubator/ent/dialect/sql"
)

// TransactionQuery is the builder for querying Transaction entities.
type TransactionQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Transaction
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (tq *TransactionQuery) Where(ps ...predicate.Transaction) *TransactionQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TransactionQuery) Limit(limit int) *TransactionQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TransactionQuery) Offset(offset int) *TransactionQuery {
	tq.offset = &offset
	return tq
}

// Order adds an order step to the query.
func (tq *TransactionQuery) Order(o ...Order) *TransactionQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Transaction entity in the query. Returns *ErrNotFound when no transaction was found.
func (tq *TransactionQuery) First(ctx context.Context) (*Transaction, error) {
	ts, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, &ErrNotFound{transaction.Label}
	}
	return ts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TransactionQuery) FirstX(ctx context.Context) *Transaction {
	t, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return t
}

// FirstID returns the first Transaction id in the query. Returns *ErrNotFound when no id was found.
func (tq *TransactionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{transaction.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (tq *TransactionQuery) FirstXID(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Transaction entity in the query, returns an error if not exactly one entity was returned.
func (tq *TransactionQuery) Only(ctx context.Context) (*Transaction, error) {
	ts, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ts) {
	case 1:
		return ts[0], nil
	case 0:
		return nil, &ErrNotFound{transaction.Label}
	default:
		return nil, &ErrNotSingular{transaction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TransactionQuery) OnlyX(ctx context.Context) *Transaction {
	t, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// OnlyID returns the only Transaction id in the query, returns an error if not exactly one id was returned.
func (tq *TransactionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{transaction.Label}
	default:
		err = &ErrNotSingular{transaction.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (tq *TransactionQuery) OnlyXID(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Transactions.
func (tq *TransactionQuery) All(ctx context.Context) ([]*Transaction, error) {
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TransactionQuery) AllX(ctx context.Context) []*Transaction {
	ts, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ts
}

// IDs executes the query and returns a list of Transaction ids.
func (tq *TransactionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(transaction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TransactionQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TransactionQuery) Count(ctx context.Context) (int, error) {
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TransactionQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TransactionQuery) Exist(ctx context.Context) (bool, error) {
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TransactionQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TransactionQuery) Clone() *TransactionQuery {
	return &TransactionQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]Order{}, tq.order...),
		unique:     append([]string{}, tq.unique...),
		predicates: append([]predicate.Transaction{}, tq.predicates...),
		// clone intermediate queries.
		sql: tq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Orderphone string `json:"orderphone,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Transaction.Query().
//		GroupBy(transaction.FieldOrderphone).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TransactionQuery) GroupBy(field string, fields ...string) *TransactionGroupBy {
	group := &TransactionGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = tq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Orderphone string `json:"orderphone,omitempty"`
//	}
//
//	client.Transaction.Query().
//		Select(transaction.FieldOrderphone).
//		Scan(ctx, &v)
//
func (tq *TransactionQuery) Select(field string, fields ...string) *TransactionSelect {
	selector := &TransactionSelect{config: tq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = tq.sqlQuery()
	return selector
}

func (tq *TransactionQuery) sqlAll(ctx context.Context) ([]*Transaction, error) {
	rows := &sql.Rows{}
	selector := tq.sqlQuery()
	if unique := tq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := tq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ts Transactions
	if err := ts.FromRows(rows); err != nil {
		return nil, err
	}
	ts.config(tq.config)
	return ts, nil
}

func (tq *TransactionQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := tq.sqlQuery()
	unique := []string{transaction.FieldID}
	if len(tq.unique) > 0 {
		unique = tq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := tq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (tq *TransactionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (tq *TransactionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(transaction.Table)
	selector := builder.Select(t1.Columns(transaction.Columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(transaction.Columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransactionGroupBy is the builder for group-by Transaction entities.
type TransactionGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TransactionGroupBy) Aggregate(fns ...Aggregate) *TransactionGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scan the result into the given value.
func (tgb *TransactionGroupBy) Scan(ctx context.Context, v interface{}) error {
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TransactionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (tgb *TransactionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransactionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TransactionGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (tgb *TransactionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransactionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TransactionGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (tgb *TransactionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransactionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TransactionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (tgb *TransactionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransactionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TransactionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TransactionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tgb.sqlQuery().Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TransactionGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql
	columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
	columns = append(columns, tgb.fields...)
	for _, fn := range tgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(tgb.fields...)
}

// TransactionSelect is the builder for select fields of Transaction entities.
type TransactionSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ts *TransactionSelect) Scan(ctx context.Context, v interface{}) error {
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TransactionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ts *TransactionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransactionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TransactionSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ts *TransactionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransactionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TransactionSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ts *TransactionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransactionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TransactionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ts *TransactionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransactionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TransactionSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TransactionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sqlQuery().Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ts *TransactionSelect) sqlQuery() sql.Querier {
	view := "transaction_view"
	return sql.Dialect(ts.driver.Dialect()).
		Select(ts.fields...).From(ts.sql.As(view))
}
