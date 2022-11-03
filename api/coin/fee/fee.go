//nolint:nolintlint,dupl
package fee

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/fee"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/fee"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/fee"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fee"

	"github.com/google/uuid"
)

func (s *Server) CreateFee(
	ctx context.Context,
	in *npool.CreateFeeRequest,
) (
	*npool.CreateFeeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFee")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "fee", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fee: %v", err.Error())
		return &npool.CreateFeeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFeeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateFees(
	ctx context.Context,
	in *npool.CreateFeesRequest,
) (
	*npool.CreateFeesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFees")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateFeesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	// TODO: verify infput

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "fee", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create fees: %v", err)
		return &npool.CreateFeesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFeesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateFee(ctx context.Context, in *npool.UpdateFeeRequest) (*npool.UpdateFeeResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateFee")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "fee", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fee: %v", err.Error())
		return &npool.UpdateFeeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFeeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFee(ctx context.Context, in *npool.GetFeeRequest) (*npool.GetFeeResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFee")
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
		return &npool.GetFeeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fee", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get fee: %v", err)
		return &npool.GetFeeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFeeOnly(
	ctx context.Context,
	in *npool.GetFeeOnlyRequest,
) (
	*npool.GetFeeOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFeeOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "fee", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get fees: %v", err)
		return &npool.GetFeeOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeeOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFees(ctx context.Context, in *npool.GetFeesRequest) (*npool.GetFeesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFees")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "fee", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get fees: %v", err)
		return &npool.GetFeesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistFee(ctx context.Context, in *npool.ExistFeeRequest) (*npool.ExistFeeResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFee")
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
		return &npool.ExistFeeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fee", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check fee: %v", err)
		return &npool.ExistFeeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFeeResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFeeConds(ctx context.Context,
	in *npool.ExistFeeCondsRequest) (*npool.ExistFeeCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFeeConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "fee", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check fee: %v", err)
		return &npool.ExistFeeCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFeeCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountFees(ctx context.Context, in *npool.CountFeesRequest) (*npool.CountFeesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountFees")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "fee", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count fees: %v", err)
		return &npool.CountFeesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountFeesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteFee(ctx context.Context, in *npool.DeleteFeeRequest) (*npool.DeleteFeeResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteFee")
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
		return &npool.DeleteFeeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fee", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete fee: %v", err)
		return &npool.DeleteFeeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFeeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
