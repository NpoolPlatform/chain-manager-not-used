package fee

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fee"
)

func Ent2Grpc(row *ent.Fee) *npool.Fee {
	if row == nil {
		return nil
	}

	return &npool.Fee{
		ID:                 row.ID.String(),
		CoinTypeID:         row.CoinTypeID.String(),
		FeeCoinTypeID:      row.FeeCoinTypeID.String(),
		WithdrawFeeAmount:  row.WithdrawFeeAmount.String(),
		CollectFeeAmount:   row.CollectFeeAmount.String(),
		HotWalletFeeAmount: row.HotWalletFeeAmount.String(),
		LowFeeAmount:       row.LowFeeAmount.String(),
		CreatedAt:          row.CreatedAt,
		UpdatedAt:          row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Fee) []*npool.Fee {
	infos := []*npool.Fee{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
