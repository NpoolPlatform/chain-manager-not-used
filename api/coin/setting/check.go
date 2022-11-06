package setting

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func validate(in *npool.SettingReq) error { //nolint
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
	if _, err := uuid.Parse(in.GetFeeCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "FeeCoinTypeID", in.GetFeeCoinTypeID(), "error", err)
		return err
	}
	amount, err := decimal.NewFromString(in.GetWithdrawFeeAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "WithdrawFeeAmount", in.GetWithdrawFeeAmount(), "error", err)
		return err
	}
	amount, err = decimal.NewFromString(in.GetCollectFeeAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "CollectFeeAmount", in.GetCollectFeeAmount(), "error", err)
		return err
	}
	amount, err = decimal.NewFromString(in.GetHotWalletFeeAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "HotWalletFeeAmount", in.GetHotWalletFeeAmount(), "error", err)
		return err
	}
	amount, err = decimal.NewFromString(in.GetLowFeeAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "LowFeeAmount", in.GetLowFeeAmount(), "error", err)
		return err
	}
	amount, err = decimal.NewFromString(in.GetHotWalletAccountAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "HotWalletAccountAmount", in.GetHotWalletAccountAmount(), "error", err)
		return err
	}
	amount, err = decimal.NewFromString(in.GetPaymentAccountCollectAmount())
	if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "PaymentAccountCollectAmount", in.GetPaymentAccountCollectAmount(), "error", err)
		return err
	}

	return nil
}

func validateMany(in []*npool.SettingReq) error {
	coinTypeIDs := map[string]int{}

	for _, info := range in {
		if err := validate(info); err != nil {
			return err
		}
		coinTypeIDs[info.GetCoinTypeID()] += 1
	}

	for k, v := range coinTypeIDs {
		if v > 1 {
			return fmt.Errorf("cointypeid %v is invalid", k)
		}
	}

	return nil
}

func validateConds(conds *npool.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "ID", conds.GetID().GetValue(), "error", err)
			return err
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "CoinTypeID", conds.GetCoinTypeID().GetValue(), "error", err)
			return err
		}
	}
	if conds.FeeCoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetFeeCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "FeeCoinTypeID", conds.GetFeeCoinTypeID().GetValue(), "error", err)
			return err
		}
	}

	return nil
}

func validateUpdate(in *npool.SettingReq) error { //nolint
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return err
	}

	if in.WithdrawFeeAmount != nil {
		amount, err := decimal.NewFromString(in.GetWithdrawFeeAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "WithdrawFeeAmount", in.GetWithdrawFeeAmount(), "error", err)
			return err
		}
	}

	if in.CollectFeeAmount != nil {
		amount, err := decimal.NewFromString(in.GetCollectFeeAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "CollectFeeAmount", in.GetCollectFeeAmount(), "error", err)
			return err
		}
	}

	if in.HotWalletFeeAmount != nil {
		amount, err := decimal.NewFromString(in.GetHotWalletFeeAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "HotWalletFeeAmount", in.GetHotWalletFeeAmount(), "error", err)
			return err
		}
	}

	if in.LowFeeAmount != nil {
		amount, err := decimal.NewFromString(in.GetLowFeeAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "LowFeeAmount", in.GetLowFeeAmount(), "error", err)
			return err
		}
	}

	if in.HotWalletAccountAmount != nil {
		amount, err := decimal.NewFromString(in.GetHotWalletAccountAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "HotWalletAccountAmount", in.GetHotWalletAccountAmount(), "error", err)
			return err
		}
	}

	if in.PaymentAccountCollectAmount != nil {
		amount, err := decimal.NewFromString(in.GetPaymentAccountCollectAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "PaymentAccountCollectAmount", in.GetPaymentAccountCollectAmount(), "error", err)
			return err
		}
	}

	return nil
}
