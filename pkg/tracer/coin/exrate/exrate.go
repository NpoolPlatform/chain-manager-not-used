package exrate

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/exrate"
)

func trace(span trace1.Span, in *npool.ExchangeRateReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("MarketValue.%v", index), in.GetMarketValue()),
		attribute.String(fmt.Sprintf("SettleValue.%v", index), in.GetSettleValue()),
		attribute.Int(fmt.Sprintf("SettlePercent.%v", index), int(in.GetSettlePercent())),
		attribute.String(fmt.Sprintf("Setter.%v", index), in.GetSetter()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.ExchangeRateReq) trace1.Span {
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

func TraceMany(span trace1.Span, infos []*npool.ExchangeRateReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
