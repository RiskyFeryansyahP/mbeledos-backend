// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/confus1on/mbeledos/ent/bengkel"
	"github.com/confus1on/mbeledos/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// BengkelQuery is the builder for querying Bengkel entities.
type BengkelQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Bengkel
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (bq *BengkelQuery) Where(ps ...predicate.Bengkel) *BengkelQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BengkelQuery) Limit(limit int) *BengkelQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BengkelQuery) Offset(offset int) *BengkelQuery {
	bq.offset = &offset
	return bq
}

// Order adds an order step to the query.
func (bq *BengkelQuery) Order(o ...Order) *BengkelQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Bengkel entity in the query. Returns *ErrNotFound when no bengkel was found.
func (bq *BengkelQuery) First(ctx context.Context) (*Bengkel, error) {
	bs, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, &ErrNotFound{bengkel.Label}
	}
	return bs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BengkelQuery) FirstX(ctx context.Context) *Bengkel {
	b, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return b
}

// FirstID returns the first Bengkel id in the query. Returns *ErrNotFound when no id was found.
func (bq *BengkelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{bengkel.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bq *BengkelQuery) FirstXID(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Bengkel entity in the query, returns an error if not exactly one entity was returned.
func (bq *BengkelQuery) Only(ctx context.Context) (*Bengkel, error) {
	bs, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(bs) {
	case 1:
		return bs[0], nil
	case 0:
		return nil, &ErrNotFound{bengkel.Label}
	default:
		return nil, &ErrNotSingular{bengkel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BengkelQuery) OnlyX(ctx context.Context) *Bengkel {
	b, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// OnlyID returns the only Bengkel id in the query, returns an error if not exactly one id was returned.
func (bq *BengkelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{bengkel.Label}
	default:
		err = &ErrNotSingular{bengkel.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (bq *BengkelQuery) OnlyXID(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bengkels.
func (bq *BengkelQuery) All(ctx context.Context) ([]*Bengkel, error) {
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BengkelQuery) AllX(ctx context.Context) []*Bengkel {
	bs, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return bs
}

// IDs executes the query and returns a list of Bengkel ids.
func (bq *BengkelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(bengkel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BengkelQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BengkelQuery) Count(ctx context.Context) (int, error) {
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BengkelQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BengkelQuery) Exist(ctx context.Context) (bool, error) {
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BengkelQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BengkelQuery) Clone() *BengkelQuery {
	return &BengkelQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]Order{}, bq.order...),
		unique:     append([]string{}, bq.unique...),
		predicates: append([]predicate.Bengkel{}, bq.predicates...),
		// clone intermediate queries.
		sql: bq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		KodeBengkel string `json:"kode_bengkel,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bengkel.Query().
//		GroupBy(bengkel.FieldKodeBengkel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BengkelQuery) GroupBy(field string, fields ...string) *BengkelGroupBy {
	group := &BengkelGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = bq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		KodeBengkel string `json:"kode_bengkel,omitempty"`
//	}
//
//	client.Bengkel.Query().
//		Select(bengkel.FieldKodeBengkel).
//		Scan(ctx, &v)
//
func (bq *BengkelQuery) Select(field string, fields ...string) *BengkelSelect {
	selector := &BengkelSelect{config: bq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = bq.sqlQuery()
	return selector
}

func (bq *BengkelQuery) sqlAll(ctx context.Context) ([]*Bengkel, error) {
	rows := &sql.Rows{}
	selector := bq.sqlQuery()
	if unique := bq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := bq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var bs Bengkels
	if err := bs.FromRows(rows); err != nil {
		return nil, err
	}
	bs.config(bq.config)
	return bs, nil
}

func (bq *BengkelQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := bq.sqlQuery()
	unique := []string{bengkel.FieldID}
	if len(bq.unique) > 0 {
		unique = bq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := bq.driver.Query(ctx, query, args, rows); err != nil {
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

func (bq *BengkelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bq *BengkelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bengkel.Table)
	selector := builder.Select(t1.Columns(bengkel.Columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(bengkel.Columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BengkelGroupBy is the builder for group-by Bengkel entities.
type BengkelGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BengkelGroupBy) Aggregate(fns ...Aggregate) *BengkelGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bgb *BengkelGroupBy) Scan(ctx context.Context, v interface{}) error {
	return bgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BengkelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bgb *BengkelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BengkelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BengkelGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bgb *BengkelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BengkelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BengkelGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bgb *BengkelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BengkelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BengkelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bgb *BengkelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BengkelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BengkelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BengkelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bgb.sqlQuery().Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BengkelGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql
	columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
	columns = append(columns, bgb.fields...)
	for _, fn := range bgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(bgb.fields...)
}

// BengkelSelect is the builder for select fields of Bengkel entities.
type BengkelSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (bs *BengkelSelect) Scan(ctx context.Context, v interface{}) error {
	return bs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bs *BengkelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (bs *BengkelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BengkelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bs *BengkelSelect) StringsX(ctx context.Context) []string {
	v, err := bs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (bs *BengkelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BengkelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bs *BengkelSelect) IntsX(ctx context.Context) []int {
	v, err := bs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (bs *BengkelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BengkelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bs *BengkelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (bs *BengkelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BengkelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bs *BengkelSelect) BoolsX(ctx context.Context) []bool {
	v, err := bs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bs *BengkelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sqlQuery().Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bs *BengkelSelect) sqlQuery() sql.Querier {
	view := "bengkel_view"
	return sql.Dialect(bs.driver.Dialect()).
		Select(bs.fields...).From(bs.sql.As(view))
}
