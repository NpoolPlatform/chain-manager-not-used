package appcoin

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"
)

func Ent2Grpc(row *ent.AppCoin) *npool.AppCoin {
	if row == nil {
		return nil
	}

	return &npool.AppCoin{
		ID:                       row.ID.String(),
		AppID:                    row.AppID.String(),
		CoinTypeID:               row.CoinTypeID.String(),
		Name:                     row.Name,
		DisplayNames:             row.DisplayNames,
		Logo:                     row.Logo,
		ForPay:                   row.ForPay,
		WithdrawAutoReviewAmount: row.WithdrawAutoReviewAmount.String(),
		ProductPage:              row.ProductPage,
		DailyRewardAmount:        row.DailyRewardAmount.String(),
		CreatedAt:                row.CreatedAt,
		UpdatedAt:                row.UpdatedAt,
		Disabled:                 row.Disabled,
		Display:                  row.Display,
	}
}

func Ent2GrpcMany(rows []*ent.AppCoin) []*npool.AppCoin {
	infos := []*npool.AppCoin{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
