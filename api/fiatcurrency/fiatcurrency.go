//nolint:nolintlint,dupl
package fiatcurrency

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/fiatcurrency"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/fiatcurrency"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiatcurrency"

	"github.com/google/uuid"
)

func (s *Server) CreateFiatCurrency(
	ctx context.Context,
	in *npool.CreateFiatCurrencyRequest,
) (
	*npool.CreateFiatCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFiatCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fiatcurrency: %v", err.Error())
		return &npool.CreateFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateFiatCurrencies(
	ctx context.Context,
	in *npool.CreateFiatCurrenciesRequest,
) (
	*npool.CreateFiatCurrenciesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFiatCurrencies")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create fiatcurrencys: %v", err)
		return &npool.CreateFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatCurrenciesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateFiatCurrency(
	ctx context.Context,
	in *npool.UpdateFiatCurrencyRequest,
) (
	*npool.UpdateFiatCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateFiatCurrency")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateUpdate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("UpdateFiatCurrency", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fiatcurrency: %v", err.Error())
		return &npool.UpdateFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFiatCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFiatCurrency(
	ctx context.Context,
	in *npool.GetFiatCurrencyRequest,
) (
	*npool.GetFiatCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFiatCurrency")
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
		return &npool.GetFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get fiatcurrency: %v", err)
		return &npool.GetFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFiatCurrencyOnly(
	ctx context.Context,
	in *npool.GetFiatCurrencyOnlyRequest,
) (
	*npool.GetFiatCurrencyOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFiatCurrencyOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetFiatCurrencyOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get fiatcurrencys: %v", err)
		return &npool.GetFiatCurrencyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFiatCurrencies(ctx context.Context, in *npool.GetFiatCurrenciesRequest) (*npool.GetFiatCurrenciesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFiatCurrencies")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get fiatcurrencys: %v", err)
		return &npool.GetFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrenciesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistFiatCurrency(ctx context.Context, in *npool.ExistFiatCurrencyRequest) (*npool.ExistFiatCurrencyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFiatCurrency")
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
		return &npool.ExistFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check fiatcurrency: %v", err)
		return &npool.ExistFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFiatCurrencyResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFiatCurrencyConds(ctx context.Context,
	in *npool.ExistFiatCurrencyCondsRequest) (*npool.ExistFiatCurrencyCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFiatCurrencyConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistFiatCurrencyCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check fiatcurrency: %v", err)
		return &npool.ExistFiatCurrencyCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFiatCurrencyCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountFiatCurrencies(
	ctx context.Context,
	in *npool.CountFiatCurrenciesRequest,
) (
	*npool.CountFiatCurrenciesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountFiatCurrencies")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count fiatcurrencys: %v", err)
		return &npool.CountFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountFiatCurrenciesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteFiatCurrency(
	ctx context.Context,
	in *npool.DeleteFiatCurrencyRequest,
) (
	*npool.DeleteFiatCurrencyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteFiatCurrency")
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
		return &npool.DeleteFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete fiatcurrency: %v", err)
		return &npool.DeleteFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFiatCurrencyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
