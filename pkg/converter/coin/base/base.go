package coinbase

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"
)

func Ent2Grpc(row *ent.CoinBase) *npool.CoinBase {
	if row == nil {
		return nil
	}

	return &npool.CoinBase{
		ID:             row.ID.String(),
		Name:           row.Name,
		Logo:           row.Logo,
		Presale:        row.Presale,
		Unit:           row.Unit,
		ENV:            row.Env,
		ReservedAmount: row.ReservedAmount.String(),
		ForPay:         row.ForPay,
		Disabled:       row.Disabled,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CoinBase) []*npool.CoinBase {
	infos := []*npool.CoinBase{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
