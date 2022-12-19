package fiatcurrencytype

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiat/currencytype"
)

func Ent2Grpc(row *ent.FiatCurrencyType) *npool.FiatCurrencyType {
	if row == nil {
		return nil
	}

	return &npool.FiatCurrencyType{
		ID:        row.ID.String(),
		Name:      row.Name,
		Logo:      row.Logo,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.FiatCurrencyType) []*npool.FiatCurrencyType {
	infos := []*npool.FiatCurrencyType{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
