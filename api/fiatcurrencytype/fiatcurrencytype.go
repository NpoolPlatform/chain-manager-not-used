//nolint:nolintlint,dupl
package fiatcurrencytype

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/fiatcurrencytype"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/fiatcurrencytype"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrencytype"

	"github.com/google/uuid"
)

func (s *Server) CreateFiatCurrencyType(
	ctx context.Context,
	in *npool.CreateFiatCurrencyTypeRequest,
) (
	*npool.CreateFiatCurrencyTypeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFiatCurrencyType")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateFiatCurrencyTypeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fiatcurrency: %v", err.Error())
		return &npool.CreateFiatCurrencyTypeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatCurrencyTypeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateFiatCurrencyTypes(
	ctx context.Context,
	in *npool.CreateFiatCurrencyTypesRequest,
) (
	*npool.CreateFiatCurrencyTypesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFiatCurrencyTypes")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateFiatCurrencyTypesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateFiatCurrencyTypesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create fiatcurrencys: %v", err)
		return &npool.CreateFiatCurrencyTypesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatCurrencyTypesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateFiatCurrencyType(
	ctx context.Context,
	in *npool.UpdateFiatCurrencyTypeRequest,
) (
	*npool.UpdateFiatCurrencyTypeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateFiatCurrencyType")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateUpdate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("UpdateFiatCurrencyType", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateFiatCurrencyTypeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fiatcurrency: %v", err.Error())
		return &npool.UpdateFiatCurrencyTypeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFiatCurrencyTypeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFiatCurrencyType(
	ctx context.Context,
	in *npool.GetFiatCurrencyTypeRequest,
) (
	*npool.GetFiatCurrencyTypeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFiatCurrencyType")
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
		return &npool.GetFiatCurrencyTypeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get fiatcurrency: %v", err)
		return &npool.GetFiatCurrencyTypeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyTypeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFiatCurrencyTypeOnly(
	ctx context.Context,
	in *npool.GetFiatCurrencyTypeOnlyRequest,
) (
	*npool.GetFiatCurrencyTypeOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFiatCurrencyTypeOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetFiatCurrencyTypeOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get fiatcurrencys: %v", err)
		return &npool.GetFiatCurrencyTypeOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyTypeOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFiatCurrencyTypes(
	ctx context.Context,
	in *npool.GetFiatCurrencyTypesRequest,
) (
	*npool.GetFiatCurrencyTypesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFiatCurrencyTypes")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetFiatCurrencyTypesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get fiatcurrencys: %v", err)
		return &npool.GetFiatCurrencyTypesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyTypesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistFiatCurrencyType(
	ctx context.Context,
	in *npool.ExistFiatCurrencyTypeRequest,
) (
	*npool.ExistFiatCurrencyTypeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFiatCurrencyType")
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
		return &npool.ExistFiatCurrencyTypeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check fiatcurrency: %v", err)
		return &npool.ExistFiatCurrencyTypeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFiatCurrencyTypeResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFiatCurrencyTypeConds(ctx context.Context,
	in *npool.ExistFiatCurrencyTypeCondsRequest) (*npool.ExistFiatCurrencyTypeCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFiatCurrencyTypeConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistFiatCurrencyTypeCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check fiatcurrency: %v", err)
		return &npool.ExistFiatCurrencyTypeCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFiatCurrencyTypeCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountFiatCurrencyTypes(
	ctx context.Context,
	in *npool.CountFiatCurrencyTypesRequest,
) (
	*npool.CountFiatCurrencyTypesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountFiatCurrencyTypes")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountFiatCurrencyTypesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count fiatcurrencys: %v", err)
		return &npool.CountFiatCurrencyTypesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountFiatCurrencyTypesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteFiatCurrencyType(
	ctx context.Context,
	in *npool.DeleteFiatCurrencyTypeRequest,
) (
	*npool.DeleteFiatCurrencyTypeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteFiatCurrencyType")
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
		return &npool.DeleteFiatCurrencyTypeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fiatcurrency", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete fiatcurrency: %v", err)
		return &npool.DeleteFiatCurrencyTypeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFiatCurrencyTypeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
