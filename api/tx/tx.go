//nolint:nolintlint,dupl
package tx

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/tx"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/tx"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/tx"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"

	"github.com/google/uuid"
)

func (s *Server) CreateTx(ctx context.Context, in *npool.CreateTxRequest) (*npool.CreateTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "tx", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create tx: %v", err.Error())
		return &npool.CreateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTxs(ctx context.Context, in *npool.CreateTxsRequest) (*npool.CreateTxsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateTxsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	// TODO: verify infput

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "tx", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create txs: %v", err)
		return &npool.CreateTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateTx(ctx context.Context, in *npool.UpdateTxRequest) (*npool.UpdateTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "tx", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update tx: %v", err.Error())
		return &npool.UpdateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTx(ctx context.Context, in *npool.GetTxRequest) (*npool.GetTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTx")
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
		return &npool.GetTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get tx: %v", err)
		return &npool.GetTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxOnly(ctx context.Context, in *npool.GetTxOnlyRequest) (*npool.GetTxOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get txs: %v", err)
		return &npool.GetTxOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxs(ctx context.Context, in *npool.GetTxsRequest) (*npool.GetTxsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "tx", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get txs: %v", err)
		return &npool.GetTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistTx(ctx context.Context, in *npool.ExistTxRequest) (*npool.ExistTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTx")
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
		return &npool.ExistTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check tx: %v", err)
		return &npool.ExistTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTxConds(ctx context.Context,
	in *npool.ExistTxCondsRequest) (*npool.ExistTxCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTxConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check tx: %v", err)
		return &npool.ExistTxCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTxs(ctx context.Context, in *npool.CountTxsRequest) (*npool.CountTxsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count txs: %v", err)
		return &npool.CountTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTxsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteTx(ctx context.Context, in *npool.DeleteTxRequest) (*npool.DeleteTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteTx")
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
		return &npool.DeleteTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete tx: %v", err)
		return &npool.DeleteTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
