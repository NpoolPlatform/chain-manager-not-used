// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TranUpdate is the builder for updating Tran entities.
type TranUpdate struct {
	config
	hooks     []Hook
	mutation  *TranMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TranUpdate builder.
func (tu *TranUpdate) Where(ps ...predicate.Tran) *TranUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TranUpdate) SetCreatedAt(u uint32) *TranUpdate {
	tu.mutation.ResetCreatedAt()
	tu.mutation.SetCreatedAt(u)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TranUpdate) SetNillableCreatedAt(u *uint32) *TranUpdate {
	if u != nil {
		tu.SetCreatedAt(*u)
	}
	return tu
}

// AddCreatedAt adds u to the "created_at" field.
func (tu *TranUpdate) AddCreatedAt(u int32) *TranUpdate {
	tu.mutation.AddCreatedAt(u)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TranUpdate) SetUpdatedAt(u uint32) *TranUpdate {
	tu.mutation.ResetUpdatedAt()
	tu.mutation.SetUpdatedAt(u)
	return tu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tu *TranUpdate) AddUpdatedAt(u int32) *TranUpdate {
	tu.mutation.AddUpdatedAt(u)
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TranUpdate) SetDeletedAt(u uint32) *TranUpdate {
	tu.mutation.ResetDeletedAt()
	tu.mutation.SetDeletedAt(u)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TranUpdate) SetNillableDeletedAt(u *uint32) *TranUpdate {
	if u != nil {
		tu.SetDeletedAt(*u)
	}
	return tu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tu *TranUpdate) AddDeletedAt(u int32) *TranUpdate {
	tu.mutation.AddDeletedAt(u)
	return tu
}

// SetAppID sets the "app_id" field.
func (tu *TranUpdate) SetAppID(u uuid.UUID) *TranUpdate {
	tu.mutation.SetAppID(u)
	return tu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (tu *TranUpdate) SetNillableAppID(u *uuid.UUID) *TranUpdate {
	if u != nil {
		tu.SetAppID(*u)
	}
	return tu
}

// ClearAppID clears the value of the "app_id" field.
func (tu *TranUpdate) ClearAppID() *TranUpdate {
	tu.mutation.ClearAppID()
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TranUpdate) SetUserID(u uuid.UUID) *TranUpdate {
	tu.mutation.SetUserID(u)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TranUpdate) SetNillableUserID(u *uuid.UUID) *TranUpdate {
	if u != nil {
		tu.SetUserID(*u)
	}
	return tu
}

// ClearUserID clears the value of the "user_id" field.
func (tu *TranUpdate) ClearUserID() *TranUpdate {
	tu.mutation.ClearUserID()
	return tu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (tu *TranUpdate) SetCoinTypeID(u uuid.UUID) *TranUpdate {
	tu.mutation.SetCoinTypeID(u)
	return tu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (tu *TranUpdate) SetNillableCoinTypeID(u *uuid.UUID) *TranUpdate {
	if u != nil {
		tu.SetCoinTypeID(*u)
	}
	return tu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (tu *TranUpdate) ClearCoinTypeID() *TranUpdate {
	tu.mutation.ClearCoinTypeID()
	return tu
}

// SetIncoming sets the "incoming" field.
func (tu *TranUpdate) SetIncoming(d decimal.Decimal) *TranUpdate {
	tu.mutation.SetIncoming(d)
	return tu
}

// SetNillableIncoming sets the "incoming" field if the given value is not nil.
func (tu *TranUpdate) SetNillableIncoming(d *decimal.Decimal) *TranUpdate {
	if d != nil {
		tu.SetIncoming(*d)
	}
	return tu
}

// ClearIncoming clears the value of the "incoming" field.
func (tu *TranUpdate) ClearIncoming() *TranUpdate {
	tu.mutation.ClearIncoming()
	return tu
}

// SetLocked sets the "locked" field.
func (tu *TranUpdate) SetLocked(d decimal.Decimal) *TranUpdate {
	tu.mutation.SetLocked(d)
	return tu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (tu *TranUpdate) SetNillableLocked(d *decimal.Decimal) *TranUpdate {
	if d != nil {
		tu.SetLocked(*d)
	}
	return tu
}

// ClearLocked clears the value of the "locked" field.
func (tu *TranUpdate) ClearLocked() *TranUpdate {
	tu.mutation.ClearLocked()
	return tu
}

// SetOutcoming sets the "outcoming" field.
func (tu *TranUpdate) SetOutcoming(d decimal.Decimal) *TranUpdate {
	tu.mutation.SetOutcoming(d)
	return tu
}

// SetNillableOutcoming sets the "outcoming" field if the given value is not nil.
func (tu *TranUpdate) SetNillableOutcoming(d *decimal.Decimal) *TranUpdate {
	if d != nil {
		tu.SetOutcoming(*d)
	}
	return tu
}

// ClearOutcoming clears the value of the "outcoming" field.
func (tu *TranUpdate) ClearOutcoming() *TranUpdate {
	tu.mutation.ClearOutcoming()
	return tu
}

// SetSpendable sets the "spendable" field.
func (tu *TranUpdate) SetSpendable(d decimal.Decimal) *TranUpdate {
	tu.mutation.SetSpendable(d)
	return tu
}

// SetNillableSpendable sets the "spendable" field if the given value is not nil.
func (tu *TranUpdate) SetNillableSpendable(d *decimal.Decimal) *TranUpdate {
	if d != nil {
		tu.SetSpendable(*d)
	}
	return tu
}

// ClearSpendable clears the value of the "spendable" field.
func (tu *TranUpdate) ClearSpendable() *TranUpdate {
	tu.mutation.ClearSpendable()
	return tu
}

// Mutation returns the TranMutation object of the builder.
func (tu *TranUpdate) Mutation() *TranMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TranUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TranUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TranUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TranUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TranUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		if tran.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tran.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tran.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TranUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TranUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TranUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tran.Table,
			Columns: tran.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tran.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldDeletedAt,
		})
	}
	if value, ok := tu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldDeletedAt,
		})
	}
	if value, ok := tu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldAppID,
		})
	}
	if tu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldAppID,
		})
	}
	if value, ok := tu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldUserID,
		})
	}
	if tu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldUserID,
		})
	}
	if value, ok := tu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldCoinTypeID,
		})
	}
	if tu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldCoinTypeID,
		})
	}
	if value, ok := tu.mutation.Incoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldIncoming,
		})
	}
	if tu.mutation.IncomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldIncoming,
		})
	}
	if value, ok := tu.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldLocked,
		})
	}
	if tu.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldLocked,
		})
	}
	if value, ok := tu.mutation.Outcoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldOutcoming,
		})
	}
	if tu.mutation.OutcomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldOutcoming,
		})
	}
	if value, ok := tu.mutation.Spendable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldSpendable,
		})
	}
	if tu.mutation.SpendableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldSpendable,
		})
	}
	_spec.Modifiers = tu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tran.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TranUpdateOne is the builder for updating a single Tran entity.
type TranUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TranMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TranUpdateOne) SetCreatedAt(u uint32) *TranUpdateOne {
	tuo.mutation.ResetCreatedAt()
	tuo.mutation.SetCreatedAt(u)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableCreatedAt(u *uint32) *TranUpdateOne {
	if u != nil {
		tuo.SetCreatedAt(*u)
	}
	return tuo
}

// AddCreatedAt adds u to the "created_at" field.
func (tuo *TranUpdateOne) AddCreatedAt(u int32) *TranUpdateOne {
	tuo.mutation.AddCreatedAt(u)
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TranUpdateOne) SetUpdatedAt(u uint32) *TranUpdateOne {
	tuo.mutation.ResetUpdatedAt()
	tuo.mutation.SetUpdatedAt(u)
	return tuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tuo *TranUpdateOne) AddUpdatedAt(u int32) *TranUpdateOne {
	tuo.mutation.AddUpdatedAt(u)
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TranUpdateOne) SetDeletedAt(u uint32) *TranUpdateOne {
	tuo.mutation.ResetDeletedAt()
	tuo.mutation.SetDeletedAt(u)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableDeletedAt(u *uint32) *TranUpdateOne {
	if u != nil {
		tuo.SetDeletedAt(*u)
	}
	return tuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tuo *TranUpdateOne) AddDeletedAt(u int32) *TranUpdateOne {
	tuo.mutation.AddDeletedAt(u)
	return tuo
}

// SetAppID sets the "app_id" field.
func (tuo *TranUpdateOne) SetAppID(u uuid.UUID) *TranUpdateOne {
	tuo.mutation.SetAppID(u)
	return tuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableAppID(u *uuid.UUID) *TranUpdateOne {
	if u != nil {
		tuo.SetAppID(*u)
	}
	return tuo
}

// ClearAppID clears the value of the "app_id" field.
func (tuo *TranUpdateOne) ClearAppID() *TranUpdateOne {
	tuo.mutation.ClearAppID()
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TranUpdateOne) SetUserID(u uuid.UUID) *TranUpdateOne {
	tuo.mutation.SetUserID(u)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableUserID(u *uuid.UUID) *TranUpdateOne {
	if u != nil {
		tuo.SetUserID(*u)
	}
	return tuo
}

// ClearUserID clears the value of the "user_id" field.
func (tuo *TranUpdateOne) ClearUserID() *TranUpdateOne {
	tuo.mutation.ClearUserID()
	return tuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (tuo *TranUpdateOne) SetCoinTypeID(u uuid.UUID) *TranUpdateOne {
	tuo.mutation.SetCoinTypeID(u)
	return tuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *TranUpdateOne {
	if u != nil {
		tuo.SetCoinTypeID(*u)
	}
	return tuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (tuo *TranUpdateOne) ClearCoinTypeID() *TranUpdateOne {
	tuo.mutation.ClearCoinTypeID()
	return tuo
}

// SetIncoming sets the "incoming" field.
func (tuo *TranUpdateOne) SetIncoming(d decimal.Decimal) *TranUpdateOne {
	tuo.mutation.SetIncoming(d)
	return tuo
}

// SetNillableIncoming sets the "incoming" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableIncoming(d *decimal.Decimal) *TranUpdateOne {
	if d != nil {
		tuo.SetIncoming(*d)
	}
	return tuo
}

// ClearIncoming clears the value of the "incoming" field.
func (tuo *TranUpdateOne) ClearIncoming() *TranUpdateOne {
	tuo.mutation.ClearIncoming()
	return tuo
}

// SetLocked sets the "locked" field.
func (tuo *TranUpdateOne) SetLocked(d decimal.Decimal) *TranUpdateOne {
	tuo.mutation.SetLocked(d)
	return tuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableLocked(d *decimal.Decimal) *TranUpdateOne {
	if d != nil {
		tuo.SetLocked(*d)
	}
	return tuo
}

// ClearLocked clears the value of the "locked" field.
func (tuo *TranUpdateOne) ClearLocked() *TranUpdateOne {
	tuo.mutation.ClearLocked()
	return tuo
}

// SetOutcoming sets the "outcoming" field.
func (tuo *TranUpdateOne) SetOutcoming(d decimal.Decimal) *TranUpdateOne {
	tuo.mutation.SetOutcoming(d)
	return tuo
}

// SetNillableOutcoming sets the "outcoming" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableOutcoming(d *decimal.Decimal) *TranUpdateOne {
	if d != nil {
		tuo.SetOutcoming(*d)
	}
	return tuo
}

// ClearOutcoming clears the value of the "outcoming" field.
func (tuo *TranUpdateOne) ClearOutcoming() *TranUpdateOne {
	tuo.mutation.ClearOutcoming()
	return tuo
}

// SetSpendable sets the "spendable" field.
func (tuo *TranUpdateOne) SetSpendable(d decimal.Decimal) *TranUpdateOne {
	tuo.mutation.SetSpendable(d)
	return tuo
}

// SetNillableSpendable sets the "spendable" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableSpendable(d *decimal.Decimal) *TranUpdateOne {
	if d != nil {
		tuo.SetSpendable(*d)
	}
	return tuo
}

// ClearSpendable clears the value of the "spendable" field.
func (tuo *TranUpdateOne) ClearSpendable() *TranUpdateOne {
	tuo.mutation.ClearSpendable()
	return tuo
}

// Mutation returns the TranMutation object of the builder.
func (tuo *TranUpdateOne) Mutation() *TranMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TranUpdateOne) Select(field string, fields ...string) *TranUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tran entity.
func (tuo *TranUpdateOne) Save(ctx context.Context) (*Tran, error) {
	var (
		err  error
		node *Tran
	)
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tran)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TranMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TranUpdateOne) SaveX(ctx context.Context) *Tran {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TranUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TranUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TranUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		if tran.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tran.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tran.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TranUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TranUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TranUpdateOne) sqlSave(ctx context.Context) (_node *Tran, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tran.Table,
			Columns: tran.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tran.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tran.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tran.FieldID)
		for _, f := range fields {
			if !tran.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tran.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldDeletedAt,
		})
	}
	if value, ok := tuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldDeletedAt,
		})
	}
	if value, ok := tuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldAppID,
		})
	}
	if tuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldAppID,
		})
	}
	if value, ok := tuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldUserID,
		})
	}
	if tuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldUserID,
		})
	}
	if value, ok := tuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldCoinTypeID,
		})
	}
	if tuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldCoinTypeID,
		})
	}
	if value, ok := tuo.mutation.Incoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldIncoming,
		})
	}
	if tuo.mutation.IncomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldIncoming,
		})
	}
	if value, ok := tuo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldLocked,
		})
	}
	if tuo.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldLocked,
		})
	}
	if value, ok := tuo.mutation.Outcoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldOutcoming,
		})
	}
	if tuo.mutation.OutcomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldOutcoming,
		})
	}
	if value, ok := tuo.mutation.Spendable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldSpendable,
		})
	}
	if tuo.mutation.SpendableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldSpendable,
		})
	}
	_spec.Modifiers = tuo.modifiers
	_node = &Tran{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tran.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}