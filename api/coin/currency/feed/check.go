package currencyfeed

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"

	"github.com/google/uuid"
)

func validate(in *npool.CurrencyFeedReq) error {
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

	if in.GetFeedSource() == "" {
		logger.Sugar().Errorw("validate", "FeedSource", in.GetFeedSource())
		return fmt.Errorf("feedsource is invalid")
	}

	switch in.GetFeedType() {
	case npool.FeedType_CoinBase, npool.FeedType_CoinGecko:
	default:
		logger.Sugar().Errorw("validate", "FeedType", in.GetFeedType())
		return fmt.Errorf("feedtype is invalid")
	}

	return nil
}

func validateMany(in []*npool.CurrencyFeedReq) error {
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
