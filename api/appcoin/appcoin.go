//nolint:nolintlint,dupl
package appcoin

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/appcoin"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/appcoin"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"

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

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppCoin", "error", err)
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

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateAppCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppCoins", "error", err)
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

	if in.GetInfo().Name != nil && in.GetInfo().GetName() == "" {
		logger.Sugar().Errorw("UpdateAppCoin", "Name", in.GetInfo().GetName())
		return &npool.UpdateAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Logo != nil && in.GetInfo().GetLogo() == "" {
		logger.Sugar().Errorw("UpdateAppCoin", "Logo", in.GetInfo().GetLogo())
		return &npool.UpdateAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().WithdrawAutoReviewAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetWithdrawAutoReviewAmount()); err != nil {
			logger.Sugar().Errorw("CreateAppCoin", "WithdrawAutoReviewAmount", in.GetInfo().GetWithdrawAutoReviewAmount(), "error", err)
			return &npool.UpdateAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppCoin", "error", err)
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
		logger.Sugar().Errorw("GetAppCoin", "ID", in.GetID(), "error", err)
		return &npool.GetAppCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppCoin", "error", err)
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

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetAppCoinOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcoin", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppCoinOnly", "error", err)
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

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetAppCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppCoins", "error", err)
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
		logger.Sugar().Errorw("ExistAppCoin", "error", err)
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

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistAppCoinCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppCoinConds", "error", err)
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

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountAppCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountAppCoins", "error", err)
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
		logger.Sugar().Errorw("DeleteAppCoin", "error", err)
		return &npool.DeleteAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppCoinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
