package currencyvalue

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
)

func Ent2Grpc(row *ent.CurrencyValue) *npool.CurrencyValue {
	if row == nil {
		return nil
	}

	return &npool.CurrencyValue{
		ID:              row.ID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		FeedSourceID:    row.FeedSourceID.String(),
		MarketValueHigh: row.MarketValueHigh.String(),
		MarketValueLow:  row.MarketValueLow.String(),
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CurrencyValue) []*npool.CurrencyValue {
	infos := []*npool.CurrencyValue{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
