package exrate

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func validate(in *npool.ExchangeRateReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return err
	}

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}

	if _, err := decimal.NewFromString(in.GetMarketValue()); err != nil {
		logger.Sugar().Errorw("validate", "MarketValue", in.GetMarketValue(), "error", err)
		return err
	}

	if in.GetSettlePercent() > 100 || in.GetSettlePercent() <= 0 {
		logger.Sugar().Errorw("validate", "SettlePercent", in.GetSettlePercent())
		return fmt.Errorf("settlepercent is invalidl")
	}

	return nil
}

func validateMany(in []*npool.ExchangeRateReq) error {
	apps := map[string]struct{}{}

	for _, info := range in {
		if err := validate(info); err != nil {
			return err
		}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return fmt.Errorf("appid is invalid")
	}

	return nil
}

func validateConds(conds *npool.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return err
		}
	}

	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return err
		}
	}

	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			return err
		}
	}

	return nil
}
