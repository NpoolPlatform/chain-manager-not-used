//nolint:nolintlint,dupl
package exrate

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/exrate"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/exrate"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/exrate"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/exrate"

	"github.com/google/uuid"
)

func (s *Server) CreateExchangeRate(
	ctx context.Context,
	in *npool.CreateExchangeRateRequest,
) (
	*npool.CreateExchangeRateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateExchangeRate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "exrate", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create exrate: %v", err.Error())
		return &npool.CreateExchangeRateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateExchangeRateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateExchangeRates(
	ctx context.Context,
	in *npool.CreateExchangeRatesRequest,
) (
	*npool.CreateExchangeRatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateExchangeRates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateExchangeRatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	// TODO: verify infput

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "exrate", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create exrates: %v", err)
		return &npool.CreateExchangeRatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateExchangeRatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateExchangeRate(ctx context.Context, in *npool.UpdateExchangeRateRequest) (*npool.UpdateExchangeRateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateExchangeRate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "exrate", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create exrate: %v", err.Error())
		return &npool.UpdateExchangeRateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateExchangeRateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetExchangeRate(ctx context.Context, in *npool.GetExchangeRateRequest) (*npool.GetExchangeRateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetExchangeRate")
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
		return &npool.GetExchangeRateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "exrate", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get exrate: %v", err)
		return &npool.GetExchangeRateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetExchangeRateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetExchangeRateOnly(
	ctx context.Context,
	in *npool.GetExchangeRateOnlyRequest,
) (
	*npool.GetExchangeRateOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetExchangeRateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "exrate", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get exrates: %v", err)
		return &npool.GetExchangeRateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetExchangeRateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetExchangeRates(ctx context.Context, in *npool.GetExchangeRatesRequest) (*npool.GetExchangeRatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetExchangeRates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "exrate", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get exrates: %v", err)
		return &npool.GetExchangeRatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetExchangeRatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistExchangeRate(ctx context.Context, in *npool.ExistExchangeRateRequest) (*npool.ExistExchangeRateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistExchangeRate")
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
		return &npool.ExistExchangeRateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "exrate", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check exrate: %v", err)
		return &npool.ExistExchangeRateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistExchangeRateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistExchangeRateConds(ctx context.Context,
	in *npool.ExistExchangeRateCondsRequest) (*npool.ExistExchangeRateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistExchangeRateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "exrate", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check exrate: %v", err)
		return &npool.ExistExchangeRateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistExchangeRateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountExchangeRates(ctx context.Context, in *npool.CountExchangeRatesRequest) (*npool.CountExchangeRatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountExchangeRates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "exrate", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count exrates: %v", err)
		return &npool.CountExchangeRatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountExchangeRatesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteExchangeRate(ctx context.Context, in *npool.DeleteExchangeRateRequest) (*npool.DeleteExchangeRateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteExchangeRate")
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
		return &npool.DeleteExchangeRateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "exrate", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete exrate: %v", err)
		return &npool.DeleteExchangeRateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteExchangeRateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
