//nolint:nolintlint,dupl
package currency

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/currency"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	"github.com/google/uuid"
)

func (s *Server) CreateCurrency(
	ctx context.Context,
	in *npool.CreateCurrencyRequest,
) (
	*npool.CreateCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currency: %v", err.Error())
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCurrencies(
	ctx context.Context,
	in *npool.CreateCurrenciesRequest,
) (
	*npool.CreateCurrenciesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCurrencies")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCurrenciesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create currencys: %v", err)
		return &npool.CreateCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrenciesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCurrency(
	ctx context.Context,
	in *npool.UpdateCurrencyRequest,
) (
	*npool.UpdateCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateUpdate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("UpdateCurrency", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currency: %v", err.Error())
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrency(
	ctx context.Context,
	in *npool.GetCurrencyRequest,
) (
	*npool.GetCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get currency: %v", err)
		return &npool.GetCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyOnly(
	ctx context.Context,
	in *npool.GetCurrencyOnlyRequest,
) (
	*npool.GetCurrencyOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCurrencyOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get currencys: %v", err)
		return &npool.GetCurrencyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencies")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get currencys: %v", err)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrenciesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCurrency(ctx context.Context, in *npool.ExistCurrencyRequest) (*npool.ExistCurrencyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check currency: %v", err)
		return &npool.ExistCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCurrencyResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCurrencyConds(ctx context.Context,
	in *npool.ExistCurrencyCondsRequest) (*npool.ExistCurrencyCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCurrencyConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistCurrencyCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check currency: %v", err)
		return &npool.ExistCurrencyCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCurrencyCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCurrencies(
	ctx context.Context,
	in *npool.CountCurrenciesRequest,
) (
	*npool.CountCurrenciesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCurrencies")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count currencys: %v", err)
		return &npool.CountCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCurrenciesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCurrency(
	ctx context.Context,
	in *npool.DeleteCurrencyRequest,
) (
	*npool.DeleteCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currency", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete currency: %v", err)
		return &npool.DeleteCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
