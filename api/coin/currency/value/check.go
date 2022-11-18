package currencyvalue

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func validate(in *npool.CurrencyValueReq) error {
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

	if _, err := uuid.Parse(in.GetFeedSourceID()); err != nil {
		logger.Sugar().Errorw("validate", "FeedSourceID", in.GetFeedSourceID(), "error", err)
		return err
	}

	highValue, err := decimal.NewFromString(in.GetMarketValueHigh())
	if err != nil {
		logger.Sugar().Errorw("validate", "MarketValueHigh", in.GetMarketValueHigh(), "error", err)
		return err
	}

	lowValue, err := decimal.NewFromString(in.GetMarketValueLow())
	if err != nil {
		logger.Sugar().Errorw("validate", "MarketValueLow", in.GetMarketValueLow(), "error", err)
		return err
	}

	if highValue.Cmp(lowValue) < 0 {
		logger.Sugar().Errorw("validate", "High", highValue, "Low", lowValue)
		return fmt.Errorf("invalid market value")
	}

	return nil
}

func validateMany(in []*npool.CurrencyValueReq) error {
	for _, info := range in {
		if err := validate(info); err != nil {
			return err
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

	return nil
}

func validateUpdate(in *npool.CurrencyValueReq) error {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return err
	}

	highValue, err := decimal.NewFromString(in.GetMarketValueHigh())
	if err != nil {
		logger.Sugar().Errorw("validate", "MarketValueHigh", in.GetMarketValueHigh(), "error", err)
		return err
	}

	lowValue, err := decimal.NewFromString(in.GetMarketValueLow())
	if err != nil {
		logger.Sugar().Errorw("validate", "MarketValueLow", in.GetMarketValueLow(), "error", err)
		return err
	}

	if highValue.Cmp(lowValue) < 0 {
		logger.Sugar().Errorw("validate", "High", highValue, "Low", lowValue)
		return fmt.Errorf("invalid market value")
	}

	return nil
}
