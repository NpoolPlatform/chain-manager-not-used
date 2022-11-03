// Code generated by ent, DO NOT EDIT.

package fee

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// FeeCoinTypeID applies equality check predicate on the "fee_coin_type_id" field. It's identical to FeeCoinTypeIDEQ.
func FeeCoinTypeID(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeeCoinTypeID), v))
	})
}

// WithdrawFeeByStableUsd applies equality check predicate on the "withdraw_fee_by_stable_usd" field. It's identical to WithdrawFeeByStableUsdEQ.
func WithdrawFeeByStableUsd(v bool) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawFeeByStableUsd), v))
	})
}

// WithdrawFeeAmount applies equality check predicate on the "withdraw_fee_amount" field. It's identical to WithdrawFeeAmountEQ.
func WithdrawFeeAmount(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawFeeAmount), v))
	})
}

// CollectFeeAmount applies equality check predicate on the "collect_fee_amount" field. It's identical to CollectFeeAmountEQ.
func CollectFeeAmount(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCollectFeeAmount), v))
	})
}

// HotWalletFeeAmount applies equality check predicate on the "hot_wallet_fee_amount" field. It's identical to HotWalletFeeAmountEQ.
func HotWalletFeeAmount(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHotWalletFeeAmount), v))
	})
}

// LowFeeAmount applies equality check predicate on the "low_fee_amount" field. It's identical to LowFeeAmountEQ.
func LowFeeAmount(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLowFeeAmount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// FeeCoinTypeIDEQ applies the EQ predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDEQ(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeeCoinTypeID), v))
	})
}

// FeeCoinTypeIDNEQ applies the NEQ predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDNEQ(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeeCoinTypeID), v))
	})
}

// FeeCoinTypeIDIn applies the In predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDIn(vs ...uuid.UUID) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeeCoinTypeID), v...))
	})
}

// FeeCoinTypeIDNotIn applies the NotIn predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDNotIn(vs ...uuid.UUID) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeeCoinTypeID), v...))
	})
}

// FeeCoinTypeIDGT applies the GT predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDGT(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeeCoinTypeID), v))
	})
}

// FeeCoinTypeIDGTE applies the GTE predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDGTE(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeeCoinTypeID), v))
	})
}

// FeeCoinTypeIDLT applies the LT predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDLT(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeeCoinTypeID), v))
	})
}

// FeeCoinTypeIDLTE applies the LTE predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDLTE(v uuid.UUID) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeeCoinTypeID), v))
	})
}

// FeeCoinTypeIDIsNil applies the IsNil predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFeeCoinTypeID)))
	})
}

// FeeCoinTypeIDNotNil applies the NotNil predicate on the "fee_coin_type_id" field.
func FeeCoinTypeIDNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFeeCoinTypeID)))
	})
}

// WithdrawFeeByStableUsdEQ applies the EQ predicate on the "withdraw_fee_by_stable_usd" field.
func WithdrawFeeByStableUsdEQ(v bool) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawFeeByStableUsd), v))
	})
}

// WithdrawFeeByStableUsdNEQ applies the NEQ predicate on the "withdraw_fee_by_stable_usd" field.
func WithdrawFeeByStableUsdNEQ(v bool) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawFeeByStableUsd), v))
	})
}

// WithdrawFeeByStableUsdIsNil applies the IsNil predicate on the "withdraw_fee_by_stable_usd" field.
func WithdrawFeeByStableUsdIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWithdrawFeeByStableUsd)))
	})
}

// WithdrawFeeByStableUsdNotNil applies the NotNil predicate on the "withdraw_fee_by_stable_usd" field.
func WithdrawFeeByStableUsdNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWithdrawFeeByStableUsd)))
	})
}

// WithdrawFeeAmountEQ applies the EQ predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawFeeAmount), v))
	})
}

// WithdrawFeeAmountNEQ applies the NEQ predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountNEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawFeeAmount), v))
	})
}

// WithdrawFeeAmountIn applies the In predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWithdrawFeeAmount), v...))
	})
}

// WithdrawFeeAmountNotIn applies the NotIn predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountNotIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWithdrawFeeAmount), v...))
	})
}

// WithdrawFeeAmountGT applies the GT predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountGT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawFeeAmount), v))
	})
}

// WithdrawFeeAmountGTE applies the GTE predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountGTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawFeeAmount), v))
	})
}

// WithdrawFeeAmountLT applies the LT predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountLT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawFeeAmount), v))
	})
}

// WithdrawFeeAmountLTE applies the LTE predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountLTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawFeeAmount), v))
	})
}

// WithdrawFeeAmountIsNil applies the IsNil predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWithdrawFeeAmount)))
	})
}

// WithdrawFeeAmountNotNil applies the NotNil predicate on the "withdraw_fee_amount" field.
func WithdrawFeeAmountNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWithdrawFeeAmount)))
	})
}

// CollectFeeAmountEQ applies the EQ predicate on the "collect_fee_amount" field.
func CollectFeeAmountEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCollectFeeAmount), v))
	})
}

// CollectFeeAmountNEQ applies the NEQ predicate on the "collect_fee_amount" field.
func CollectFeeAmountNEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCollectFeeAmount), v))
	})
}

// CollectFeeAmountIn applies the In predicate on the "collect_fee_amount" field.
func CollectFeeAmountIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCollectFeeAmount), v...))
	})
}

// CollectFeeAmountNotIn applies the NotIn predicate on the "collect_fee_amount" field.
func CollectFeeAmountNotIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCollectFeeAmount), v...))
	})
}

// CollectFeeAmountGT applies the GT predicate on the "collect_fee_amount" field.
func CollectFeeAmountGT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCollectFeeAmount), v))
	})
}

// CollectFeeAmountGTE applies the GTE predicate on the "collect_fee_amount" field.
func CollectFeeAmountGTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCollectFeeAmount), v))
	})
}

// CollectFeeAmountLT applies the LT predicate on the "collect_fee_amount" field.
func CollectFeeAmountLT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCollectFeeAmount), v))
	})
}

// CollectFeeAmountLTE applies the LTE predicate on the "collect_fee_amount" field.
func CollectFeeAmountLTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCollectFeeAmount), v))
	})
}

// CollectFeeAmountIsNil applies the IsNil predicate on the "collect_fee_amount" field.
func CollectFeeAmountIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCollectFeeAmount)))
	})
}

// CollectFeeAmountNotNil applies the NotNil predicate on the "collect_fee_amount" field.
func CollectFeeAmountNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCollectFeeAmount)))
	})
}

// HotWalletFeeAmountEQ applies the EQ predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHotWalletFeeAmount), v))
	})
}

// HotWalletFeeAmountNEQ applies the NEQ predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountNEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHotWalletFeeAmount), v))
	})
}

// HotWalletFeeAmountIn applies the In predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHotWalletFeeAmount), v...))
	})
}

// HotWalletFeeAmountNotIn applies the NotIn predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountNotIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHotWalletFeeAmount), v...))
	})
}

// HotWalletFeeAmountGT applies the GT predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountGT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHotWalletFeeAmount), v))
	})
}

// HotWalletFeeAmountGTE applies the GTE predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountGTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHotWalletFeeAmount), v))
	})
}

// HotWalletFeeAmountLT applies the LT predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountLT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHotWalletFeeAmount), v))
	})
}

// HotWalletFeeAmountLTE applies the LTE predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountLTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHotWalletFeeAmount), v))
	})
}

// HotWalletFeeAmountIsNil applies the IsNil predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHotWalletFeeAmount)))
	})
}

// HotWalletFeeAmountNotNil applies the NotNil predicate on the "hot_wallet_fee_amount" field.
func HotWalletFeeAmountNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHotWalletFeeAmount)))
	})
}

// LowFeeAmountEQ applies the EQ predicate on the "low_fee_amount" field.
func LowFeeAmountEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLowFeeAmount), v))
	})
}

// LowFeeAmountNEQ applies the NEQ predicate on the "low_fee_amount" field.
func LowFeeAmountNEQ(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLowFeeAmount), v))
	})
}

// LowFeeAmountIn applies the In predicate on the "low_fee_amount" field.
func LowFeeAmountIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLowFeeAmount), v...))
	})
}

// LowFeeAmountNotIn applies the NotIn predicate on the "low_fee_amount" field.
func LowFeeAmountNotIn(vs ...decimal.Decimal) predicate.Fee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLowFeeAmount), v...))
	})
}

// LowFeeAmountGT applies the GT predicate on the "low_fee_amount" field.
func LowFeeAmountGT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLowFeeAmount), v))
	})
}

// LowFeeAmountGTE applies the GTE predicate on the "low_fee_amount" field.
func LowFeeAmountGTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLowFeeAmount), v))
	})
}

// LowFeeAmountLT applies the LT predicate on the "low_fee_amount" field.
func LowFeeAmountLT(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLowFeeAmount), v))
	})
}

// LowFeeAmountLTE applies the LTE predicate on the "low_fee_amount" field.
func LowFeeAmountLTE(v decimal.Decimal) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLowFeeAmount), v))
	})
}

// LowFeeAmountIsNil applies the IsNil predicate on the "low_fee_amount" field.
func LowFeeAmountIsNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLowFeeAmount)))
	})
}

// LowFeeAmountNotNil applies the NotNil predicate on the "low_fee_amount" field.
func LowFeeAmountNotNil() predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLowFeeAmount)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Fee) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Fee) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
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
func Not(p predicate.Fee) predicate.Fee {
	return predicate.Fee(func(s *sql.Selector) {
		p(s.Not())
	})
}
