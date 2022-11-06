package appcoin

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func validate(in *npool.AppCoinReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateAppCoin", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateAppCoin", "AppID", in.GetAppID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateAppCoin", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	if in.GetName() == "" {
		logger.Sugar().Errorw("CreateAppCoin", "Name", in.GetName())
		return fmt.Errorf("name is invalid")
	}
	if in.Logo != nil && in.GetLogo() == "" {
		logger.Sugar().Errorw("CreateAppCoin", "Logo", in.GetLogo())
		return fmt.Errorf("logo is invalid")
	}
	if _, err := decimal.NewFromString(in.GetWithdrawAutoReviewAmount()); err != nil {
		logger.Sugar().Errorw("CreateAppCoin", "WithdrawAutoReviewAmount", in.GetWithdrawAutoReviewAmount(), "error", err)
		return err
	}
	return nil
}

func validateMany(in []*npool.AppCoinReq) error {
	apps := map[string]struct{}{}
	keys := map[string]int{}
	names := map[string]int{}

	for _, info := range in {
		if err := validate(info); err != nil {
			return err
		}

		apps[info.GetAppID()] = struct{}{}
		keys[fmt.Sprintf("%v:%v", info.GetAppID(), info.GetCoinTypeID())] += 1
		names[info.GetName()] += 1
	}

	if len(apps) > 1 {
		return fmt.Errorf("AppID is invalid")
	}
	for k, v := range keys {
		if v > 1 {
			return fmt.Errorf("%v is invalid", k)
		}
	}
	for k, v := range names {
		if v > 1 {
			return fmt.Errorf("%v is invalid", k)
		}
	}

	return nil
}
