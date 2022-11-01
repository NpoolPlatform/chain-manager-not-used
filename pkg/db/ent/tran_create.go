// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TranCreate is the builder for creating a Tran entity.
type TranCreate struct {
	config
	mutation *TranMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TranCreate) SetCreatedAt(u uint32) *TranCreate {
	tc.mutation.SetCreatedAt(u)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TranCreate) SetNillableCreatedAt(u *uint32) *TranCreate {
	if u != nil {
		tc.SetCreatedAt(*u)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TranCreate) SetUpdatedAt(u uint32) *TranCreate {
	tc.mutation.SetUpdatedAt(u)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TranCreate) SetNillableUpdatedAt(u *uint32) *TranCreate {
	if u != nil {
		tc.SetUpdatedAt(*u)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TranCreate) SetDeletedAt(u uint32) *TranCreate {
	tc.mutation.SetDeletedAt(u)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TranCreate) SetNillableDeletedAt(u *uint32) *TranCreate {
	if u != nil {
		tc.SetDeletedAt(*u)
	}
	return tc
}

// SetAppID sets the "app_id" field.
func (tc *TranCreate) SetAppID(u uuid.UUID) *TranCreate {
	tc.mutation.SetAppID(u)
	return tc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (tc *TranCreate) SetNillableAppID(u *uuid.UUID) *TranCreate {
	if u != nil {
		tc.SetAppID(*u)
	}
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TranCreate) SetUserID(u uuid.UUID) *TranCreate {
	tc.mutation.SetUserID(u)
	return tc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tc *TranCreate) SetNillableUserID(u *uuid.UUID) *TranCreate {
	if u != nil {
		tc.SetUserID(*u)
	}
	return tc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (tc *TranCreate) SetCoinTypeID(u uuid.UUID) *TranCreate {
	tc.mutation.SetCoinTypeID(u)
	return tc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (tc *TranCreate) SetNillableCoinTypeID(u *uuid.UUID) *TranCreate {
	if u != nil {
		tc.SetCoinTypeID(*u)
	}
	return tc
}

// SetIncoming sets the "incoming" field.
func (tc *TranCreate) SetIncoming(d decimal.Decimal) *TranCreate {
	tc.mutation.SetIncoming(d)
	return tc
}

// SetNillableIncoming sets the "incoming" field if the given value is not nil.
func (tc *TranCreate) SetNillableIncoming(d *decimal.Decimal) *TranCreate {
	if d != nil {
		tc.SetIncoming(*d)
	}
	return tc
}

// SetLocked sets the "locked" field.
func (tc *TranCreate) SetLocked(d decimal.Decimal) *TranCreate {
	tc.mutation.SetLocked(d)
	return tc
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (tc *TranCreate) SetNillableLocked(d *decimal.Decimal) *TranCreate {
	if d != nil {
		tc.SetLocked(*d)
	}
	return tc
}

// SetOutcoming sets the "outcoming" field.
func (tc *TranCreate) SetOutcoming(d decimal.Decimal) *TranCreate {
	tc.mutation.SetOutcoming(d)
	return tc
}

// SetNillableOutcoming sets the "outcoming" field if the given value is not nil.
func (tc *TranCreate) SetNillableOutcoming(d *decimal.Decimal) *TranCreate {
	if d != nil {
		tc.SetOutcoming(*d)
	}
	return tc
}

// SetSpendable sets the "spendable" field.
func (tc *TranCreate) SetSpendable(d decimal.Decimal) *TranCreate {
	tc.mutation.SetSpendable(d)
	return tc
}

// SetNillableSpendable sets the "spendable" field if the given value is not nil.
func (tc *TranCreate) SetNillableSpendable(d *decimal.Decimal) *TranCreate {
	if d != nil {
		tc.SetSpendable(*d)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TranCreate) SetID(u uuid.UUID) *TranCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TranCreate) SetNillableID(u *uuid.UUID) *TranCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TranMutation object of the builder.
func (tc *TranCreate) Mutation() *TranMutation {
	return tc.mutation
}

// Save creates the Tran in the database.
func (tc *TranCreate) Save(ctx context.Context) (*Tran, error) {
	var (
		err  error
		node *Tran
	)
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (tc *TranCreate) SaveX(ctx context.Context) *Tran {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TranCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TranCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TranCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if tran.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := tran.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if tran.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tran.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		if tran.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := tran.DefaultDeletedAt()
		tc.mutation.SetDeletedAt(v)
	}
	if _, ok := tc.mutation.AppID(); !ok {
		if tran.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := tran.DefaultAppID()
		tc.mutation.SetAppID(v)
	}
	if _, ok := tc.mutation.UserID(); !ok {
		if tran.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := tran.DefaultUserID()
		tc.mutation.SetUserID(v)
	}
	if _, ok := tc.mutation.CoinTypeID(); !ok {
		if tran.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := tran.DefaultCoinTypeID()
		tc.mutation.SetCoinTypeID(v)
	}
	if _, ok := tc.mutation.Incoming(); !ok {
		v := tran.DefaultIncoming
		tc.mutation.SetIncoming(v)
	}
	if _, ok := tc.mutation.Locked(); !ok {
		v := tran.DefaultLocked
		tc.mutation.SetLocked(v)
	}
	if _, ok := tc.mutation.Outcoming(); !ok {
		v := tran.DefaultOutcoming
		tc.mutation.SetOutcoming(v)
	}
	if _, ok := tc.mutation.Spendable(); !ok {
		v := tran.DefaultSpendable
		tc.mutation.SetSpendable(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		if tran.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized tran.DefaultID (forgotten import ent/runtime?)")
		}
		v := tran.DefaultID()
		tc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TranCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tran.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Tran.updated_at"`)}
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Tran.deleted_at"`)}
	}
	return nil
}

func (tc *TranCreate) sqlSave(ctx context.Context) (*Tran, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TranCreate) createSpec() (*Tran, *sqlgraph.CreateSpec) {
	var (
		_node = &Tran{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tran.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tran.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tran.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := tc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := tc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tran.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := tc.mutation.Incoming(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldIncoming,
		})
		_node.Incoming = value
	}
	if value, ok := tc.mutation.Locked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldLocked,
		})
		_node.Locked = value
	}
	if value, ok := tc.mutation.Outcoming(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldOutcoming,
		})
		_node.Outcoming = value
	}
	if value, ok := tc.mutation.Spendable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: tran.FieldSpendable,
		})
		_node.Spendable = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tran.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TranUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TranCreate) OnConflict(opts ...sql.ConflictOption) *TranUpsertOne {
	tc.conflict = opts
	return &TranUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tran.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TranCreate) OnConflictColumns(columns ...string) *TranUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TranUpsertOne{
		create: tc,
	}
}

type (
	// TranUpsertOne is the builder for "upsert"-ing
	//  one Tran node.
	TranUpsertOne struct {
		create *TranCreate
	}

	// TranUpsert is the "OnConflict" setter.
	TranUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *TranUpsert) SetCreatedAt(v uint32) *TranUpsert {
	u.Set(tran.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TranUpsert) UpdateCreatedAt() *TranUpsert {
	u.SetExcluded(tran.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TranUpsert) AddCreatedAt(v uint32) *TranUpsert {
	u.Add(tran.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TranUpsert) SetUpdatedAt(v uint32) *TranUpsert {
	u.Set(tran.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TranUpsert) UpdateUpdatedAt() *TranUpsert {
	u.SetExcluded(tran.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TranUpsert) AddUpdatedAt(v uint32) *TranUpsert {
	u.Add(tran.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TranUpsert) SetDeletedAt(v uint32) *TranUpsert {
	u.Set(tran.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TranUpsert) UpdateDeletedAt() *TranUpsert {
	u.SetExcluded(tran.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TranUpsert) AddDeletedAt(v uint32) *TranUpsert {
	u.Add(tran.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *TranUpsert) SetAppID(v uuid.UUID) *TranUpsert {
	u.Set(tran.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *TranUpsert) UpdateAppID() *TranUpsert {
	u.SetExcluded(tran.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *TranUpsert) ClearAppID() *TranUpsert {
	u.SetNull(tran.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *TranUpsert) SetUserID(v uuid.UUID) *TranUpsert {
	u.Set(tran.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TranUpsert) UpdateUserID() *TranUpsert {
	u.SetExcluded(tran.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *TranUpsert) ClearUserID() *TranUpsert {
	u.SetNull(tran.FieldUserID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *TranUpsert) SetCoinTypeID(v uuid.UUID) *TranUpsert {
	u.Set(tran.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *TranUpsert) UpdateCoinTypeID() *TranUpsert {
	u.SetExcluded(tran.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *TranUpsert) ClearCoinTypeID() *TranUpsert {
	u.SetNull(tran.FieldCoinTypeID)
	return u
}

// SetIncoming sets the "incoming" field.
func (u *TranUpsert) SetIncoming(v decimal.Decimal) *TranUpsert {
	u.Set(tran.FieldIncoming, v)
	return u
}

// UpdateIncoming sets the "incoming" field to the value that was provided on create.
func (u *TranUpsert) UpdateIncoming() *TranUpsert {
	u.SetExcluded(tran.FieldIncoming)
	return u
}

// ClearIncoming clears the value of the "incoming" field.
func (u *TranUpsert) ClearIncoming() *TranUpsert {
	u.SetNull(tran.FieldIncoming)
	return u
}

// SetLocked sets the "locked" field.
func (u *TranUpsert) SetLocked(v decimal.Decimal) *TranUpsert {
	u.Set(tran.FieldLocked, v)
	return u
}

// UpdateLocked sets the "locked" field to the value that was provided on create.
func (u *TranUpsert) UpdateLocked() *TranUpsert {
	u.SetExcluded(tran.FieldLocked)
	return u
}

// ClearLocked clears the value of the "locked" field.
func (u *TranUpsert) ClearLocked() *TranUpsert {
	u.SetNull(tran.FieldLocked)
	return u
}

// SetOutcoming sets the "outcoming" field.
func (u *TranUpsert) SetOutcoming(v decimal.Decimal) *TranUpsert {
	u.Set(tran.FieldOutcoming, v)
	return u
}

// UpdateOutcoming sets the "outcoming" field to the value that was provided on create.
func (u *TranUpsert) UpdateOutcoming() *TranUpsert {
	u.SetExcluded(tran.FieldOutcoming)
	return u
}

// ClearOutcoming clears the value of the "outcoming" field.
func (u *TranUpsert) ClearOutcoming() *TranUpsert {
	u.SetNull(tran.FieldOutcoming)
	return u
}

// SetSpendable sets the "spendable" field.
func (u *TranUpsert) SetSpendable(v decimal.Decimal) *TranUpsert {
	u.Set(tran.FieldSpendable, v)
	return u
}

// UpdateSpendable sets the "spendable" field to the value that was provided on create.
func (u *TranUpsert) UpdateSpendable() *TranUpsert {
	u.SetExcluded(tran.FieldSpendable)
	return u
}

// ClearSpendable clears the value of the "spendable" field.
func (u *TranUpsert) ClearSpendable() *TranUpsert {
	u.SetNull(tran.FieldSpendable)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Tran.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tran.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TranUpsertOne) UpdateNewValues() *TranUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tran.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Tran.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TranUpsertOne) Ignore() *TranUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TranUpsertOne) DoNothing() *TranUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TranCreate.OnConflict
// documentation for more info.
func (u *TranUpsertOne) Update(set func(*TranUpsert)) *TranUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TranUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TranUpsertOne) SetCreatedAt(v uint32) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TranUpsertOne) AddCreatedAt(v uint32) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateCreatedAt() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TranUpsertOne) SetUpdatedAt(v uint32) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TranUpsertOne) AddUpdatedAt(v uint32) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateUpdatedAt() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TranUpsertOne) SetDeletedAt(v uint32) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TranUpsertOne) AddDeletedAt(v uint32) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateDeletedAt() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *TranUpsertOne) SetAppID(v uuid.UUID) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateAppID() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *TranUpsertOne) ClearAppID() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *TranUpsertOne) SetUserID(v uuid.UUID) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateUserID() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *TranUpsertOne) ClearUserID() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *TranUpsertOne) SetCoinTypeID(v uuid.UUID) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateCoinTypeID() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *TranUpsertOne) ClearCoinTypeID() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetIncoming sets the "incoming" field.
func (u *TranUpsertOne) SetIncoming(v decimal.Decimal) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetIncoming(v)
	})
}

// UpdateIncoming sets the "incoming" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateIncoming() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateIncoming()
	})
}

// ClearIncoming clears the value of the "incoming" field.
func (u *TranUpsertOne) ClearIncoming() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearIncoming()
	})
}

// SetLocked sets the "locked" field.
func (u *TranUpsertOne) SetLocked(v decimal.Decimal) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetLocked(v)
	})
}

// UpdateLocked sets the "locked" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateLocked() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateLocked()
	})
}

// ClearLocked clears the value of the "locked" field.
func (u *TranUpsertOne) ClearLocked() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearLocked()
	})
}

// SetOutcoming sets the "outcoming" field.
func (u *TranUpsertOne) SetOutcoming(v decimal.Decimal) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetOutcoming(v)
	})
}

// UpdateOutcoming sets the "outcoming" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateOutcoming() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateOutcoming()
	})
}

// ClearOutcoming clears the value of the "outcoming" field.
func (u *TranUpsertOne) ClearOutcoming() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearOutcoming()
	})
}

// SetSpendable sets the "spendable" field.
func (u *TranUpsertOne) SetSpendable(v decimal.Decimal) *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.SetSpendable(v)
	})
}

// UpdateSpendable sets the "spendable" field to the value that was provided on create.
func (u *TranUpsertOne) UpdateSpendable() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.UpdateSpendable()
	})
}

// ClearSpendable clears the value of the "spendable" field.
func (u *TranUpsertOne) ClearSpendable() *TranUpsertOne {
	return u.Update(func(s *TranUpsert) {
		s.ClearSpendable()
	})
}

// Exec executes the query.
func (u *TranUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TranCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TranUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TranUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TranUpsertOne.ID is not supported by MySQL driver. Use TranUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TranUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TranCreateBulk is the builder for creating many Tran entities in bulk.
type TranCreateBulk struct {
	config
	builders []*TranCreate
	conflict []sql.ConflictOption
}

// Save creates the Tran entities in the database.
func (tcb *TranCreateBulk) Save(ctx context.Context) ([]*Tran, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tran, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TranMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TranCreateBulk) SaveX(ctx context.Context) []*Tran {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TranCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TranCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tran.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TranUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TranCreateBulk) OnConflict(opts ...sql.ConflictOption) *TranUpsertBulk {
	tcb.conflict = opts
	return &TranUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tran.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TranCreateBulk) OnConflictColumns(columns ...string) *TranUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TranUpsertBulk{
		create: tcb,
	}
}

// TranUpsertBulk is the builder for "upsert"-ing
// a bulk of Tran nodes.
type TranUpsertBulk struct {
	create *TranCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tran.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tran.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TranUpsertBulk) UpdateNewValues() *TranUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tran.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tran.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TranUpsertBulk) Ignore() *TranUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TranUpsertBulk) DoNothing() *TranUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TranCreateBulk.OnConflict
// documentation for more info.
func (u *TranUpsertBulk) Update(set func(*TranUpsert)) *TranUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TranUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TranUpsertBulk) SetCreatedAt(v uint32) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TranUpsertBulk) AddCreatedAt(v uint32) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateCreatedAt() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TranUpsertBulk) SetUpdatedAt(v uint32) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TranUpsertBulk) AddUpdatedAt(v uint32) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateUpdatedAt() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TranUpsertBulk) SetDeletedAt(v uint32) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TranUpsertBulk) AddDeletedAt(v uint32) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateDeletedAt() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *TranUpsertBulk) SetAppID(v uuid.UUID) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateAppID() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *TranUpsertBulk) ClearAppID() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *TranUpsertBulk) SetUserID(v uuid.UUID) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateUserID() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *TranUpsertBulk) ClearUserID() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *TranUpsertBulk) SetCoinTypeID(v uuid.UUID) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateCoinTypeID() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *TranUpsertBulk) ClearCoinTypeID() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetIncoming sets the "incoming" field.
func (u *TranUpsertBulk) SetIncoming(v decimal.Decimal) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetIncoming(v)
	})
}

// UpdateIncoming sets the "incoming" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateIncoming() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateIncoming()
	})
}

// ClearIncoming clears the value of the "incoming" field.
func (u *TranUpsertBulk) ClearIncoming() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearIncoming()
	})
}

// SetLocked sets the "locked" field.
func (u *TranUpsertBulk) SetLocked(v decimal.Decimal) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetLocked(v)
	})
}

// UpdateLocked sets the "locked" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateLocked() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateLocked()
	})
}

// ClearLocked clears the value of the "locked" field.
func (u *TranUpsertBulk) ClearLocked() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearLocked()
	})
}

// SetOutcoming sets the "outcoming" field.
func (u *TranUpsertBulk) SetOutcoming(v decimal.Decimal) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetOutcoming(v)
	})
}

// UpdateOutcoming sets the "outcoming" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateOutcoming() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateOutcoming()
	})
}

// ClearOutcoming clears the value of the "outcoming" field.
func (u *TranUpsertBulk) ClearOutcoming() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearOutcoming()
	})
}

// SetSpendable sets the "spendable" field.
func (u *TranUpsertBulk) SetSpendable(v decimal.Decimal) *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.SetSpendable(v)
	})
}

// UpdateSpendable sets the "spendable" field to the value that was provided on create.
func (u *TranUpsertBulk) UpdateSpendable() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.UpdateSpendable()
	})
}

// ClearSpendable clears the value of the "spendable" field.
func (u *TranUpsertBulk) ClearSpendable() *TranUpsertBulk {
	return u.Update(func(s *TranUpsert) {
		s.ClearSpendable()
	})
}

// Exec executes the query.
func (u *TranUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TranCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TranCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TranUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}