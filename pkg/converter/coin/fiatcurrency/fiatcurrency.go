package fiatcurrency

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiatcurrency"
)

func Ent2Grpc(row *ent.FiatCurrency) *npool.FiatCurrency {
	if row == nil {
		return nil
	}

	return &npool.FiatCurrency{
		ID:                 row.ID.String(),
		FiatCurrencyTypeID: row.FiatCurrencyTypeID.String(),
		FeedType:           currency.FeedType(currency.FeedType_value[row.FeedType]),
		MarketValueHigh:    row.MarketValueHigh.String(),
		MarketValueLow:     row.MarketValueLow.String(),
		CreatedAt:          row.CreatedAt,
		UpdatedAt:          row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.FiatCurrency) []*npool.FiatCurrency {
	infos := []*npool.FiatCurrency{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
