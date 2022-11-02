package appcoin

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/appcoin"
)

func trace(span trace1.Span, in *npool.AppCoinReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
		attribute.String(fmt.Sprintf("Logo.%v", index), in.GetLogo()),
		attribute.Bool(fmt.Sprintf("ForPay.%v", index), in.GetForPay()),
		attribute.String(fmt.Sprintf("WithdrawAutoReviewAmount.%v", index), in.GetWithdrawAutoReviewAmount()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppCoinReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppCoinReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
