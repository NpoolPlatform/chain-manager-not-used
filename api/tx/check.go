package tx

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func validate(in *npool.TxReq) error { //nolint
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetFromAccountID()); err != nil {
		logger.Sugar().Errorw("validate", "FromAccountID", in.GetFromAccountID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetToAccountID()); err != nil {
		logger.Sugar().Errorw("validate", "ToAccountID", in.GetToAccountID(), "error", err)
		return err
	}

	amount, err := decimal.NewFromString(in.GetAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "Amount", in.GetAmount(), "error", err)
		return fmt.Errorf("amount is invalid: %v", err)
	}
	amount, err = decimal.NewFromString(in.GetFeeAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("validate", "FeeAmount", in.GetFeeAmount(), "error", err)
		return fmt.Errorf("feeamount is invalid: %v", err)
	}
	if in.ChainTxID != nil && in.GetChainTxID() == "" {
		logger.Sugar().Errorw("validate", "ChainTxID", in.GetChainTxID(), "error", err)
		return fmt.Errorf("chaintxid is invalid")
	}

	if in.State != nil {
		switch in.GetState() {
		case npool.TxState_StateCreated:
		case npool.TxState_StateWait:
		case npool.TxState_StateTransferring:
		case npool.TxState_StateSuccessful:
		case npool.TxState_StateFail:
		default:
			logger.Sugar().Errorw("validate", "State", in.GetState())
			return fmt.Errorf("state is invalid")
		}
	}

	switch in.GetType() {
	case basetypes.TxType_TxWithdraw:
	case basetypes.TxType_TxFeedGas:
	case basetypes.TxType_TxPaymentCollect:
	case basetypes.TxType_TxBenefit:
	case basetypes.TxType_TxLimitation:
	case basetypes.TxType_TxPlatformBenefit:
	case basetypes.TxType_TxUserBenefit:
	default:
		logger.Sugar().Errorw("validate", "Type", in.GetType())
		return fmt.Errorf("type is invalid")
	}

	return nil
}

func validateMany(in []*npool.TxReq) error {
	for _, info := range in {
		if err := validate(info); err != nil {
			return err
		}
	}
	return nil
}

func validateConds(conds *npool.Conds) error { //nolint
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "ID", conds.GetID().GetValue(), "error", err)
			return err
		}
	}

	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "CoinTypeID", conds.GetCoinTypeID().GetValue(), "error", err)
			return err
		}
	}

	if conds.AccountID != nil {
		if _, err := uuid.Parse(conds.GetAccountID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "AccountID", conds.GetAccountID().GetValue(), "error", err)
			return err
		}
	}

	for _, id := range conds.GetAccountIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("validateConds", "AccountID", id, "error", err)
			return err
		}
	}

	if conds.State != nil {
		switch npool.TxState(conds.GetState().GetValue()) {
		case npool.TxState_StateCreated:
		case npool.TxState_StateWait:
		case npool.TxState_StateTransferring:
		case npool.TxState_StateSuccessful:
		case npool.TxState_StateFail:
		default:
			logger.Sugar().Errorw("validateConds", "State", conds.GetState().GetValue())
			return fmt.Errorf("state is invalid")
		}
	}

	if conds.Type != nil {
		switch basetypes.TxType(conds.GetType().GetValue()) {
		case basetypes.TxType_TxWithdraw:
		case basetypes.TxType_TxFeedGas:
		case basetypes.TxType_TxPaymentCollect:
		case basetypes.TxType_TxBenefit:
		case basetypes.TxType_TxLimitation:
		case basetypes.TxType_TxPlatformBenefit:
		case basetypes.TxType_TxUserBenefit:
		default:
			logger.Sugar().Errorw("validateConds", "Type", conds.GetType().GetValue())
			return fmt.Errorf("type is invalid")
		}
	}

	return nil
}
