package currencyfeed

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
)

func Ent2Grpc(row *ent.CurrencyFeed) *npool.CurrencyFeed {
	if row == nil {
		return nil
	}

	return &npool.CurrencyFeed{
		ID:         row.ID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		FeedSource: row.FeedSource,
		FeedType:   npool.FeedType(npool.FeedType_value[row.FeedType]),
		Disabled:   row.Disabled,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CurrencyFeed) []*npool.CurrencyFeed {
	infos := []*npool.CurrencyFeed{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
