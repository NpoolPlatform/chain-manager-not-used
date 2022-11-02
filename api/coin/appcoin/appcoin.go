//nolint:nolintlint,dupl
package appcoin

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/appcoin"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/appcoin"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/appcoin"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/appcoin"

	"github.com/google/uuid"
)

func (s *Server) CreateAppCoin(ctx context.Context, in *npool.CreateAppCoinRequest) (*npool.CreateAppCoinResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create appcoin: %v", err.Error())
		return &npool.CreateAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppCoinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppCoins(ctx context.Context, in *npool.CreateAppCoinsRequest) (*npool.CreateAppCoinsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppCoins")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppCoinsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	// TODO: verify infput

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create appcoins: %v", err)
		return &npool.CreateAppCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppCoinsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppCoin(ctx context.Context, in *npool.UpdateAppCoinRequest) (*npool.UpdateAppCoinResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create appcoin: %v", err.Error())
		return &npool.UpdateAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppCoinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppCoin(ctx context.Context, in *npool.GetAppCoinRequest) (*npool.GetAppCoinResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppCoin")
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
		return &npool.GetAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get appcoin: %v", err)
		return &npool.GetAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppCoinOnly(ctx context.Context, in *npool.GetAppCoinOnlyRequest) (*npool.GetAppCoinOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppCoinOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get appcoins: %v", err)
		return &npool.GetAppCoinOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppCoins(ctx context.Context, in *npool.GetAppCoinsRequest) (*npool.GetAppCoinsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppCoins")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get appcoins: %v", err)
		return &npool.GetAppCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppCoin(ctx context.Context, in *npool.ExistAppCoinRequest) (*npool.ExistAppCoinResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppCoin")
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
		return &npool.ExistAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check appcoin: %v", err)
		return &npool.ExistAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppCoinResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppCoinConds(ctx context.Context,
	in *npool.ExistAppCoinCondsRequest) (*npool.ExistAppCoinCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppCoinConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check appcoin: %v", err)
		return &npool.ExistAppCoinCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppCoinCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppCoins(ctx context.Context, in *npool.CountAppCoinsRequest) (*npool.CountAppCoinsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppCoins")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count appcoins: %v", err)
		return &npool.CountAppCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppCoinsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppCoin(ctx context.Context, in *npool.DeleteAppCoinRequest) (*npool.DeleteAppCoinResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppCoin")
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
		return &npool.DeleteAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete appcoin: %v", err)
		return &npool.DeleteAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppCoinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
