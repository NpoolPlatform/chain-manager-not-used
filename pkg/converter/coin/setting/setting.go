package setting

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"
)

func Ent2Grpc(row *ent.Setting) *npool.Setting {
	if row == nil {
		return nil
	}

	return &npool.Setting{
		ID:                          row.ID.String(),
		CoinTypeID:                  row.CoinTypeID.String(),
		FeeCoinTypeID:               row.FeeCoinTypeID.String(),
		WithdrawFeeByStableUSD:      row.WithdrawFeeByStableUsd,
		WithdrawFeeAmount:           row.WithdrawFeeAmount.String(),
		CollectFeeAmount:            row.CollectFeeAmount.String(),
		HotWalletFeeAmount:          row.HotWalletFeeAmount.String(),
		LowFeeAmount:                row.LowFeeAmount.String(),
		WarmAccountAmount:           row.WarmAccountAmount.String(),
		PaymentAccountCollectAmount: row.PaymentAccountCollectAmount.String(),
		CreatedAt:                   row.CreatedAt,
		UpdatedAt:                   row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Setting) []*npool.Setting {
	infos := []*npool.Setting{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
