package exrate

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/exrate"
)

func Ent2Grpc(row *ent.ExchangeRate) *npool.ExchangeRate {
	if row == nil {
		return nil
	}

	return &npool.ExchangeRate{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		CoinTypeID:    row.CoinTypeID.String(),
		MarketValue:   row.MarketValue.String(),
		SettleValue:   row.SettleValue.String(),
		SettlePercent: row.SettlePercent,
		Setter:        row.Setter.String(),
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ExchangeRate) []*npool.ExchangeRate {
	infos := []*npool.ExchangeRate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
