// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currency"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/schema"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appcoinMixin := schema.AppCoin{}.Mixin()
	appcoin.Policy = privacy.NewPolicies(appcoinMixin[0], schema.AppCoin{})
	appcoin.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := appcoin.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	appcoinMixinFields0 := appcoinMixin[0].Fields()
	_ = appcoinMixinFields0
	appcoinFields := schema.AppCoin{}.Fields()
	_ = appcoinFields
	// appcoinDescCreatedAt is the schema descriptor for created_at field.
	appcoinDescCreatedAt := appcoinMixinFields0[0].Descriptor()
	// appcoin.DefaultCreatedAt holds the default value on creation for the created_at field.
	appcoin.DefaultCreatedAt = appcoinDescCreatedAt.Default.(func() uint32)
	// appcoinDescUpdatedAt is the schema descriptor for updated_at field.
	appcoinDescUpdatedAt := appcoinMixinFields0[1].Descriptor()
	// appcoin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appcoin.DefaultUpdatedAt = appcoinDescUpdatedAt.Default.(func() uint32)
	// appcoin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appcoin.UpdateDefaultUpdatedAt = appcoinDescUpdatedAt.UpdateDefault.(func() uint32)
	// appcoinDescDeletedAt is the schema descriptor for deleted_at field.
	appcoinDescDeletedAt := appcoinMixinFields0[2].Descriptor()
	// appcoin.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	appcoin.DefaultDeletedAt = appcoinDescDeletedAt.Default.(func() uint32)
	// appcoinDescAppID is the schema descriptor for app_id field.
	appcoinDescAppID := appcoinFields[1].Descriptor()
	// appcoin.DefaultAppID holds the default value on creation for the app_id field.
	appcoin.DefaultAppID = appcoinDescAppID.Default.(func() uuid.UUID)
	// appcoinDescCoinTypeID is the schema descriptor for coin_type_id field.
	appcoinDescCoinTypeID := appcoinFields[2].Descriptor()
	// appcoin.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	appcoin.DefaultCoinTypeID = appcoinDescCoinTypeID.Default.(func() uuid.UUID)
	// appcoinDescName is the schema descriptor for name field.
	appcoinDescName := appcoinFields[3].Descriptor()
	// appcoin.DefaultName holds the default value on creation for the name field.
	appcoin.DefaultName = appcoinDescName.Default.(string)
	// appcoinDescLogo is the schema descriptor for logo field.
	appcoinDescLogo := appcoinFields[4].Descriptor()
	// appcoin.DefaultLogo holds the default value on creation for the logo field.
	appcoin.DefaultLogo = appcoinDescLogo.Default.(string)
	// appcoinDescForPay is the schema descriptor for for_pay field.
	appcoinDescForPay := appcoinFields[5].Descriptor()
	// appcoin.DefaultForPay holds the default value on creation for the for_pay field.
	appcoin.DefaultForPay = appcoinDescForPay.Default.(bool)
	// appcoinDescWithdrawAutoReviewAmount is the schema descriptor for withdraw_auto_review_amount field.
	appcoinDescWithdrawAutoReviewAmount := appcoinFields[6].Descriptor()
	// appcoin.DefaultWithdrawAutoReviewAmount holds the default value on creation for the withdraw_auto_review_amount field.
	appcoin.DefaultWithdrawAutoReviewAmount = appcoinDescWithdrawAutoReviewAmount.Default.(decimal.Decimal)
	// appcoinDescProductPage is the schema descriptor for product_page field.
	appcoinDescProductPage := appcoinFields[7].Descriptor()
	// appcoin.DefaultProductPage holds the default value on creation for the product_page field.
	appcoin.DefaultProductPage = appcoinDescProductPage.Default.(string)
	// appcoinDescDisabled is the schema descriptor for disabled field.
	appcoinDescDisabled := appcoinFields[8].Descriptor()
	// appcoin.DefaultDisabled holds the default value on creation for the disabled field.
	appcoin.DefaultDisabled = appcoinDescDisabled.Default.(bool)
	// appcoinDescDailyRewardAmount is the schema descriptor for daily_reward_amount field.
	appcoinDescDailyRewardAmount := appcoinFields[9].Descriptor()
	// appcoin.DefaultDailyRewardAmount holds the default value on creation for the daily_reward_amount field.
	appcoin.DefaultDailyRewardAmount = appcoinDescDailyRewardAmount.Default.(decimal.Decimal)
	// appcoinDescID is the schema descriptor for id field.
	appcoinDescID := appcoinFields[0].Descriptor()
	// appcoin.DefaultID holds the default value on creation for the id field.
	appcoin.DefaultID = appcoinDescID.Default.(func() uuid.UUID)
	coinbaseMixin := schema.CoinBase{}.Mixin()
	coinbase.Policy = privacy.NewPolicies(coinbaseMixin[0], schema.CoinBase{})
	coinbase.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coinbase.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coinbaseMixinFields0 := coinbaseMixin[0].Fields()
	_ = coinbaseMixinFields0
	coinbaseFields := schema.CoinBase{}.Fields()
	_ = coinbaseFields
	// coinbaseDescCreatedAt is the schema descriptor for created_at field.
	coinbaseDescCreatedAt := coinbaseMixinFields0[0].Descriptor()
	// coinbase.DefaultCreatedAt holds the default value on creation for the created_at field.
	coinbase.DefaultCreatedAt = coinbaseDescCreatedAt.Default.(func() uint32)
	// coinbaseDescUpdatedAt is the schema descriptor for updated_at field.
	coinbaseDescUpdatedAt := coinbaseMixinFields0[1].Descriptor()
	// coinbase.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coinbase.DefaultUpdatedAt = coinbaseDescUpdatedAt.Default.(func() uint32)
	// coinbase.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coinbase.UpdateDefaultUpdatedAt = coinbaseDescUpdatedAt.UpdateDefault.(func() uint32)
	// coinbaseDescDeletedAt is the schema descriptor for deleted_at field.
	coinbaseDescDeletedAt := coinbaseMixinFields0[2].Descriptor()
	// coinbase.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coinbase.DefaultDeletedAt = coinbaseDescDeletedAt.Default.(func() uint32)
	// coinbaseDescName is the schema descriptor for name field.
	coinbaseDescName := coinbaseFields[1].Descriptor()
	// coinbase.DefaultName holds the default value on creation for the name field.
	coinbase.DefaultName = coinbaseDescName.Default.(string)
	// coinbaseDescLogo is the schema descriptor for logo field.
	coinbaseDescLogo := coinbaseFields[2].Descriptor()
	// coinbase.DefaultLogo holds the default value on creation for the logo field.
	coinbase.DefaultLogo = coinbaseDescLogo.Default.(string)
	// coinbaseDescPresale is the schema descriptor for presale field.
	coinbaseDescPresale := coinbaseFields[3].Descriptor()
	// coinbase.DefaultPresale holds the default value on creation for the presale field.
	coinbase.DefaultPresale = coinbaseDescPresale.Default.(bool)
	// coinbaseDescUnit is the schema descriptor for unit field.
	coinbaseDescUnit := coinbaseFields[4].Descriptor()
	// coinbase.DefaultUnit holds the default value on creation for the unit field.
	coinbase.DefaultUnit = coinbaseDescUnit.Default.(string)
	// coinbaseDescEnv is the schema descriptor for env field.
	coinbaseDescEnv := coinbaseFields[5].Descriptor()
	// coinbase.DefaultEnv holds the default value on creation for the env field.
	coinbase.DefaultEnv = coinbaseDescEnv.Default.(string)
	// coinbaseDescReservedAmount is the schema descriptor for reserved_amount field.
	coinbaseDescReservedAmount := coinbaseFields[6].Descriptor()
	// coinbase.DefaultReservedAmount holds the default value on creation for the reserved_amount field.
	coinbase.DefaultReservedAmount = coinbaseDescReservedAmount.Default.(decimal.Decimal)
	// coinbaseDescForPay is the schema descriptor for for_pay field.
	coinbaseDescForPay := coinbaseFields[7].Descriptor()
	// coinbase.DefaultForPay holds the default value on creation for the for_pay field.
	coinbase.DefaultForPay = coinbaseDescForPay.Default.(bool)
	// coinbaseDescDisabled is the schema descriptor for disabled field.
	coinbaseDescDisabled := coinbaseFields[8].Descriptor()
	// coinbase.DefaultDisabled holds the default value on creation for the disabled field.
	coinbase.DefaultDisabled = coinbaseDescDisabled.Default.(bool)
	// coinbaseDescID is the schema descriptor for id field.
	coinbaseDescID := coinbaseFields[0].Descriptor()
	// coinbase.DefaultID holds the default value on creation for the id field.
	coinbase.DefaultID = coinbaseDescID.Default.(func() uuid.UUID)
	coindescriptionMixin := schema.CoinDescription{}.Mixin()
	coindescription.Policy = privacy.NewPolicies(coindescriptionMixin[0], schema.CoinDescription{})
	coindescription.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coindescription.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coindescriptionMixinFields0 := coindescriptionMixin[0].Fields()
	_ = coindescriptionMixinFields0
	coindescriptionFields := schema.CoinDescription{}.Fields()
	_ = coindescriptionFields
	// coindescriptionDescCreatedAt is the schema descriptor for created_at field.
	coindescriptionDescCreatedAt := coindescriptionMixinFields0[0].Descriptor()
	// coindescription.DefaultCreatedAt holds the default value on creation for the created_at field.
	coindescription.DefaultCreatedAt = coindescriptionDescCreatedAt.Default.(func() uint32)
	// coindescriptionDescUpdatedAt is the schema descriptor for updated_at field.
	coindescriptionDescUpdatedAt := coindescriptionMixinFields0[1].Descriptor()
	// coindescription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coindescription.DefaultUpdatedAt = coindescriptionDescUpdatedAt.Default.(func() uint32)
	// coindescription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coindescription.UpdateDefaultUpdatedAt = coindescriptionDescUpdatedAt.UpdateDefault.(func() uint32)
	// coindescriptionDescDeletedAt is the schema descriptor for deleted_at field.
	coindescriptionDescDeletedAt := coindescriptionMixinFields0[2].Descriptor()
	// coindescription.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coindescription.DefaultDeletedAt = coindescriptionDescDeletedAt.Default.(func() uint32)
	// coindescriptionDescAppID is the schema descriptor for app_id field.
	coindescriptionDescAppID := coindescriptionFields[1].Descriptor()
	// coindescription.DefaultAppID holds the default value on creation for the app_id field.
	coindescription.DefaultAppID = coindescriptionDescAppID.Default.(func() uuid.UUID)
	// coindescriptionDescCoinTypeID is the schema descriptor for coin_type_id field.
	coindescriptionDescCoinTypeID := coindescriptionFields[2].Descriptor()
	// coindescription.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	coindescription.DefaultCoinTypeID = coindescriptionDescCoinTypeID.Default.(func() uuid.UUID)
	// coindescriptionDescUsedFor is the schema descriptor for used_for field.
	coindescriptionDescUsedFor := coindescriptionFields[3].Descriptor()
	// coindescription.DefaultUsedFor holds the default value on creation for the used_for field.
	coindescription.DefaultUsedFor = coindescriptionDescUsedFor.Default.(string)
	// coindescriptionDescTitle is the schema descriptor for title field.
	coindescriptionDescTitle := coindescriptionFields[4].Descriptor()
	// coindescription.DefaultTitle holds the default value on creation for the title field.
	coindescription.DefaultTitle = coindescriptionDescTitle.Default.(string)
	// coindescriptionDescMessage is the schema descriptor for message field.
	coindescriptionDescMessage := coindescriptionFields[5].Descriptor()
	// coindescription.DefaultMessage holds the default value on creation for the message field.
	coindescription.DefaultMessage = coindescriptionDescMessage.Default.(string)
	// coindescriptionDescID is the schema descriptor for id field.
	coindescriptionDescID := coindescriptionFields[0].Descriptor()
	// coindescription.DefaultID holds the default value on creation for the id field.
	coindescription.DefaultID = coindescriptionDescID.Default.(func() uuid.UUID)
	coinextraMixin := schema.CoinExtra{}.Mixin()
	coinextra.Policy = privacy.NewPolicies(coinextraMixin[0], schema.CoinExtra{})
	coinextra.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coinextra.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coinextraMixinFields0 := coinextraMixin[0].Fields()
	_ = coinextraMixinFields0
	coinextraFields := schema.CoinExtra{}.Fields()
	_ = coinextraFields
	// coinextraDescCreatedAt is the schema descriptor for created_at field.
	coinextraDescCreatedAt := coinextraMixinFields0[0].Descriptor()
	// coinextra.DefaultCreatedAt holds the default value on creation for the created_at field.
	coinextra.DefaultCreatedAt = coinextraDescCreatedAt.Default.(func() uint32)
	// coinextraDescUpdatedAt is the schema descriptor for updated_at field.
	coinextraDescUpdatedAt := coinextraMixinFields0[1].Descriptor()
	// coinextra.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coinextra.DefaultUpdatedAt = coinextraDescUpdatedAt.Default.(func() uint32)
	// coinextra.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coinextra.UpdateDefaultUpdatedAt = coinextraDescUpdatedAt.UpdateDefault.(func() uint32)
	// coinextraDescDeletedAt is the schema descriptor for deleted_at field.
	coinextraDescDeletedAt := coinextraMixinFields0[2].Descriptor()
	// coinextra.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coinextra.DefaultDeletedAt = coinextraDescDeletedAt.Default.(func() uint32)
	// coinextraDescCoinTypeID is the schema descriptor for coin_type_id field.
	coinextraDescCoinTypeID := coinextraFields[1].Descriptor()
	// coinextra.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	coinextra.DefaultCoinTypeID = coinextraDescCoinTypeID.Default.(func() uuid.UUID)
	// coinextraDescHomePage is the schema descriptor for home_page field.
	coinextraDescHomePage := coinextraFields[2].Descriptor()
	// coinextra.DefaultHomePage holds the default value on creation for the home_page field.
	coinextra.DefaultHomePage = coinextraDescHomePage.Default.(string)
	// coinextraDescSpecs is the schema descriptor for specs field.
	coinextraDescSpecs := coinextraFields[3].Descriptor()
	// coinextra.DefaultSpecs holds the default value on creation for the specs field.
	coinextra.DefaultSpecs = coinextraDescSpecs.Default.(string)
	// coinextraDescStableUsd is the schema descriptor for stable_usd field.
	coinextraDescStableUsd := coinextraFields[4].Descriptor()
	// coinextra.DefaultStableUsd holds the default value on creation for the stable_usd field.
	coinextra.DefaultStableUsd = coinextraDescStableUsd.Default.(bool)
	// coinextraDescID is the schema descriptor for id field.
	coinextraDescID := coinextraFields[0].Descriptor()
	// coinextra.DefaultID holds the default value on creation for the id field.
	coinextra.DefaultID = coinextraDescID.Default.(func() uuid.UUID)
	currencyMixin := schema.Currency{}.Mixin()
	currency.Policy = privacy.NewPolicies(currencyMixin[0], schema.Currency{})
	currency.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := currency.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	currencyMixinFields0 := currencyMixin[0].Fields()
	_ = currencyMixinFields0
	currencyFields := schema.Currency{}.Fields()
	_ = currencyFields
	// currencyDescCreatedAt is the schema descriptor for created_at field.
	currencyDescCreatedAt := currencyMixinFields0[0].Descriptor()
	// currency.DefaultCreatedAt holds the default value on creation for the created_at field.
	currency.DefaultCreatedAt = currencyDescCreatedAt.Default.(func() uint32)
	// currencyDescUpdatedAt is the schema descriptor for updated_at field.
	currencyDescUpdatedAt := currencyMixinFields0[1].Descriptor()
	// currency.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	currency.DefaultUpdatedAt = currencyDescUpdatedAt.Default.(func() uint32)
	// currency.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	currency.UpdateDefaultUpdatedAt = currencyDescUpdatedAt.UpdateDefault.(func() uint32)
	// currencyDescDeletedAt is the schema descriptor for deleted_at field.
	currencyDescDeletedAt := currencyMixinFields0[2].Descriptor()
	// currency.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	currency.DefaultDeletedAt = currencyDescDeletedAt.Default.(func() uint32)
	// currencyDescCoinTypeID is the schema descriptor for coin_type_id field.
	currencyDescCoinTypeID := currencyFields[1].Descriptor()
	// currency.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	currency.DefaultCoinTypeID = currencyDescCoinTypeID.Default.(func() uuid.UUID)
	// currencyDescFeedType is the schema descriptor for feed_type field.
	currencyDescFeedType := currencyFields[2].Descriptor()
	// currency.DefaultFeedType holds the default value on creation for the feed_type field.
	currency.DefaultFeedType = currencyDescFeedType.Default.(string)
	// currencyDescMarketValueHigh is the schema descriptor for market_value_high field.
	currencyDescMarketValueHigh := currencyFields[3].Descriptor()
	// currency.DefaultMarketValueHigh holds the default value on creation for the market_value_high field.
	currency.DefaultMarketValueHigh = currencyDescMarketValueHigh.Default.(decimal.Decimal)
	// currencyDescMarketValueLow is the schema descriptor for market_value_low field.
	currencyDescMarketValueLow := currencyFields[4].Descriptor()
	// currency.DefaultMarketValueLow holds the default value on creation for the market_value_low field.
	currency.DefaultMarketValueLow = currencyDescMarketValueLow.Default.(decimal.Decimal)
	// currencyDescID is the schema descriptor for id field.
	currencyDescID := currencyFields[0].Descriptor()
	// currency.DefaultID holds the default value on creation for the id field.
	currency.DefaultID = currencyDescID.Default.(func() uuid.UUID)
	exchangerateMixin := schema.ExchangeRate{}.Mixin()
	exchangerate.Policy = privacy.NewPolicies(exchangerateMixin[0], schema.ExchangeRate{})
	exchangerate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := exchangerate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	exchangerateMixinFields0 := exchangerateMixin[0].Fields()
	_ = exchangerateMixinFields0
	exchangerateFields := schema.ExchangeRate{}.Fields()
	_ = exchangerateFields
	// exchangerateDescCreatedAt is the schema descriptor for created_at field.
	exchangerateDescCreatedAt := exchangerateMixinFields0[0].Descriptor()
	// exchangerate.DefaultCreatedAt holds the default value on creation for the created_at field.
	exchangerate.DefaultCreatedAt = exchangerateDescCreatedAt.Default.(func() uint32)
	// exchangerateDescUpdatedAt is the schema descriptor for updated_at field.
	exchangerateDescUpdatedAt := exchangerateMixinFields0[1].Descriptor()
	// exchangerate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	exchangerate.DefaultUpdatedAt = exchangerateDescUpdatedAt.Default.(func() uint32)
	// exchangerate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	exchangerate.UpdateDefaultUpdatedAt = exchangerateDescUpdatedAt.UpdateDefault.(func() uint32)
	// exchangerateDescDeletedAt is the schema descriptor for deleted_at field.
	exchangerateDescDeletedAt := exchangerateMixinFields0[2].Descriptor()
	// exchangerate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	exchangerate.DefaultDeletedAt = exchangerateDescDeletedAt.Default.(func() uint32)
	// exchangerateDescAppID is the schema descriptor for app_id field.
	exchangerateDescAppID := exchangerateFields[1].Descriptor()
	// exchangerate.DefaultAppID holds the default value on creation for the app_id field.
	exchangerate.DefaultAppID = exchangerateDescAppID.Default.(func() uuid.UUID)
	// exchangerateDescCoinTypeID is the schema descriptor for coin_type_id field.
	exchangerateDescCoinTypeID := exchangerateFields[2].Descriptor()
	// exchangerate.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	exchangerate.DefaultCoinTypeID = exchangerateDescCoinTypeID.Default.(func() uuid.UUID)
	// exchangerateDescMarketValue is the schema descriptor for market_value field.
	exchangerateDescMarketValue := exchangerateFields[3].Descriptor()
	// exchangerate.DefaultMarketValue holds the default value on creation for the market_value field.
	exchangerate.DefaultMarketValue = exchangerateDescMarketValue.Default.(decimal.Decimal)
	// exchangerateDescSettleValue is the schema descriptor for settle_value field.
	exchangerateDescSettleValue := exchangerateFields[4].Descriptor()
	// exchangerate.DefaultSettleValue holds the default value on creation for the settle_value field.
	exchangerate.DefaultSettleValue = exchangerateDescSettleValue.Default.(decimal.Decimal)
	// exchangerateDescSettlePercent is the schema descriptor for settle_percent field.
	exchangerateDescSettlePercent := exchangerateFields[5].Descriptor()
	// exchangerate.DefaultSettlePercent holds the default value on creation for the settle_percent field.
	exchangerate.DefaultSettlePercent = exchangerateDescSettlePercent.Default.(uint32)
	// exchangerateDescSetter is the schema descriptor for setter field.
	exchangerateDescSetter := exchangerateFields[6].Descriptor()
	// exchangerate.DefaultSetter holds the default value on creation for the setter field.
	exchangerate.DefaultSetter = exchangerateDescSetter.Default.(func() uuid.UUID)
	// exchangerateDescID is the schema descriptor for id field.
	exchangerateDescID := exchangerateFields[0].Descriptor()
	// exchangerate.DefaultID holds the default value on creation for the id field.
	exchangerate.DefaultID = exchangerateDescID.Default.(func() uuid.UUID)
	settingMixin := schema.Setting{}.Mixin()
	setting.Policy = privacy.NewPolicies(settingMixin[0], schema.Setting{})
	setting.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := setting.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	settingMixinFields0 := settingMixin[0].Fields()
	_ = settingMixinFields0
	settingFields := schema.Setting{}.Fields()
	_ = settingFields
	// settingDescCreatedAt is the schema descriptor for created_at field.
	settingDescCreatedAt := settingMixinFields0[0].Descriptor()
	// setting.DefaultCreatedAt holds the default value on creation for the created_at field.
	setting.DefaultCreatedAt = settingDescCreatedAt.Default.(func() uint32)
	// settingDescUpdatedAt is the schema descriptor for updated_at field.
	settingDescUpdatedAt := settingMixinFields0[1].Descriptor()
	// setting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	setting.DefaultUpdatedAt = settingDescUpdatedAt.Default.(func() uint32)
	// setting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	setting.UpdateDefaultUpdatedAt = settingDescUpdatedAt.UpdateDefault.(func() uint32)
	// settingDescDeletedAt is the schema descriptor for deleted_at field.
	settingDescDeletedAt := settingMixinFields0[2].Descriptor()
	// setting.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	setting.DefaultDeletedAt = settingDescDeletedAt.Default.(func() uint32)
	// settingDescCoinTypeID is the schema descriptor for coin_type_id field.
	settingDescCoinTypeID := settingFields[1].Descriptor()
	// setting.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	setting.DefaultCoinTypeID = settingDescCoinTypeID.Default.(func() uuid.UUID)
	// settingDescFeeCoinTypeID is the schema descriptor for fee_coin_type_id field.
	settingDescFeeCoinTypeID := settingFields[2].Descriptor()
	// setting.DefaultFeeCoinTypeID holds the default value on creation for the fee_coin_type_id field.
	setting.DefaultFeeCoinTypeID = settingDescFeeCoinTypeID.Default.(func() uuid.UUID)
	// settingDescWithdrawFeeByStableUsd is the schema descriptor for withdraw_fee_by_stable_usd field.
	settingDescWithdrawFeeByStableUsd := settingFields[3].Descriptor()
	// setting.DefaultWithdrawFeeByStableUsd holds the default value on creation for the withdraw_fee_by_stable_usd field.
	setting.DefaultWithdrawFeeByStableUsd = settingDescWithdrawFeeByStableUsd.Default.(bool)
	// settingDescWithdrawFeeAmount is the schema descriptor for withdraw_fee_amount field.
	settingDescWithdrawFeeAmount := settingFields[4].Descriptor()
	// setting.DefaultWithdrawFeeAmount holds the default value on creation for the withdraw_fee_amount field.
	setting.DefaultWithdrawFeeAmount = settingDescWithdrawFeeAmount.Default.(decimal.Decimal)
	// settingDescCollectFeeAmount is the schema descriptor for collect_fee_amount field.
	settingDescCollectFeeAmount := settingFields[5].Descriptor()
	// setting.DefaultCollectFeeAmount holds the default value on creation for the collect_fee_amount field.
	setting.DefaultCollectFeeAmount = settingDescCollectFeeAmount.Default.(decimal.Decimal)
	// settingDescHotWalletFeeAmount is the schema descriptor for hot_wallet_fee_amount field.
	settingDescHotWalletFeeAmount := settingFields[6].Descriptor()
	// setting.DefaultHotWalletFeeAmount holds the default value on creation for the hot_wallet_fee_amount field.
	setting.DefaultHotWalletFeeAmount = settingDescHotWalletFeeAmount.Default.(decimal.Decimal)
	// settingDescLowFeeAmount is the schema descriptor for low_fee_amount field.
	settingDescLowFeeAmount := settingFields[7].Descriptor()
	// setting.DefaultLowFeeAmount holds the default value on creation for the low_fee_amount field.
	setting.DefaultLowFeeAmount = settingDescLowFeeAmount.Default.(decimal.Decimal)
	// settingDescHotWalletAccountAmount is the schema descriptor for hot_wallet_account_amount field.
	settingDescHotWalletAccountAmount := settingFields[8].Descriptor()
	// setting.DefaultHotWalletAccountAmount holds the default value on creation for the hot_wallet_account_amount field.
	setting.DefaultHotWalletAccountAmount = settingDescHotWalletAccountAmount.Default.(decimal.Decimal)
	// settingDescPaymentAccountCollectAmount is the schema descriptor for payment_account_collect_amount field.
	settingDescPaymentAccountCollectAmount := settingFields[9].Descriptor()
	// setting.DefaultPaymentAccountCollectAmount holds the default value on creation for the payment_account_collect_amount field.
	setting.DefaultPaymentAccountCollectAmount = settingDescPaymentAccountCollectAmount.Default.(decimal.Decimal)
	// settingDescID is the schema descriptor for id field.
	settingDescID := settingFields[0].Descriptor()
	// setting.DefaultID holds the default value on creation for the id field.
	setting.DefaultID = settingDescID.Default.(func() uuid.UUID)
	tranMixin := schema.Tran{}.Mixin()
	tran.Policy = privacy.NewPolicies(tranMixin[0], schema.Tran{})
	tran.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := tran.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tranMixinFields0 := tranMixin[0].Fields()
	_ = tranMixinFields0
	tranFields := schema.Tran{}.Fields()
	_ = tranFields
	// tranDescCreatedAt is the schema descriptor for created_at field.
	tranDescCreatedAt := tranMixinFields0[0].Descriptor()
	// tran.DefaultCreatedAt holds the default value on creation for the created_at field.
	tran.DefaultCreatedAt = tranDescCreatedAt.Default.(func() uint32)
	// tranDescUpdatedAt is the schema descriptor for updated_at field.
	tranDescUpdatedAt := tranMixinFields0[1].Descriptor()
	// tran.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tran.DefaultUpdatedAt = tranDescUpdatedAt.Default.(func() uint32)
	// tran.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tran.UpdateDefaultUpdatedAt = tranDescUpdatedAt.UpdateDefault.(func() uint32)
	// tranDescDeletedAt is the schema descriptor for deleted_at field.
	tranDescDeletedAt := tranMixinFields0[2].Descriptor()
	// tran.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	tran.DefaultDeletedAt = tranDescDeletedAt.Default.(func() uint32)
	// tranDescCoinTypeID is the schema descriptor for coin_type_id field.
	tranDescCoinTypeID := tranFields[1].Descriptor()
	// tran.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	tran.DefaultCoinTypeID = tranDescCoinTypeID.Default.(func() uuid.UUID)
	// tranDescFromAccountID is the schema descriptor for from_account_id field.
	tranDescFromAccountID := tranFields[2].Descriptor()
	// tran.DefaultFromAccountID holds the default value on creation for the from_account_id field.
	tran.DefaultFromAccountID = tranDescFromAccountID.Default.(func() uuid.UUID)
	// tranDescToAccountID is the schema descriptor for to_account_id field.
	tranDescToAccountID := tranFields[3].Descriptor()
	// tran.DefaultToAccountID holds the default value on creation for the to_account_id field.
	tran.DefaultToAccountID = tranDescToAccountID.Default.(func() uuid.UUID)
	// tranDescAmount is the schema descriptor for amount field.
	tranDescAmount := tranFields[4].Descriptor()
	// tran.DefaultAmount holds the default value on creation for the amount field.
	tran.DefaultAmount = tranDescAmount.Default.(decimal.Decimal)
	// tranDescFeeAmount is the schema descriptor for fee_amount field.
	tranDescFeeAmount := tranFields[5].Descriptor()
	// tran.DefaultFeeAmount holds the default value on creation for the fee_amount field.
	tran.DefaultFeeAmount = tranDescFeeAmount.Default.(decimal.Decimal)
	// tranDescChainTxID is the schema descriptor for chain_tx_id field.
	tranDescChainTxID := tranFields[6].Descriptor()
	// tran.DefaultChainTxID holds the default value on creation for the chain_tx_id field.
	tran.DefaultChainTxID = tranDescChainTxID.Default.(string)
	// tranDescState is the schema descriptor for state field.
	tranDescState := tranFields[7].Descriptor()
	// tran.DefaultState holds the default value on creation for the state field.
	tran.DefaultState = tranDescState.Default.(string)
	// tranDescExtra is the schema descriptor for extra field.
	tranDescExtra := tranFields[8].Descriptor()
	// tran.DefaultExtra holds the default value on creation for the extra field.
	tran.DefaultExtra = tranDescExtra.Default.(string)
	// tranDescType is the schema descriptor for type field.
	tranDescType := tranFields[9].Descriptor()
	// tran.DefaultType holds the default value on creation for the type field.
	tran.DefaultType = tranDescType.Default.(string)
	// tranDescID is the schema descriptor for id field.
	tranDescID := tranFields[0].Descriptor()
	// tran.DefaultID holds the default value on creation for the id field.
	tran.DefaultID = tranDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)
