package description

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
)

func Ent2Grpc(row *ent.CoinDescription) *npool.CoinDescription {
	if row == nil {
		return nil
	}

	return &npool.CoinDescription{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		UsedFor:    npool.UsedFor(npool.UsedFor_value[row.UsedFor]),
		Title:      row.Title,
		Message:    row.Message,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CoinDescription) []*npool.CoinDescription {
	infos := []*npool.CoinDescription{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
