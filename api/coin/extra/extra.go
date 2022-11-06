//nolint:nolintlint,dupl
package coinextra

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/extra"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/extra"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/extra"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"

	"github.com/google/uuid"
)

func (s *Server) CreateCoinExtra(ctx context.Context, in *npool.CreateCoinExtraRequest) (*npool.CreateCoinExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinExtra")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateCoinExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create coinextra: %v", err.Error())
		return &npool.CreateCoinExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCoinExtras(ctx context.Context, in *npool.CreateCoinExtrasRequest) (*npool.CreateCoinExtrasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinExtras")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCoinExtrasResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateCoinExtrasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "coinextra", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create coinextras: %v", err)
		return &npool.CreateCoinExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinExtrasResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCoinExtra(ctx context.Context, in *npool.UpdateCoinExtraRequest) (*npool.UpdateCoinExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCoinExtra")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create coinextra: %v", err.Error())
		return &npool.UpdateCoinExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinExtra(ctx context.Context, in *npool.GetCoinExtraRequest) (*npool.GetCoinExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinExtra")
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
		return &npool.GetCoinExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get coinextra: %v", err)
		return &npool.GetCoinExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinExtraOnly(ctx context.Context, in *npool.GetCoinExtraOnlyRequest) (*npool.GetCoinExtraOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinExtraOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCoinExtraOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get coinextras: %v", err)
		return &npool.GetCoinExtraOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinExtraOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinExtras(ctx context.Context, in *npool.GetCoinExtrasRequest) (*npool.GetCoinExtrasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinExtras")
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
		return &npool.GetCoinExtrasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get coinextras: %v", err)
		return &npool.GetCoinExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinExtrasResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCoinExtra(ctx context.Context, in *npool.ExistCoinExtraRequest) (*npool.ExistCoinExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCoinExtra")
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
		return &npool.ExistCoinExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check coinextra: %v", err)
		return &npool.ExistCoinExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinExtraResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCoinExtraConds(ctx context.Context,
	in *npool.ExistCoinExtraCondsRequest) (*npool.ExistCoinExtraCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCoinExtraConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistCoinExtraCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check coinextra: %v", err)
		return &npool.ExistCoinExtraCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinExtraCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCoinExtras(ctx context.Context, in *npool.CountCoinExtrasRequest) (*npool.CountCoinExtrasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCoinExtras")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountCoinExtrasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count coinextras: %v", err)
		return &npool.CountCoinExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCoinExtrasResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCoinExtra(ctx context.Context, in *npool.DeleteCoinExtraRequest) (*npool.DeleteCoinExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCoinExtra")
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
		return &npool.DeleteCoinExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coinextra", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete coinextra: %v", err)
		return &npool.DeleteCoinExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
