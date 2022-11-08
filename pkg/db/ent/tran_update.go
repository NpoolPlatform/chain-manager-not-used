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

// SetFromAccountID sets the "from_account_id" field.
func (tu *TranUpdate) SetFromAccountID(u uuid.UUID) *TranUpdate {
	tu.mutation.SetFromAccountID(u)
	return tu
}

// SetNillableFromAccountID sets the "from_account_id" field if the given value is not nil.
func (tu *TranUpdate) SetNillableFromAccountID(u *uuid.UUID) *TranUpdate {
	if u != nil {
		tu.SetFromAccountID(*u)
	}
	return tu
}

// ClearFromAccountID clears the value of the "from_account_id" field.
func (tu *TranUpdate) ClearFromAccountID() *TranUpdate {
	tu.mutation.ClearFromAccountID()
	return tu
}

// SetToAccountID sets the "to_account_id" field.
func (tu *TranUpdate) SetToAccountID(u uuid.UUID) *TranUpdate {
	tu.mutation.SetToAccountID(u)
	return tu
}

// SetNillableToAccountID sets the "to_account_id" field if the given value is not nil.
func (tu *TranUpdate) SetNillableToAccountID(u *uuid.UUID) *TranUpdate {
	if u != nil {
		tu.SetToAccountID(*u)
	}
	return tu
}

// ClearToAccountID clears the value of the "to_account_id" field.
func (tu *TranUpdate) ClearToAccountID() *TranUpdate {
	tu.mutation.ClearToAccountID()
	return tu
}

// SetAmount sets the "amount" field.
func (tu *TranUpdate) SetAmount(d decimal.Decimal) *TranUpdate {
	tu.mutation.SetAmount(d)
	return tu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tu *TranUpdate) SetNillableAmount(d *decimal.Decimal) *TranUpdate {
	if d != nil {
		tu.SetAmount(*d)
	}
	return tu
}

// ClearAmount clears the value of the "amount" field.
func (tu *TranUpdate) ClearAmount() *TranUpdate {
	tu.mutation.ClearAmount()
	return tu
}

// SetFeeAmount sets the "fee_amount" field.
func (tu *TranUpdate) SetFeeAmount(d decimal.Decimal) *TranUpdate {
	tu.mutation.SetFeeAmount(d)
	return tu
}

// SetNillableFeeAmount sets the "fee_amount" field if the given value is not nil.
func (tu *TranUpdate) SetNillableFeeAmount(d *decimal.Decimal) *TranUpdate {
	if d != nil {
		tu.SetFeeAmount(*d)
	}
	return tu
}

// ClearFeeAmount clears the value of the "fee_amount" field.
func (tu *TranUpdate) ClearFeeAmount() *TranUpdate {
	tu.mutation.ClearFeeAmount()
	return tu
}

// SetChainTxID sets the "chain_tx_id" field.
func (tu *TranUpdate) SetChainTxID(s string) *TranUpdate {
	tu.mutation.SetChainTxID(s)
	return tu
}

// SetNillableChainTxID sets the "chain_tx_id" field if the given value is not nil.
func (tu *TranUpdate) SetNillableChainTxID(s *string) *TranUpdate {
	if s != nil {
		tu.SetChainTxID(*s)
	}
	return tu
}

// ClearChainTxID clears the value of the "chain_tx_id" field.
func (tu *TranUpdate) ClearChainTxID() *TranUpdate {
	tu.mutation.ClearChainTxID()
	return tu
}

// SetState sets the "state" field.
func (tu *TranUpdate) SetState(s string) *TranUpdate {
	tu.mutation.SetState(s)
	return tu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (tu *TranUpdate) SetNillableState(s *string) *TranUpdate {
	if s != nil {
		tu.SetState(*s)
	}
	return tu
}

// ClearState clears the value of the "state" field.
func (tu *TranUpdate) ClearState() *TranUpdate {
	tu.mutation.ClearState()
	return tu
}

// SetExtra sets the "extra" field.
func (tu *TranUpdate) SetExtra(s string) *TranUpdate {
	tu.mutation.SetExtra(s)
	return tu
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (tu *TranUpdate) SetNillableExtra(s *string) *TranUpdate {
	if s != nil {
		tu.SetExtra(*s)
	}
	return tu
}

// ClearExtra clears the value of the "extra" field.
func (tu *TranUpdate) ClearExtra() *TranUpdate {
	tu.mutation.ClearExtra()
	return tu
}

// SetType sets the "type" field.
func (tu *TranUpdate) SetType(s string) *TranUpdate {
	tu.mutation.SetType(s)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TranUpdate) SetNillableType(s *string) *TranUpdate {
	if s != nil {
		tu.SetType(*s)
	}
	return tu
}

// ClearType clears the value of the "type" field.
func (tu *TranUpdate) ClearType() *TranUpdate {
	tu.mutation.ClearType()
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
	if value, ok := tu.mutation.FromAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldFromAccountID,
		})
	}
	if tu.mutation.FromAccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldFromAccountID,
		})
	}
	if value, ok := tu.mutation.ToAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldToAccountID,
		})
	}
	if tu.mutation.ToAccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldToAccountID,
		})
	}
	if value, ok := tu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldAmount,
		})
	}
	if tu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldAmount,
		})
	}
	if value, ok := tu.mutation.FeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldFeeAmount,
		})
	}
	if tu.mutation.FeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldFeeAmount,
		})
	}
	if value, ok := tu.mutation.ChainTxID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldChainTxID,
		})
	}
	if tu.mutation.ChainTxIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldChainTxID,
		})
	}
	if value, ok := tu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldState,
		})
	}
	if tu.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldState,
		})
	}
	if value, ok := tu.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldExtra,
		})
	}
	if tu.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldExtra,
		})
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldType,
		})
	}
	if tu.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldType,
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

// SetFromAccountID sets the "from_account_id" field.
func (tuo *TranUpdateOne) SetFromAccountID(u uuid.UUID) *TranUpdateOne {
	tuo.mutation.SetFromAccountID(u)
	return tuo
}

// SetNillableFromAccountID sets the "from_account_id" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableFromAccountID(u *uuid.UUID) *TranUpdateOne {
	if u != nil {
		tuo.SetFromAccountID(*u)
	}
	return tuo
}

// ClearFromAccountID clears the value of the "from_account_id" field.
func (tuo *TranUpdateOne) ClearFromAccountID() *TranUpdateOne {
	tuo.mutation.ClearFromAccountID()
	return tuo
}

// SetToAccountID sets the "to_account_id" field.
func (tuo *TranUpdateOne) SetToAccountID(u uuid.UUID) *TranUpdateOne {
	tuo.mutation.SetToAccountID(u)
	return tuo
}

// SetNillableToAccountID sets the "to_account_id" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableToAccountID(u *uuid.UUID) *TranUpdateOne {
	if u != nil {
		tuo.SetToAccountID(*u)
	}
	return tuo
}

// ClearToAccountID clears the value of the "to_account_id" field.
func (tuo *TranUpdateOne) ClearToAccountID() *TranUpdateOne {
	tuo.mutation.ClearToAccountID()
	return tuo
}

// SetAmount sets the "amount" field.
func (tuo *TranUpdateOne) SetAmount(d decimal.Decimal) *TranUpdateOne {
	tuo.mutation.SetAmount(d)
	return tuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableAmount(d *decimal.Decimal) *TranUpdateOne {
	if d != nil {
		tuo.SetAmount(*d)
	}
	return tuo
}

// ClearAmount clears the value of the "amount" field.
func (tuo *TranUpdateOne) ClearAmount() *TranUpdateOne {
	tuo.mutation.ClearAmount()
	return tuo
}

// SetFeeAmount sets the "fee_amount" field.
func (tuo *TranUpdateOne) SetFeeAmount(d decimal.Decimal) *TranUpdateOne {
	tuo.mutation.SetFeeAmount(d)
	return tuo
}

// SetNillableFeeAmount sets the "fee_amount" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableFeeAmount(d *decimal.Decimal) *TranUpdateOne {
	if d != nil {
		tuo.SetFeeAmount(*d)
	}
	return tuo
}

// ClearFeeAmount clears the value of the "fee_amount" field.
func (tuo *TranUpdateOne) ClearFeeAmount() *TranUpdateOne {
	tuo.mutation.ClearFeeAmount()
	return tuo
}

// SetChainTxID sets the "chain_tx_id" field.
func (tuo *TranUpdateOne) SetChainTxID(s string) *TranUpdateOne {
	tuo.mutation.SetChainTxID(s)
	return tuo
}

// SetNillableChainTxID sets the "chain_tx_id" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableChainTxID(s *string) *TranUpdateOne {
	if s != nil {
		tuo.SetChainTxID(*s)
	}
	return tuo
}

// ClearChainTxID clears the value of the "chain_tx_id" field.
func (tuo *TranUpdateOne) ClearChainTxID() *TranUpdateOne {
	tuo.mutation.ClearChainTxID()
	return tuo
}

// SetState sets the "state" field.
func (tuo *TranUpdateOne) SetState(s string) *TranUpdateOne {
	tuo.mutation.SetState(s)
	return tuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableState(s *string) *TranUpdateOne {
	if s != nil {
		tuo.SetState(*s)
	}
	return tuo
}

// ClearState clears the value of the "state" field.
func (tuo *TranUpdateOne) ClearState() *TranUpdateOne {
	tuo.mutation.ClearState()
	return tuo
}

// SetExtra sets the "extra" field.
func (tuo *TranUpdateOne) SetExtra(s string) *TranUpdateOne {
	tuo.mutation.SetExtra(s)
	return tuo
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableExtra(s *string) *TranUpdateOne {
	if s != nil {
		tuo.SetExtra(*s)
	}
	return tuo
}

// ClearExtra clears the value of the "extra" field.
func (tuo *TranUpdateOne) ClearExtra() *TranUpdateOne {
	tuo.mutation.ClearExtra()
	return tuo
}

// SetType sets the "type" field.
func (tuo *TranUpdateOne) SetType(s string) *TranUpdateOne {
	tuo.mutation.SetType(s)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TranUpdateOne) SetNillableType(s *string) *TranUpdateOne {
	if s != nil {
		tuo.SetType(*s)
	}
	return tuo
}

// ClearType clears the value of the "type" field.
func (tuo *TranUpdateOne) ClearType() *TranUpdateOne {
	tuo.mutation.ClearType()
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
	if value, ok := tuo.mutation.FromAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldFromAccountID,
		})
	}
	if tuo.mutation.FromAccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldFromAccountID,
		})
	}
	if value, ok := tuo.mutation.ToAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldToAccountID,
		})
	}
	if tuo.mutation.ToAccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: tran.FieldToAccountID,
		})
	}
	if value, ok := tuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldAmount,
		})
	}
	if tuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldAmount,
		})
	}
	if value, ok := tuo.mutation.FeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldFeeAmount,
		})
	}
	if tuo.mutation.FeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: tran.FieldFeeAmount,
		})
	}
	if value, ok := tuo.mutation.ChainTxID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldChainTxID,
		})
	}
	if tuo.mutation.ChainTxIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldChainTxID,
		})
	}
	if value, ok := tuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldState,
		})
	}
	if tuo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldState,
		})
	}
	if value, ok := tuo.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldExtra,
		})
	}
	if tuo.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldExtra,
		})
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tran.FieldType,
		})
	}
	if tuo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tran.FieldType,
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
