package fiatcurrencytype

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrencytype"

	"github.com/google/uuid"
)

func validate(in *npool.FiatCurrencyTypeReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if in.Name != nil && in.GetName() == "" {
		logger.Sugar().Errorw("validate", "Name", in.GetName())
		return fmt.Errorf("Name is empty")
	}

	return nil
}

func validateMany(in []*npool.FiatCurrencyTypeReq) error {
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
	if conds.Name != nil {
		if _, err := uuid.Parse(conds.GetName().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "FiatTypeID", conds.GetName().GetValue(), "error", err)
			return err
		}
	}

	return nil
}

func validateUpdate(in *npool.FiatCurrencyTypeReq) error {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return err
	}

	if in.GetName() == "" {
		logger.Sugar().Errorw("validate", "Name", in.GetName())
		return fmt.Errorf("Name is empty")
	}
	return nil
}
