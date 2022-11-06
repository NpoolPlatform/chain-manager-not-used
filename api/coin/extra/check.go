package coinextra

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"

	"github.com/google/uuid"
)

func validate(in *npool.CoinExtraReq) error {
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

	return nil
}

func validateMany(in []*npool.CoinExtraReq) error {
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

	return nil
}
