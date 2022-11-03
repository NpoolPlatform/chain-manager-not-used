// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/fee"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FeeUpdate is the builder for updating Fee entities.
type FeeUpdate struct {
	config
	hooks     []Hook
	mutation  *FeeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FeeUpdate builder.
func (fu *FeeUpdate) Where(ps ...predicate.Fee) *FeeUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FeeUpdate) SetCreatedAt(u uint32) *FeeUpdate {
	fu.mutation.ResetCreatedAt()
	fu.mutation.SetCreatedAt(u)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableCreatedAt(u *uint32) *FeeUpdate {
	if u != nil {
		fu.SetCreatedAt(*u)
	}
	return fu
}

// AddCreatedAt adds u to the "created_at" field.
func (fu *FeeUpdate) AddCreatedAt(u int32) *FeeUpdate {
	fu.mutation.AddCreatedAt(u)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FeeUpdate) SetUpdatedAt(u uint32) *FeeUpdate {
	fu.mutation.ResetUpdatedAt()
	fu.mutation.SetUpdatedAt(u)
	return fu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fu *FeeUpdate) AddUpdatedAt(u int32) *FeeUpdate {
	fu.mutation.AddUpdatedAt(u)
	return fu
}

// SetDeletedAt sets the "deleted_at" field.
func (fu *FeeUpdate) SetDeletedAt(u uint32) *FeeUpdate {
	fu.mutation.ResetDeletedAt()
	fu.mutation.SetDeletedAt(u)
	return fu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableDeletedAt(u *uint32) *FeeUpdate {
	if u != nil {
		fu.SetDeletedAt(*u)
	}
	return fu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fu *FeeUpdate) AddDeletedAt(u int32) *FeeUpdate {
	fu.mutation.AddDeletedAt(u)
	return fu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (fu *FeeUpdate) SetCoinTypeID(u uuid.UUID) *FeeUpdate {
	fu.mutation.SetCoinTypeID(u)
	return fu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableCoinTypeID(u *uuid.UUID) *FeeUpdate {
	if u != nil {
		fu.SetCoinTypeID(*u)
	}
	return fu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (fu *FeeUpdate) ClearCoinTypeID() *FeeUpdate {
	fu.mutation.ClearCoinTypeID()
	return fu
}

// SetFeeCoinTypeID sets the "fee_coin_type_id" field.
func (fu *FeeUpdate) SetFeeCoinTypeID(u uuid.UUID) *FeeUpdate {
	fu.mutation.SetFeeCoinTypeID(u)
	return fu
}

// SetNillableFeeCoinTypeID sets the "fee_coin_type_id" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableFeeCoinTypeID(u *uuid.UUID) *FeeUpdate {
	if u != nil {
		fu.SetFeeCoinTypeID(*u)
	}
	return fu
}

// ClearFeeCoinTypeID clears the value of the "fee_coin_type_id" field.
func (fu *FeeUpdate) ClearFeeCoinTypeID() *FeeUpdate {
	fu.mutation.ClearFeeCoinTypeID()
	return fu
}

// SetWithdrawFeeByStableUsd sets the "withdraw_fee_by_stable_usd" field.
func (fu *FeeUpdate) SetWithdrawFeeByStableUsd(b bool) *FeeUpdate {
	fu.mutation.SetWithdrawFeeByStableUsd(b)
	return fu
}

// SetNillableWithdrawFeeByStableUsd sets the "withdraw_fee_by_stable_usd" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableWithdrawFeeByStableUsd(b *bool) *FeeUpdate {
	if b != nil {
		fu.SetWithdrawFeeByStableUsd(*b)
	}
	return fu
}

// ClearWithdrawFeeByStableUsd clears the value of the "withdraw_fee_by_stable_usd" field.
func (fu *FeeUpdate) ClearWithdrawFeeByStableUsd() *FeeUpdate {
	fu.mutation.ClearWithdrawFeeByStableUsd()
	return fu
}

// SetWithdrawFeeAmount sets the "withdraw_fee_amount" field.
func (fu *FeeUpdate) SetWithdrawFeeAmount(d decimal.Decimal) *FeeUpdate {
	fu.mutation.SetWithdrawFeeAmount(d)
	return fu
}

// SetNillableWithdrawFeeAmount sets the "withdraw_fee_amount" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableWithdrawFeeAmount(d *decimal.Decimal) *FeeUpdate {
	if d != nil {
		fu.SetWithdrawFeeAmount(*d)
	}
	return fu
}

// ClearWithdrawFeeAmount clears the value of the "withdraw_fee_amount" field.
func (fu *FeeUpdate) ClearWithdrawFeeAmount() *FeeUpdate {
	fu.mutation.ClearWithdrawFeeAmount()
	return fu
}

// SetCollectFeeAmount sets the "collect_fee_amount" field.
func (fu *FeeUpdate) SetCollectFeeAmount(d decimal.Decimal) *FeeUpdate {
	fu.mutation.SetCollectFeeAmount(d)
	return fu
}

// SetNillableCollectFeeAmount sets the "collect_fee_amount" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableCollectFeeAmount(d *decimal.Decimal) *FeeUpdate {
	if d != nil {
		fu.SetCollectFeeAmount(*d)
	}
	return fu
}

// ClearCollectFeeAmount clears the value of the "collect_fee_amount" field.
func (fu *FeeUpdate) ClearCollectFeeAmount() *FeeUpdate {
	fu.mutation.ClearCollectFeeAmount()
	return fu
}

// SetHotWalletFeeAmount sets the "hot_wallet_fee_amount" field.
func (fu *FeeUpdate) SetHotWalletFeeAmount(d decimal.Decimal) *FeeUpdate {
	fu.mutation.SetHotWalletFeeAmount(d)
	return fu
}

// SetNillableHotWalletFeeAmount sets the "hot_wallet_fee_amount" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableHotWalletFeeAmount(d *decimal.Decimal) *FeeUpdate {
	if d != nil {
		fu.SetHotWalletFeeAmount(*d)
	}
	return fu
}

// ClearHotWalletFeeAmount clears the value of the "hot_wallet_fee_amount" field.
func (fu *FeeUpdate) ClearHotWalletFeeAmount() *FeeUpdate {
	fu.mutation.ClearHotWalletFeeAmount()
	return fu
}

// SetLowFeeAmount sets the "low_fee_amount" field.
func (fu *FeeUpdate) SetLowFeeAmount(d decimal.Decimal) *FeeUpdate {
	fu.mutation.SetLowFeeAmount(d)
	return fu
}

// SetNillableLowFeeAmount sets the "low_fee_amount" field if the given value is not nil.
func (fu *FeeUpdate) SetNillableLowFeeAmount(d *decimal.Decimal) *FeeUpdate {
	if d != nil {
		fu.SetLowFeeAmount(*d)
	}
	return fu
}

// ClearLowFeeAmount clears the value of the "low_fee_amount" field.
func (fu *FeeUpdate) ClearLowFeeAmount() *FeeUpdate {
	fu.mutation.ClearLowFeeAmount()
	return fu
}

// Mutation returns the FeeMutation object of the builder.
func (fu *FeeUpdate) Mutation() *FeeMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FeeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := fu.defaults(); err != nil {
		return 0, err
	}
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FeeUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FeeUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FeeUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FeeUpdate) defaults() error {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		if fee.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fee.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fee.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fu *FeeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FeeUpdate {
	fu.modifiers = append(fu.modifiers, modifiers...)
	return fu
}

func (fu *FeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fee.Table,
			Columns: fee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fee.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldDeletedAt,
		})
	}
	if value, ok := fu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldDeletedAt,
		})
	}
	if value, ok := fu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fee.FieldCoinTypeID,
		})
	}
	if fu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fee.FieldCoinTypeID,
		})
	}
	if value, ok := fu.mutation.FeeCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fee.FieldFeeCoinTypeID,
		})
	}
	if fu.mutation.FeeCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fee.FieldFeeCoinTypeID,
		})
	}
	if value, ok := fu.mutation.WithdrawFeeByStableUsd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fee.FieldWithdrawFeeByStableUsd,
		})
	}
	if fu.mutation.WithdrawFeeByStableUsdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: fee.FieldWithdrawFeeByStableUsd,
		})
	}
	if value, ok := fu.mutation.WithdrawFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldWithdrawFeeAmount,
		})
	}
	if fu.mutation.WithdrawFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldWithdrawFeeAmount,
		})
	}
	if value, ok := fu.mutation.CollectFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldCollectFeeAmount,
		})
	}
	if fu.mutation.CollectFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldCollectFeeAmount,
		})
	}
	if value, ok := fu.mutation.HotWalletFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldHotWalletFeeAmount,
		})
	}
	if fu.mutation.HotWalletFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldHotWalletFeeAmount,
		})
	}
	if value, ok := fu.mutation.LowFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldLowFeeAmount,
		})
	}
	if fu.mutation.LowFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldLowFeeAmount,
		})
	}
	_spec.Modifiers = fu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FeeUpdateOne is the builder for updating a single Fee entity.
type FeeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FeeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FeeUpdateOne) SetCreatedAt(u uint32) *FeeUpdateOne {
	fuo.mutation.ResetCreatedAt()
	fuo.mutation.SetCreatedAt(u)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableCreatedAt(u *uint32) *FeeUpdateOne {
	if u != nil {
		fuo.SetCreatedAt(*u)
	}
	return fuo
}

// AddCreatedAt adds u to the "created_at" field.
func (fuo *FeeUpdateOne) AddCreatedAt(u int32) *FeeUpdateOne {
	fuo.mutation.AddCreatedAt(u)
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FeeUpdateOne) SetUpdatedAt(u uint32) *FeeUpdateOne {
	fuo.mutation.ResetUpdatedAt()
	fuo.mutation.SetUpdatedAt(u)
	return fuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fuo *FeeUpdateOne) AddUpdatedAt(u int32) *FeeUpdateOne {
	fuo.mutation.AddUpdatedAt(u)
	return fuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fuo *FeeUpdateOne) SetDeletedAt(u uint32) *FeeUpdateOne {
	fuo.mutation.ResetDeletedAt()
	fuo.mutation.SetDeletedAt(u)
	return fuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableDeletedAt(u *uint32) *FeeUpdateOne {
	if u != nil {
		fuo.SetDeletedAt(*u)
	}
	return fuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fuo *FeeUpdateOne) AddDeletedAt(u int32) *FeeUpdateOne {
	fuo.mutation.AddDeletedAt(u)
	return fuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (fuo *FeeUpdateOne) SetCoinTypeID(u uuid.UUID) *FeeUpdateOne {
	fuo.mutation.SetCoinTypeID(u)
	return fuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *FeeUpdateOne {
	if u != nil {
		fuo.SetCoinTypeID(*u)
	}
	return fuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (fuo *FeeUpdateOne) ClearCoinTypeID() *FeeUpdateOne {
	fuo.mutation.ClearCoinTypeID()
	return fuo
}

// SetFeeCoinTypeID sets the "fee_coin_type_id" field.
func (fuo *FeeUpdateOne) SetFeeCoinTypeID(u uuid.UUID) *FeeUpdateOne {
	fuo.mutation.SetFeeCoinTypeID(u)
	return fuo
}

// SetNillableFeeCoinTypeID sets the "fee_coin_type_id" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableFeeCoinTypeID(u *uuid.UUID) *FeeUpdateOne {
	if u != nil {
		fuo.SetFeeCoinTypeID(*u)
	}
	return fuo
}

// ClearFeeCoinTypeID clears the value of the "fee_coin_type_id" field.
func (fuo *FeeUpdateOne) ClearFeeCoinTypeID() *FeeUpdateOne {
	fuo.mutation.ClearFeeCoinTypeID()
	return fuo
}

// SetWithdrawFeeByStableUsd sets the "withdraw_fee_by_stable_usd" field.
func (fuo *FeeUpdateOne) SetWithdrawFeeByStableUsd(b bool) *FeeUpdateOne {
	fuo.mutation.SetWithdrawFeeByStableUsd(b)
	return fuo
}

// SetNillableWithdrawFeeByStableUsd sets the "withdraw_fee_by_stable_usd" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableWithdrawFeeByStableUsd(b *bool) *FeeUpdateOne {
	if b != nil {
		fuo.SetWithdrawFeeByStableUsd(*b)
	}
	return fuo
}

// ClearWithdrawFeeByStableUsd clears the value of the "withdraw_fee_by_stable_usd" field.
func (fuo *FeeUpdateOne) ClearWithdrawFeeByStableUsd() *FeeUpdateOne {
	fuo.mutation.ClearWithdrawFeeByStableUsd()
	return fuo
}

// SetWithdrawFeeAmount sets the "withdraw_fee_amount" field.
func (fuo *FeeUpdateOne) SetWithdrawFeeAmount(d decimal.Decimal) *FeeUpdateOne {
	fuo.mutation.SetWithdrawFeeAmount(d)
	return fuo
}

// SetNillableWithdrawFeeAmount sets the "withdraw_fee_amount" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableWithdrawFeeAmount(d *decimal.Decimal) *FeeUpdateOne {
	if d != nil {
		fuo.SetWithdrawFeeAmount(*d)
	}
	return fuo
}

// ClearWithdrawFeeAmount clears the value of the "withdraw_fee_amount" field.
func (fuo *FeeUpdateOne) ClearWithdrawFeeAmount() *FeeUpdateOne {
	fuo.mutation.ClearWithdrawFeeAmount()
	return fuo
}

// SetCollectFeeAmount sets the "collect_fee_amount" field.
func (fuo *FeeUpdateOne) SetCollectFeeAmount(d decimal.Decimal) *FeeUpdateOne {
	fuo.mutation.SetCollectFeeAmount(d)
	return fuo
}

// SetNillableCollectFeeAmount sets the "collect_fee_amount" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableCollectFeeAmount(d *decimal.Decimal) *FeeUpdateOne {
	if d != nil {
		fuo.SetCollectFeeAmount(*d)
	}
	return fuo
}

// ClearCollectFeeAmount clears the value of the "collect_fee_amount" field.
func (fuo *FeeUpdateOne) ClearCollectFeeAmount() *FeeUpdateOne {
	fuo.mutation.ClearCollectFeeAmount()
	return fuo
}

// SetHotWalletFeeAmount sets the "hot_wallet_fee_amount" field.
func (fuo *FeeUpdateOne) SetHotWalletFeeAmount(d decimal.Decimal) *FeeUpdateOne {
	fuo.mutation.SetHotWalletFeeAmount(d)
	return fuo
}

// SetNillableHotWalletFeeAmount sets the "hot_wallet_fee_amount" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableHotWalletFeeAmount(d *decimal.Decimal) *FeeUpdateOne {
	if d != nil {
		fuo.SetHotWalletFeeAmount(*d)
	}
	return fuo
}

// ClearHotWalletFeeAmount clears the value of the "hot_wallet_fee_amount" field.
func (fuo *FeeUpdateOne) ClearHotWalletFeeAmount() *FeeUpdateOne {
	fuo.mutation.ClearHotWalletFeeAmount()
	return fuo
}

// SetLowFeeAmount sets the "low_fee_amount" field.
func (fuo *FeeUpdateOne) SetLowFeeAmount(d decimal.Decimal) *FeeUpdateOne {
	fuo.mutation.SetLowFeeAmount(d)
	return fuo
}

// SetNillableLowFeeAmount sets the "low_fee_amount" field if the given value is not nil.
func (fuo *FeeUpdateOne) SetNillableLowFeeAmount(d *decimal.Decimal) *FeeUpdateOne {
	if d != nil {
		fuo.SetLowFeeAmount(*d)
	}
	return fuo
}

// ClearLowFeeAmount clears the value of the "low_fee_amount" field.
func (fuo *FeeUpdateOne) ClearLowFeeAmount() *FeeUpdateOne {
	fuo.mutation.ClearLowFeeAmount()
	return fuo
}

// Mutation returns the FeeMutation object of the builder.
func (fuo *FeeUpdateOne) Mutation() *FeeMutation {
	return fuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FeeUpdateOne) Select(field string, fields ...string) *FeeUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Fee entity.
func (fuo *FeeUpdateOne) Save(ctx context.Context) (*Fee, error) {
	var (
		err  error
		node *Fee
	)
	if err := fuo.defaults(); err != nil {
		return nil, err
	}
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Fee)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FeeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FeeUpdateOne) SaveX(ctx context.Context) *Fee {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FeeUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FeeUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FeeUpdateOne) defaults() error {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		if fee.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fee.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fee.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fuo *FeeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FeeUpdateOne {
	fuo.modifiers = append(fuo.modifiers, modifiers...)
	return fuo
}

func (fuo *FeeUpdateOne) sqlSave(ctx context.Context) (_node *Fee, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fee.Table,
			Columns: fee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fee.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Fee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fee.FieldID)
		for _, f := range fields {
			if !fee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldDeletedAt,
		})
	}
	if value, ok := fuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fee.FieldDeletedAt,
		})
	}
	if value, ok := fuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fee.FieldCoinTypeID,
		})
	}
	if fuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fee.FieldCoinTypeID,
		})
	}
	if value, ok := fuo.mutation.FeeCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fee.FieldFeeCoinTypeID,
		})
	}
	if fuo.mutation.FeeCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fee.FieldFeeCoinTypeID,
		})
	}
	if value, ok := fuo.mutation.WithdrawFeeByStableUsd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fee.FieldWithdrawFeeByStableUsd,
		})
	}
	if fuo.mutation.WithdrawFeeByStableUsdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: fee.FieldWithdrawFeeByStableUsd,
		})
	}
	if value, ok := fuo.mutation.WithdrawFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldWithdrawFeeAmount,
		})
	}
	if fuo.mutation.WithdrawFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldWithdrawFeeAmount,
		})
	}
	if value, ok := fuo.mutation.CollectFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldCollectFeeAmount,
		})
	}
	if fuo.mutation.CollectFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldCollectFeeAmount,
		})
	}
	if value, ok := fuo.mutation.HotWalletFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldHotWalletFeeAmount,
		})
	}
	if fuo.mutation.HotWalletFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldHotWalletFeeAmount,
		})
	}
	if value, ok := fuo.mutation.LowFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fee.FieldLowFeeAmount,
		})
	}
	if fuo.mutation.LowFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fee.FieldLowFeeAmount,
		})
	}
	_spec.Modifiers = fuo.modifiers
	_node = &Fee{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}