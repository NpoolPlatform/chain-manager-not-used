package coinbase

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"
)

func trace(span trace1.Span, in *npool.CoinBaseReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
		attribute.String(fmt.Sprintf("Logo.%v", index), in.GetLogo()),
		attribute.Bool(fmt.Sprintf("Presale.%v", index), in.GetPresale()),
		attribute.String(fmt.Sprintf("Unit.%v", index), in.GetUnit()),
		attribute.String(fmt.Sprintf("ENV.%v", index), in.GetENV()),
		attribute.String(fmt.Sprintf("ReservedAmount.%v", index), in.GetReservedAmount()),
		attribute.Bool(fmt.Sprintf("ForPay.%v", index), in.GetForPay()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.CoinBaseReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("Name.Op", in.GetName().GetOp()),
		attribute.String("Name.Value", in.GetName().GetValue()),
		attribute.String("ENV.Op", in.GetENV().GetOp()),
		attribute.String("ENV.Value", in.GetENV().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.CoinBaseReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
