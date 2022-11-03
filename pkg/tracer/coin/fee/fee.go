package fee

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fee"
)

func trace(span trace1.Span, in *npool.FeeReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("FeeCoinTypeID.%v", index), in.GetFeeCoinTypeID()),
		attribute.Bool(fmt.Sprintf("WithdrawFeeByStableUSD.%v", index), in.GetWithdrawFeeByStableUSD()),
		attribute.String(fmt.Sprintf("WithdrawFeeAmount.%v", index), in.GetWithdrawFeeAmount()),
		attribute.String(fmt.Sprintf("CollectFeeAmount.%v", index), in.GetCollectFeeAmount()),
		attribute.String(fmt.Sprintf("HotWalletFeeAmount.%v", index), in.GetHotWalletFeeAmount()),
		attribute.String(fmt.Sprintf("LowFeeAmount.%v", index), in.GetLowFeeAmount()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.FeeReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("FeeCoinTypeID.Op", in.GetFeeCoinTypeID().GetOp()),
		attribute.String("FeeCoinTypeID.Value", in.GetFeeCoinTypeID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.FeeReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
