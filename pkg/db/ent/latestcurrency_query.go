// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/latestcurrency"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// LatestCurrencyQuery is the builder for querying LatestCurrency entities.
type LatestCurrencyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LatestCurrency
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LatestCurrencyQuery builder.
func (lcq *LatestCurrencyQuery) Where(ps ...predicate.LatestCurrency) *LatestCurrencyQuery {
	lcq.predicates = append(lcq.predicates, ps...)
	return lcq
}

// Limit adds a limit step to the query.
func (lcq *LatestCurrencyQuery) Limit(limit int) *LatestCurrencyQuery {
	lcq.limit = &limit
	return lcq
}

// Offset adds an offset step to the query.
func (lcq *LatestCurrencyQuery) Offset(offset int) *LatestCurrencyQuery {
	lcq.offset = &offset
	return lcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcq *LatestCurrencyQuery) Unique(unique bool) *LatestCurrencyQuery {
	lcq.unique = &unique
	return lcq
}

// Order adds an order step to the query.
func (lcq *LatestCurrencyQuery) Order(o ...OrderFunc) *LatestCurrencyQuery {
	lcq.order = append(lcq.order, o...)
	return lcq
}

// First returns the first LatestCurrency entity from the query.
// Returns a *NotFoundError when no LatestCurrency was found.
func (lcq *LatestCurrencyQuery) First(ctx context.Context) (*LatestCurrency, error) {
	nodes, err := lcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{latestcurrency.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) FirstX(ctx context.Context) *LatestCurrency {
	node, err := lcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LatestCurrency ID from the query.
// Returns a *NotFoundError when no LatestCurrency ID was found.
func (lcq *LatestCurrencyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{latestcurrency.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LatestCurrency entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LatestCurrency entity is found.
// Returns a *NotFoundError when no LatestCurrency entities are found.
func (lcq *LatestCurrencyQuery) Only(ctx context.Context) (*LatestCurrency, error) {
	nodes, err := lcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{latestcurrency.Label}
	default:
		return nil, &NotSingularError{latestcurrency.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) OnlyX(ctx context.Context) *LatestCurrency {
	node, err := lcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LatestCurrency ID in the query.
// Returns a *NotSingularError when more than one LatestCurrency ID is found.
// Returns a *NotFoundError when no entities are found.
func (lcq *LatestCurrencyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{latestcurrency.Label}
	default:
		err = &NotSingularError{latestcurrency.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LatestCurrencies.
func (lcq *LatestCurrencyQuery) All(ctx context.Context) ([]*LatestCurrency, error) {
	if err := lcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) AllX(ctx context.Context) []*LatestCurrency {
	nodes, err := lcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LatestCurrency IDs.
func (lcq *LatestCurrencyQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := lcq.Select(latestcurrency.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcq *LatestCurrencyQuery) Count(ctx context.Context) (int, error) {
	if err := lcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) CountX(ctx context.Context) int {
	count, err := lcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcq *LatestCurrencyQuery) Exist(ctx context.Context) (bool, error) {
	if err := lcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lcq *LatestCurrencyQuery) ExistX(ctx context.Context) bool {
	exist, err := lcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LatestCurrencyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcq *LatestCurrencyQuery) Clone() *LatestCurrencyQuery {
	if lcq == nil {
		return nil
	}
	return &LatestCurrencyQuery{
		config:     lcq.config,
		limit:      lcq.limit,
		offset:     lcq.offset,
		order:      append([]OrderFunc{}, lcq.order...),
		predicates: append([]predicate.LatestCurrency{}, lcq.predicates...),
		// clone intermediate query.
		sql:    lcq.sql.Clone(),
		path:   lcq.path,
		unique: lcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LatestCurrency.Query().
//		GroupBy(latestcurrency.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lcq *LatestCurrencyQuery) GroupBy(field string, fields ...string) *LatestCurrencyGroupBy {
	grbuild := &LatestCurrencyGroupBy{config: lcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lcq.sqlQuery(ctx), nil
	}
	grbuild.label = latestcurrency.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.LatestCurrency.Query().
//		Select(latestcurrency.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (lcq *LatestCurrencyQuery) Select(fields ...string) *LatestCurrencySelect {
	lcq.fields = append(lcq.fields, fields...)
	selbuild := &LatestCurrencySelect{LatestCurrencyQuery: lcq}
	selbuild.label = latestcurrency.Label
	selbuild.flds, selbuild.scan = &lcq.fields, selbuild.Scan
	return selbuild
}

func (lcq *LatestCurrencyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lcq.fields {
		if !latestcurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lcq.path != nil {
		prev, err := lcq.path(ctx)
		if err != nil {
			return err
		}
		lcq.sql = prev
	}
	if latestcurrency.Policy == nil {
		return errors.New("ent: uninitialized latestcurrency.Policy (forgotten import ent/runtime?)")
	}
	if err := latestcurrency.Policy.EvalQuery(ctx, lcq); err != nil {
		return err
	}
	return nil
}

func (lcq *LatestCurrencyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LatestCurrency, error) {
	var (
		nodes = []*LatestCurrency{}
		_spec = lcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*LatestCurrency).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &LatestCurrency{config: lcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(lcq.modifiers) > 0 {
		_spec.Modifiers = lcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lcq *LatestCurrencyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcq.querySpec()
	if len(lcq.modifiers) > 0 {
		_spec.Modifiers = lcq.modifiers
	}
	_spec.Node.Columns = lcq.fields
	if len(lcq.fields) > 0 {
		_spec.Unique = lcq.unique != nil && *lcq.unique
	}
	return sqlgraph.CountNodes(ctx, lcq.driver, _spec)
}

func (lcq *LatestCurrencyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (lcq *LatestCurrencyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   latestcurrency.Table,
			Columns: latestcurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: latestcurrency.FieldID,
			},
		},
		From:   lcq.sql,
		Unique: true,
	}
	if unique := lcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, latestcurrency.FieldID)
		for i := range fields {
			if fields[i] != latestcurrency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcq *LatestCurrencyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcq.driver.Dialect())
	t1 := builder.Table(latestcurrency.Table)
	columns := lcq.fields
	if len(columns) == 0 {
		columns = latestcurrency.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lcq.sql != nil {
		selector = lcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lcq.unique != nil && *lcq.unique {
		selector.Distinct()
	}
	for _, m := range lcq.modifiers {
		m(selector)
	}
	for _, p := range lcq.predicates {
		p(selector)
	}
	for _, p := range lcq.order {
		p(selector)
	}
	if offset := lcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (lcq *LatestCurrencyQuery) ForUpdate(opts ...sql.LockOption) *LatestCurrencyQuery {
	if lcq.driver.Dialect() == dialect.Postgres {
		lcq.Unique(false)
	}
	lcq.modifiers = append(lcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return lcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (lcq *LatestCurrencyQuery) ForShare(opts ...sql.LockOption) *LatestCurrencyQuery {
	if lcq.driver.Dialect() == dialect.Postgres {
		lcq.Unique(false)
	}
	lcq.modifiers = append(lcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return lcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lcq *LatestCurrencyQuery) Modify(modifiers ...func(s *sql.Selector)) *LatestCurrencySelect {
	lcq.modifiers = append(lcq.modifiers, modifiers...)
	return lcq.Select()
}

// LatestCurrencyGroupBy is the group-by builder for LatestCurrency entities.
type LatestCurrencyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcgb *LatestCurrencyGroupBy) Aggregate(fns ...AggregateFunc) *LatestCurrencyGroupBy {
	lcgb.fns = append(lcgb.fns, fns...)
	return lcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lcgb *LatestCurrencyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lcgb.path(ctx)
	if err != nil {
		return err
	}
	lcgb.sql = query
	return lcgb.sqlScan(ctx, v)
}

func (lcgb *LatestCurrencyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lcgb.fields {
		if !latestcurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lcgb *LatestCurrencyGroupBy) sqlQuery() *sql.Selector {
	selector := lcgb.sql.Select()
	aggregation := make([]string, 0, len(lcgb.fns))
	for _, fn := range lcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lcgb.fields)+len(lcgb.fns))
		for _, f := range lcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lcgb.fields...)...)
}

// LatestCurrencySelect is the builder for selecting fields of LatestCurrency entities.
type LatestCurrencySelect struct {
	*LatestCurrencyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lcs *LatestCurrencySelect) Scan(ctx context.Context, v interface{}) error {
	if err := lcs.prepareQuery(ctx); err != nil {
		return err
	}
	lcs.sql = lcs.LatestCurrencyQuery.sqlQuery(ctx)
	return lcs.sqlScan(ctx, v)
}

func (lcs *LatestCurrencySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lcs.sql.Query()
	if err := lcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lcs *LatestCurrencySelect) Modify(modifiers ...func(s *sql.Selector)) *LatestCurrencySelect {
	lcs.modifiers = append(lcs.modifiers, modifiers...)
	return lcs
}
