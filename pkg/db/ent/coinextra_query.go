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
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinExtraQuery is the builder for querying CoinExtra entities.
type CoinExtraQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoinExtra
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinExtraQuery builder.
func (ceq *CoinExtraQuery) Where(ps ...predicate.CoinExtra) *CoinExtraQuery {
	ceq.predicates = append(ceq.predicates, ps...)
	return ceq
}

// Limit adds a limit step to the query.
func (ceq *CoinExtraQuery) Limit(limit int) *CoinExtraQuery {
	ceq.limit = &limit
	return ceq
}

// Offset adds an offset step to the query.
func (ceq *CoinExtraQuery) Offset(offset int) *CoinExtraQuery {
	ceq.offset = &offset
	return ceq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ceq *CoinExtraQuery) Unique(unique bool) *CoinExtraQuery {
	ceq.unique = &unique
	return ceq
}

// Order adds an order step to the query.
func (ceq *CoinExtraQuery) Order(o ...OrderFunc) *CoinExtraQuery {
	ceq.order = append(ceq.order, o...)
	return ceq
}

// First returns the first CoinExtra entity from the query.
// Returns a *NotFoundError when no CoinExtra was found.
func (ceq *CoinExtraQuery) First(ctx context.Context) (*CoinExtra, error) {
	nodes, err := ceq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coinextra.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ceq *CoinExtraQuery) FirstX(ctx context.Context) *CoinExtra {
	node, err := ceq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinExtra ID from the query.
// Returns a *NotFoundError when no CoinExtra ID was found.
func (ceq *CoinExtraQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ceq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coinextra.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ceq *CoinExtraQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ceq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinExtra entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CoinExtra entity is found.
// Returns a *NotFoundError when no CoinExtra entities are found.
func (ceq *CoinExtraQuery) Only(ctx context.Context) (*CoinExtra, error) {
	nodes, err := ceq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coinextra.Label}
	default:
		return nil, &NotSingularError{coinextra.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ceq *CoinExtraQuery) OnlyX(ctx context.Context) *CoinExtra {
	node, err := ceq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinExtra ID in the query.
// Returns a *NotSingularError when more than one CoinExtra ID is found.
// Returns a *NotFoundError when no entities are found.
func (ceq *CoinExtraQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ceq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coinextra.Label}
	default:
		err = &NotSingularError{coinextra.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ceq *CoinExtraQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ceq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinExtras.
func (ceq *CoinExtraQuery) All(ctx context.Context) ([]*CoinExtra, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ceq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ceq *CoinExtraQuery) AllX(ctx context.Context) []*CoinExtra {
	nodes, err := ceq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinExtra IDs.
func (ceq *CoinExtraQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ceq.Select(coinextra.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ceq *CoinExtraQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ceq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ceq *CoinExtraQuery) Count(ctx context.Context) (int, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ceq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ceq *CoinExtraQuery) CountX(ctx context.Context) int {
	count, err := ceq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ceq *CoinExtraQuery) Exist(ctx context.Context) (bool, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ceq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ceq *CoinExtraQuery) ExistX(ctx context.Context) bool {
	exist, err := ceq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinExtraQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ceq *CoinExtraQuery) Clone() *CoinExtraQuery {
	if ceq == nil {
		return nil
	}
	return &CoinExtraQuery{
		config:     ceq.config,
		limit:      ceq.limit,
		offset:     ceq.offset,
		order:      append([]OrderFunc{}, ceq.order...),
		predicates: append([]predicate.CoinExtra{}, ceq.predicates...),
		// clone intermediate query.
		sql:    ceq.sql.Clone(),
		path:   ceq.path,
		unique: ceq.unique,
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
//	client.CoinExtra.Query().
//		GroupBy(coinextra.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ceq *CoinExtraQuery) GroupBy(field string, fields ...string) *CoinExtraGroupBy {
	grbuild := &CoinExtraGroupBy{config: ceq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ceq.sqlQuery(ctx), nil
	}
	grbuild.label = coinextra.Label
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
//	client.CoinExtra.Query().
//		Select(coinextra.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ceq *CoinExtraQuery) Select(fields ...string) *CoinExtraSelect {
	ceq.fields = append(ceq.fields, fields...)
	selbuild := &CoinExtraSelect{CoinExtraQuery: ceq}
	selbuild.label = coinextra.Label
	selbuild.flds, selbuild.scan = &ceq.fields, selbuild.Scan
	return selbuild
}

func (ceq *CoinExtraQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ceq.fields {
		if !coinextra.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ceq.path != nil {
		prev, err := ceq.path(ctx)
		if err != nil {
			return err
		}
		ceq.sql = prev
	}
	if coinextra.Policy == nil {
		return errors.New("ent: uninitialized coinextra.Policy (forgotten import ent/runtime?)")
	}
	if err := coinextra.Policy.EvalQuery(ctx, ceq); err != nil {
		return err
	}
	return nil
}

func (ceq *CoinExtraQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CoinExtra, error) {
	var (
		nodes = []*CoinExtra{}
		_spec = ceq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CoinExtra).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CoinExtra{config: ceq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ceq.modifiers) > 0 {
		_spec.Modifiers = ceq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ceq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ceq *CoinExtraQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ceq.querySpec()
	if len(ceq.modifiers) > 0 {
		_spec.Modifiers = ceq.modifiers
	}
	_spec.Node.Columns = ceq.fields
	if len(ceq.fields) > 0 {
		_spec.Unique = ceq.unique != nil && *ceq.unique
	}
	return sqlgraph.CountNodes(ctx, ceq.driver, _spec)
}

func (ceq *CoinExtraQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ceq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ceq *CoinExtraQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinextra.Table,
			Columns: coinextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinextra.FieldID,
			},
		},
		From:   ceq.sql,
		Unique: true,
	}
	if unique := ceq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ceq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinextra.FieldID)
		for i := range fields {
			if fields[i] != coinextra.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ceq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ceq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ceq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ceq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ceq *CoinExtraQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ceq.driver.Dialect())
	t1 := builder.Table(coinextra.Table)
	columns := ceq.fields
	if len(columns) == 0 {
		columns = coinextra.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ceq.sql != nil {
		selector = ceq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ceq.unique != nil && *ceq.unique {
		selector.Distinct()
	}
	for _, m := range ceq.modifiers {
		m(selector)
	}
	for _, p := range ceq.predicates {
		p(selector)
	}
	for _, p := range ceq.order {
		p(selector)
	}
	if offset := ceq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ceq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ceq *CoinExtraQuery) ForUpdate(opts ...sql.LockOption) *CoinExtraQuery {
	if ceq.driver.Dialect() == dialect.Postgres {
		ceq.Unique(false)
	}
	ceq.modifiers = append(ceq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ceq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ceq *CoinExtraQuery) ForShare(opts ...sql.LockOption) *CoinExtraQuery {
	if ceq.driver.Dialect() == dialect.Postgres {
		ceq.Unique(false)
	}
	ceq.modifiers = append(ceq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ceq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ceq *CoinExtraQuery) Modify(modifiers ...func(s *sql.Selector)) *CoinExtraSelect {
	ceq.modifiers = append(ceq.modifiers, modifiers...)
	return ceq.Select()
}

// CoinExtraGroupBy is the group-by builder for CoinExtra entities.
type CoinExtraGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cegb *CoinExtraGroupBy) Aggregate(fns ...AggregateFunc) *CoinExtraGroupBy {
	cegb.fns = append(cegb.fns, fns...)
	return cegb
}

// Scan applies the group-by query and scans the result into the given value.
func (cegb *CoinExtraGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cegb.path(ctx)
	if err != nil {
		return err
	}
	cegb.sql = query
	return cegb.sqlScan(ctx, v)
}

func (cegb *CoinExtraGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cegb.fields {
		if !coinextra.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cegb *CoinExtraGroupBy) sqlQuery() *sql.Selector {
	selector := cegb.sql.Select()
	aggregation := make([]string, 0, len(cegb.fns))
	for _, fn := range cegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cegb.fields)+len(cegb.fns))
		for _, f := range cegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cegb.fields...)...)
}

// CoinExtraSelect is the builder for selecting fields of CoinExtra entities.
type CoinExtraSelect struct {
	*CoinExtraQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ces *CoinExtraSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ces.prepareQuery(ctx); err != nil {
		return err
	}
	ces.sql = ces.CoinExtraQuery.sqlQuery(ctx)
	return ces.sqlScan(ctx, v)
}

func (ces *CoinExtraSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ces.sql.Query()
	if err := ces.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ces *CoinExtraSelect) Modify(modifiers ...func(s *sql.Selector)) *CoinExtraSelect {
	ces.modifiers = append(ces.modifiers, modifiers...)
	return ces
}
