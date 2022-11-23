// Code generated by ent, DO NOT EDIT.

package appcoin

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// ForPay applies equality check predicate on the "for_pay" field. It's identical to ForPayEQ.
func ForPay(v bool) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForPay), v))
	})
}

// WithdrawAutoReviewAmount applies equality check predicate on the "withdraw_auto_review_amount" field. It's identical to WithdrawAutoReviewAmountEQ.
func WithdrawAutoReviewAmount(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// ProductPage applies equality check predicate on the "product_page" field. It's identical to ProductPageEQ.
func ProductPage(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductPage), v))
	})
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v bool) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// DailyRewardAmount applies equality check predicate on the "daily_reward_amount" field. It's identical to DailyRewardAmountEQ.
func DailyRewardAmount(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDailyRewardAmount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoIsNil applies the IsNil predicate on the "logo" field.
func LogoIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLogo)))
	})
}

// LogoNotNil applies the NotNil predicate on the "logo" field.
func LogoNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLogo)))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// ForPayEQ applies the EQ predicate on the "for_pay" field.
func ForPayEQ(v bool) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForPay), v))
	})
}

// ForPayNEQ applies the NEQ predicate on the "for_pay" field.
func ForPayNEQ(v bool) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldForPay), v))
	})
}

// ForPayIsNil applies the IsNil predicate on the "for_pay" field.
func ForPayIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldForPay)))
	})
}

// ForPayNotNil applies the NotNil predicate on the "for_pay" field.
func ForPayNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldForPay)))
	})
}

// WithdrawAutoReviewAmountEQ applies the EQ predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountEQ(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// WithdrawAutoReviewAmountNEQ applies the NEQ predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountNEQ(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// WithdrawAutoReviewAmountIn applies the In predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountIn(vs ...decimal.Decimal) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWithdrawAutoReviewAmount), v...))
	})
}

// WithdrawAutoReviewAmountNotIn applies the NotIn predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountNotIn(vs ...decimal.Decimal) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWithdrawAutoReviewAmount), v...))
	})
}

// WithdrawAutoReviewAmountGT applies the GT predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountGT(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// WithdrawAutoReviewAmountGTE applies the GTE predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountGTE(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// WithdrawAutoReviewAmountLT applies the LT predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountLT(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// WithdrawAutoReviewAmountLTE applies the LTE predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountLTE(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawAutoReviewAmount), v))
	})
}

// WithdrawAutoReviewAmountIsNil applies the IsNil predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWithdrawAutoReviewAmount)))
	})
}

// WithdrawAutoReviewAmountNotNil applies the NotNil predicate on the "withdraw_auto_review_amount" field.
func WithdrawAutoReviewAmountNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWithdrawAutoReviewAmount)))
	})
}

// ProductPageEQ applies the EQ predicate on the "product_page" field.
func ProductPageEQ(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductPage), v))
	})
}

// ProductPageNEQ applies the NEQ predicate on the "product_page" field.
func ProductPageNEQ(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductPage), v))
	})
}

// ProductPageIn applies the In predicate on the "product_page" field.
func ProductPageIn(vs ...string) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProductPage), v...))
	})
}

// ProductPageNotIn applies the NotIn predicate on the "product_page" field.
func ProductPageNotIn(vs ...string) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProductPage), v...))
	})
}

// ProductPageGT applies the GT predicate on the "product_page" field.
func ProductPageGT(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductPage), v))
	})
}

// ProductPageGTE applies the GTE predicate on the "product_page" field.
func ProductPageGTE(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductPage), v))
	})
}

// ProductPageLT applies the LT predicate on the "product_page" field.
func ProductPageLT(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductPage), v))
	})
}

// ProductPageLTE applies the LTE predicate on the "product_page" field.
func ProductPageLTE(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductPage), v))
	})
}

// ProductPageContains applies the Contains predicate on the "product_page" field.
func ProductPageContains(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProductPage), v))
	})
}

// ProductPageHasPrefix applies the HasPrefix predicate on the "product_page" field.
func ProductPageHasPrefix(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProductPage), v))
	})
}

// ProductPageHasSuffix applies the HasSuffix predicate on the "product_page" field.
func ProductPageHasSuffix(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProductPage), v))
	})
}

// ProductPageIsNil applies the IsNil predicate on the "product_page" field.
func ProductPageIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProductPage)))
	})
}

// ProductPageNotNil applies the NotNil predicate on the "product_page" field.
func ProductPageNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProductPage)))
	})
}

// ProductPageEqualFold applies the EqualFold predicate on the "product_page" field.
func ProductPageEqualFold(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProductPage), v))
	})
}

// ProductPageContainsFold applies the ContainsFold predicate on the "product_page" field.
func ProductPageContainsFold(v string) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProductPage), v))
	})
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v bool) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v bool) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisabled), v))
	})
}

// DisabledIsNil applies the IsNil predicate on the "disabled" field.
func DisabledIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDisabled)))
	})
}

// DisabledNotNil applies the NotNil predicate on the "disabled" field.
func DisabledNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDisabled)))
	})
}

// DailyRewardAmountEQ applies the EQ predicate on the "daily_reward_amount" field.
func DailyRewardAmountEQ(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDailyRewardAmount), v))
	})
}

// DailyRewardAmountNEQ applies the NEQ predicate on the "daily_reward_amount" field.
func DailyRewardAmountNEQ(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDailyRewardAmount), v))
	})
}

// DailyRewardAmountIn applies the In predicate on the "daily_reward_amount" field.
func DailyRewardAmountIn(vs ...decimal.Decimal) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDailyRewardAmount), v...))
	})
}

// DailyRewardAmountNotIn applies the NotIn predicate on the "daily_reward_amount" field.
func DailyRewardAmountNotIn(vs ...decimal.Decimal) predicate.AppCoin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDailyRewardAmount), v...))
	})
}

// DailyRewardAmountGT applies the GT predicate on the "daily_reward_amount" field.
func DailyRewardAmountGT(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDailyRewardAmount), v))
	})
}

// DailyRewardAmountGTE applies the GTE predicate on the "daily_reward_amount" field.
func DailyRewardAmountGTE(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDailyRewardAmount), v))
	})
}

// DailyRewardAmountLT applies the LT predicate on the "daily_reward_amount" field.
func DailyRewardAmountLT(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDailyRewardAmount), v))
	})
}

// DailyRewardAmountLTE applies the LTE predicate on the "daily_reward_amount" field.
func DailyRewardAmountLTE(v decimal.Decimal) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDailyRewardAmount), v))
	})
}

// DailyRewardAmountIsNil applies the IsNil predicate on the "daily_reward_amount" field.
func DailyRewardAmountIsNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDailyRewardAmount)))
	})
}

// DailyRewardAmountNotNil applies the NotNil predicate on the "daily_reward_amount" field.
func DailyRewardAmountNotNil() predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDailyRewardAmount)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppCoin) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppCoin) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
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
func Not(p predicate.AppCoin) predicate.AppCoin {
	return predicate.AppCoin(func(s *sql.Selector) {
		p(s.Not())
	})
}
