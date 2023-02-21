package tx

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
)

func Ent2Grpc(row *ent.Tran) *npool.Tx {
	if row == nil {
		return nil
	}

	return &npool.Tx{
		ID:            row.ID.String(),
		CoinTypeID:    row.CoinTypeID.String(),
		FromAccountID: row.FromAccountID.String(),
		ToAccountID:   row.ToAccountID.String(),
		Amount:        row.Amount.String(),
		FeeAmount:     row.FeeAmount.String(),
		ChainTxID:     row.ChainTxID,
		State:         npool.TxState(npool.TxState_value[row.State]),
		Extra:         row.Extra,
		Type:          basetypes.TxType(basetypes.TxType_value[row.Type]),
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Tran) []*npool.Tx {
	infos := []*npool.Tx{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
