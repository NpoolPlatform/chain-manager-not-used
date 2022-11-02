//nolint:nolintlint,dupl
package coinbase

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/base"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/base"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/base"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	"github.com/google/uuid"
)

func (s *Server) CreateCoinBase(ctx context.Context, in *npool.CreateCoinBaseRequest) (*npool.CreateCoinBaseResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinBase")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create coinbase: %v", err.Error())
		return &npool.CreateCoinBaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinBaseResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCoinBases(ctx context.Context, in *npool.CreateCoinBasesRequest) (*npool.CreateCoinBasesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinBases")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCoinBasesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	// TODO: verify infput

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "coinbase", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create coinbases: %v", err)
		return &npool.CreateCoinBasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinBasesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCoinBase(ctx context.Context, in *npool.UpdateCoinBaseRequest) (*npool.UpdateCoinBaseResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCoinBase")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create coinbase: %v", err.Error())
		return &npool.UpdateCoinBaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinBaseResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinBase(ctx context.Context, in *npool.GetCoinBaseRequest) (*npool.GetCoinBaseResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinBase")
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
		return &npool.GetCoinBaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get coinbase: %v", err)
		return &npool.GetCoinBaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinBaseResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinBaseOnly(ctx context.Context, in *npool.GetCoinBaseOnlyRequest) (*npool.GetCoinBaseOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinBaseOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "coinbase", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get coinbases: %v", err)
		return &npool.GetCoinBaseOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinBaseOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinBases(ctx context.Context, in *npool.GetCoinBasesRequest) (*npool.GetCoinBasesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinBases")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get coinbases: %v", err)
		return &npool.GetCoinBasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinBasesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCoinBase(ctx context.Context, in *npool.ExistCoinBaseRequest) (*npool.ExistCoinBaseResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCoinBase")
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
		return &npool.ExistCoinBaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check coinbase: %v", err)
		return &npool.ExistCoinBaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinBaseResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCoinBaseConds(ctx context.Context,
	in *npool.ExistCoinBaseCondsRequest) (*npool.ExistCoinBaseCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCoinBaseConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "coinbase", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check coinbase: %v", err)
		return &npool.ExistCoinBaseCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinBaseCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCoinBases(ctx context.Context, in *npool.CountCoinBasesRequest) (*npool.CountCoinBasesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCoinBases")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count coinbases: %v", err)
		return &npool.CountCoinBasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCoinBasesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCoinBase(ctx context.Context, in *npool.DeleteCoinBaseRequest) (*npool.DeleteCoinBaseResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCoinBase")
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
		return &npool.DeleteCoinBaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinbase", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete coinbase: %v", err)
		return &npool.DeleteCoinBaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinBaseResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
