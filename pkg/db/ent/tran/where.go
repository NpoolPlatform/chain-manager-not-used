// Code generated by ent, DO NOT EDIT.

package tran

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// FromAccountID applies equality check predicate on the "from_account_id" field. It's identical to FromAccountIDEQ.
func FromAccountID(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromAccountID), v))
	})
}

// ToAccountID applies equality check predicate on the "to_account_id" field. It's identical to ToAccountIDEQ.
func ToAccountID(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToAccountID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// FeeAmount applies equality check predicate on the "fee_amount" field. It's identical to FeeAmountEQ.
func FeeAmount(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeeAmount), v))
	})
}

// ChainTxID applies equality check predicate on the "chain_tx_id" field. It's identical to ChainTxIDEQ.
func ChainTxID(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainTxID), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// Extra applies equality check predicate on the "extra" field. It's identical to ExtraEQ.
func Extra(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtra), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// FromAccountIDEQ applies the EQ predicate on the "from_account_id" field.
func FromAccountIDEQ(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDNEQ applies the NEQ predicate on the "from_account_id" field.
func FromAccountIDNEQ(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDIn applies the In predicate on the "from_account_id" field.
func FromAccountIDIn(vs ...uuid.UUID) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFromAccountID), v...))
	})
}

// FromAccountIDNotIn applies the NotIn predicate on the "from_account_id" field.
func FromAccountIDNotIn(vs ...uuid.UUID) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFromAccountID), v...))
	})
}

// FromAccountIDGT applies the GT predicate on the "from_account_id" field.
func FromAccountIDGT(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDGTE applies the GTE predicate on the "from_account_id" field.
func FromAccountIDGTE(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDLT applies the LT predicate on the "from_account_id" field.
func FromAccountIDLT(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDLTE applies the LTE predicate on the "from_account_id" field.
func FromAccountIDLTE(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDIsNil applies the IsNil predicate on the "from_account_id" field.
func FromAccountIDIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFromAccountID)))
	})
}

// FromAccountIDNotNil applies the NotNil predicate on the "from_account_id" field.
func FromAccountIDNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFromAccountID)))
	})
}

// ToAccountIDEQ applies the EQ predicate on the "to_account_id" field.
func ToAccountIDEQ(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToAccountID), v))
	})
}

// ToAccountIDNEQ applies the NEQ predicate on the "to_account_id" field.
func ToAccountIDNEQ(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToAccountID), v))
	})
}

// ToAccountIDIn applies the In predicate on the "to_account_id" field.
func ToAccountIDIn(vs ...uuid.UUID) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToAccountID), v...))
	})
}

// ToAccountIDNotIn applies the NotIn predicate on the "to_account_id" field.
func ToAccountIDNotIn(vs ...uuid.UUID) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToAccountID), v...))
	})
}

// ToAccountIDGT applies the GT predicate on the "to_account_id" field.
func ToAccountIDGT(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToAccountID), v))
	})
}

// ToAccountIDGTE applies the GTE predicate on the "to_account_id" field.
func ToAccountIDGTE(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToAccountID), v))
	})
}

// ToAccountIDLT applies the LT predicate on the "to_account_id" field.
func ToAccountIDLT(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToAccountID), v))
	})
}

// ToAccountIDLTE applies the LTE predicate on the "to_account_id" field.
func ToAccountIDLTE(v uuid.UUID) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToAccountID), v))
	})
}

// ToAccountIDIsNil applies the IsNil predicate on the "to_account_id" field.
func ToAccountIDIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldToAccountID)))
	})
}

// ToAccountIDNotNil applies the NotNil predicate on the "to_account_id" field.
func ToAccountIDNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldToAccountID)))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...decimal.Decimal) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...decimal.Decimal) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// FeeAmountEQ applies the EQ predicate on the "fee_amount" field.
func FeeAmountEQ(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeeAmount), v))
	})
}

// FeeAmountNEQ applies the NEQ predicate on the "fee_amount" field.
func FeeAmountNEQ(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeeAmount), v))
	})
}

// FeeAmountIn applies the In predicate on the "fee_amount" field.
func FeeAmountIn(vs ...decimal.Decimal) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeeAmount), v...))
	})
}

// FeeAmountNotIn applies the NotIn predicate on the "fee_amount" field.
func FeeAmountNotIn(vs ...decimal.Decimal) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeeAmount), v...))
	})
}

// FeeAmountGT applies the GT predicate on the "fee_amount" field.
func FeeAmountGT(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeeAmount), v))
	})
}

// FeeAmountGTE applies the GTE predicate on the "fee_amount" field.
func FeeAmountGTE(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeeAmount), v))
	})
}

// FeeAmountLT applies the LT predicate on the "fee_amount" field.
func FeeAmountLT(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeeAmount), v))
	})
}

// FeeAmountLTE applies the LTE predicate on the "fee_amount" field.
func FeeAmountLTE(v decimal.Decimal) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeeAmount), v))
	})
}

// FeeAmountIsNil applies the IsNil predicate on the "fee_amount" field.
func FeeAmountIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFeeAmount)))
	})
}

// FeeAmountNotNil applies the NotNil predicate on the "fee_amount" field.
func FeeAmountNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFeeAmount)))
	})
}

// ChainTxIDEQ applies the EQ predicate on the "chain_tx_id" field.
func ChainTxIDEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDNEQ applies the NEQ predicate on the "chain_tx_id" field.
func ChainTxIDNEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDIn applies the In predicate on the "chain_tx_id" field.
func ChainTxIDIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChainTxID), v...))
	})
}

// ChainTxIDNotIn applies the NotIn predicate on the "chain_tx_id" field.
func ChainTxIDNotIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChainTxID), v...))
	})
}

// ChainTxIDGT applies the GT predicate on the "chain_tx_id" field.
func ChainTxIDGT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDGTE applies the GTE predicate on the "chain_tx_id" field.
func ChainTxIDGTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDLT applies the LT predicate on the "chain_tx_id" field.
func ChainTxIDLT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDLTE applies the LTE predicate on the "chain_tx_id" field.
func ChainTxIDLTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDContains applies the Contains predicate on the "chain_tx_id" field.
func ChainTxIDContains(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDHasPrefix applies the HasPrefix predicate on the "chain_tx_id" field.
func ChainTxIDHasPrefix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDHasSuffix applies the HasSuffix predicate on the "chain_tx_id" field.
func ChainTxIDHasSuffix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDIsNil applies the IsNil predicate on the "chain_tx_id" field.
func ChainTxIDIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChainTxID)))
	})
}

// ChainTxIDNotNil applies the NotNil predicate on the "chain_tx_id" field.
func ChainTxIDNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChainTxID)))
	})
}

// ChainTxIDEqualFold applies the EqualFold predicate on the "chain_tx_id" field.
func ChainTxIDEqualFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChainTxID), v))
	})
}

// ChainTxIDContainsFold applies the ContainsFold predicate on the "chain_tx_id" field.
func ChainTxIDContainsFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChainTxID), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// ExtraEQ applies the EQ predicate on the "extra" field.
func ExtraEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtra), v))
	})
}

// ExtraNEQ applies the NEQ predicate on the "extra" field.
func ExtraNEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExtra), v))
	})
}

// ExtraIn applies the In predicate on the "extra" field.
func ExtraIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExtra), v...))
	})
}

// ExtraNotIn applies the NotIn predicate on the "extra" field.
func ExtraNotIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExtra), v...))
	})
}

// ExtraGT applies the GT predicate on the "extra" field.
func ExtraGT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExtra), v))
	})
}

// ExtraGTE applies the GTE predicate on the "extra" field.
func ExtraGTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExtra), v))
	})
}

// ExtraLT applies the LT predicate on the "extra" field.
func ExtraLT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExtra), v))
	})
}

// ExtraLTE applies the LTE predicate on the "extra" field.
func ExtraLTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExtra), v))
	})
}

// ExtraContains applies the Contains predicate on the "extra" field.
func ExtraContains(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExtra), v))
	})
}

// ExtraHasPrefix applies the HasPrefix predicate on the "extra" field.
func ExtraHasPrefix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExtra), v))
	})
}

// ExtraHasSuffix applies the HasSuffix predicate on the "extra" field.
func ExtraHasSuffix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExtra), v))
	})
}

// ExtraIsNil applies the IsNil predicate on the "extra" field.
func ExtraIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExtra)))
	})
}

// ExtraNotNil applies the NotNil predicate on the "extra" field.
func ExtraNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExtra)))
	})
}

// ExtraEqualFold applies the EqualFold predicate on the "extra" field.
func ExtraEqualFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExtra), v))
	})
}

// ExtraContainsFold applies the ContainsFold predicate on the "extra" field.
func ExtraContainsFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExtra), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Tran {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldType)))
	})
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldType)))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tran) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tran) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tran) predicate.Tran {
	return predicate.Tran(func(s *sql.Selector) {
		p(s.Not())
	})
}
