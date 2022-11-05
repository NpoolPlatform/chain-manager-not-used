package coinextra

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"
)

func Ent2Grpc(row *ent.CoinExtra) *npool.CoinExtra {
	if row == nil {
		return nil
	}

	return &npool.CoinExtra{
		ID:         row.ID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		HomePage:   row.HomePage,
		Specs:      row.Specs,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CoinExtra) []*npool.CoinExtra {
	infos := []*npool.CoinExtra{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
