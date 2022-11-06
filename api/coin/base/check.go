package coinbase

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func validate(in *npool.CoinBaseReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if in.GetName() == "" {
		logger.Sugar().Errorw("validate", "Name", in.GetName())
		return fmt.Errorf("name is invalid")
	}
	if in.Logo != nil && in.GetLogo() == "" {
		logger.Sugar().Errorw("validate", "Logo", in.GetLogo())
		return fmt.Errorf("logo is invalid")
	}
	if in.GetUnit() == "" {
		logger.Sugar().Errorw("validate", "Unit", in.GetUnit())
		return fmt.Errorf("unit is invalid")
	}

	switch in.GetENV() {
	case "main":
	case "test":
	default:
		logger.Sugar().Errorw("validate", "ENV", in.GetENV())
		return fmt.Errorf("env is invalid")
	}

	if in.ReservedAmount != nil {
		if _, err := decimal.NewFromString(in.GetReservedAmount()); err != nil {
			logger.Sugar().Errorw("validate", "ReservedAmount", in.GetReservedAmount(), "error", err)
			return err
		}
	}

	return nil
}

func validateMany(in []*npool.CoinBaseReq) error {
	names := map[string]int{}

	for _, info := range in {
		if err := validate(info); err != nil {
			return err
		}
		names[info.GetName()] += 1
	}

	for k, v := range names {
		if v > 1 {
			return fmt.Errorf("name %v is invalid", k)
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

	if conds.Name != nil {
		if conds.GetName().GetValue() == "" {
			logger.Sugar().Errorw("validate", "Name", conds.GetName().GetValue())
			return fmt.Errorf("name is invalid")
		}
	}

	if conds.ENV != nil {
		switch conds.GetENV().GetValue() {
		case "main":
		case "test":
		default:
			logger.Sugar().Errorw("validate", "ENV", conds.GetENV().GetValue())
			return fmt.Errorf("env is invalid")
		}
	}

	return nil
}
