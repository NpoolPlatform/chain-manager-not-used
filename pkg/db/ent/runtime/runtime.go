// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/schema"
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
	// coinbaseDescID is the schema descriptor for id field.
	coinbaseDescID := coinbaseFields[0].Descriptor()
	// coinbase.DefaultID holds the default value on creation for the id field.
	coinbase.DefaultID = coinbaseDescID.Default.(func() uuid.UUID)
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
	// coinextraDescID is the schema descriptor for id field.
	coinextraDescID := coinextraFields[0].Descriptor()
	// coinextra.DefaultID holds the default value on creation for the id field.
	coinextra.DefaultID = coinextraDescID.Default.(func() uuid.UUID)
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
	// tranDescFromAccountID is the schema descriptor for from_account_id field.
	tranDescFromAccountID := tranFields[1].Descriptor()
	// tran.DefaultFromAccountID holds the default value on creation for the from_account_id field.
	tran.DefaultFromAccountID = tranDescFromAccountID.Default.(func() uuid.UUID)
	// tranDescToAccountID is the schema descriptor for to_account_id field.
	tranDescToAccountID := tranFields[2].Descriptor()
	// tran.DefaultToAccountID holds the default value on creation for the to_account_id field.
	tran.DefaultToAccountID = tranDescToAccountID.Default.(func() uuid.UUID)
	// tranDescAmount is the schema descriptor for amount field.
	tranDescAmount := tranFields[3].Descriptor()
	// tran.DefaultAmount holds the default value on creation for the amount field.
	tran.DefaultAmount = tranDescAmount.Default.(decimal.Decimal)
	// tranDescFeeAmount is the schema descriptor for fee_amount field.
	tranDescFeeAmount := tranFields[4].Descriptor()
	// tran.DefaultFeeAmount holds the default value on creation for the fee_amount field.
	tran.DefaultFeeAmount = tranDescFeeAmount.Default.(decimal.Decimal)
	// tranDescChainTxID is the schema descriptor for chain_tx_id field.
	tranDescChainTxID := tranFields[5].Descriptor()
	// tran.DefaultChainTxID holds the default value on creation for the chain_tx_id field.
	tran.DefaultChainTxID = tranDescChainTxID.Default.(string)
	// tranDescState is the schema descriptor for state field.
	tranDescState := tranFields[6].Descriptor()
	// tran.DefaultState holds the default value on creation for the state field.
	tran.DefaultState = tranDescState.Default.(string)
	// tranDescExtra is the schema descriptor for extra field.
	tranDescExtra := tranFields[7].Descriptor()
	// tran.DefaultExtra holds the default value on creation for the extra field.
	tran.DefaultExtra = tranDescExtra.Default.(string)
	// tranDescID is the schema descriptor for id field.
	tranDescID := tranFields[0].Descriptor()
	// tran.DefaultID holds the default value on creation for the id field.
	tran.DefaultID = tranDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)
