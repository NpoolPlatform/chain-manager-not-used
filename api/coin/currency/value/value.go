//nolint:nolintlint,dupl
package currencyvalue

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/currency/value"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/value"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"

	"github.com/google/uuid"
)

func (s *Server) CreateCurrencyValue(
	ctx context.Context,
	in *npool.CreateCurrencyValueRequest,
) (
	*npool.CreateCurrencyValueResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCurrencyValue")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateCurrencyValueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currencyvalue: %v", err.Error())
		return &npool.CreateCurrencyValueResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyValueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCurrencyValues(
	ctx context.Context,
	in *npool.CreateCurrencyValuesRequest,
) (
	*npool.CreateCurrencyValuesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCurrencyValues")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCurrencyValuesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateCurrencyValuesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create currencyvalues: %v", err)
		return &npool.CreateCurrencyValuesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyValuesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCurrencyValue(
	ctx context.Context,
	in *npool.UpdateCurrencyValueRequest,
) (
	*npool.UpdateCurrencyValueResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCurrencyValue")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateUpdate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("UpdateCurrencyValue", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateCurrencyValueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currencyvalue: %v", err.Error())
		return &npool.UpdateCurrencyValueResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCurrencyValueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyValue(
	ctx context.Context,
	in *npool.GetCurrencyValueRequest,
) (
	*npool.GetCurrencyValueResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyValue")
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
		return &npool.GetCurrencyValueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get currencyvalue: %v", err)
		return &npool.GetCurrencyValueResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyValueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyValueOnly(
	ctx context.Context,
	in *npool.GetCurrencyValueOnlyRequest,
) (
	*npool.GetCurrencyValueOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyValueOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCurrencyValueOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get currencyvalues: %v", err)
		return &npool.GetCurrencyValueOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyValueOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyValues(ctx context.Context, in *npool.GetCurrencyValuesRequest) (*npool.GetCurrencyValuesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyValues")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCurrencyValuesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get currencyvalues: %v", err)
		return &npool.GetCurrencyValuesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyValuesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCurrencyValue(ctx context.Context, in *npool.ExistCurrencyValueRequest) (*npool.ExistCurrencyValueResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCurrencyValue")
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
		return &npool.ExistCurrencyValueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check currencyvalue: %v", err)
		return &npool.ExistCurrencyValueResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCurrencyValueResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCurrencyValueConds(ctx context.Context,
	in *npool.ExistCurrencyValueCondsRequest) (*npool.ExistCurrencyValueCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCurrencyValueConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistCurrencyValueCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check currencyvalue: %v", err)
		return &npool.ExistCurrencyValueCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCurrencyValueCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCurrencyValues(
	ctx context.Context,
	in *npool.CountCurrencyValuesRequest,
) (
	*npool.CountCurrencyValuesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCurrencyValues")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountCurrencyValuesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count currencyvalues: %v", err)
		return &npool.CountCurrencyValuesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCurrencyValuesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCurrencyValue(
	ctx context.Context,
	in *npool.DeleteCurrencyValueRequest,
) (
	*npool.DeleteCurrencyValueResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCurrencyValue")
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
		return &npool.DeleteCurrencyValueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyvalue", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete currencyvalue: %v", err)
		return &npool.DeleteCurrencyValueResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCurrencyValueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
