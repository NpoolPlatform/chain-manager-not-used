package tx

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
)

func trace(span trace1.Span, in *npool.TxReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("FromAccountID.%v", index), in.GetFromAccountID()),
		attribute.String(fmt.Sprintf("ToAccountID.%v", index), in.GetToAccountID()),
		attribute.String(fmt.Sprintf("Amount.%v", index), in.GetAmount()),
		attribute.String(fmt.Sprintf("FeeAmount.%v", index), in.GetFeeAmount()),
		attribute.String(fmt.Sprintf("ChainTxID.%v", index), in.GetChainTxID()),
		attribute.String(fmt.Sprintf("State.%v", index), in.GetState().String()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.TxReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AccountID.Op", in.GetAccountID().GetOp()),
		attribute.String("AccountID.Value", in.GetAccountID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.TxReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
