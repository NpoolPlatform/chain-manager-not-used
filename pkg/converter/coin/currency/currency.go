package currency

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
)

func Ent2Grpc(row *ent.Currency) *npool.Currency {
	if row == nil {
		return nil
	}

	return &npool.Currency{
		ID:              row.ID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		FeedType:        npool.FeedType(npool.FeedType_value[row.FeedType]),
		MarketValueHigh: row.MarketValueHigh.String(),
		MarketValueLow:  row.MarketValueLow.String(),
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Currency) []*npool.Currency {
	infos := []*npool.Currency{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
