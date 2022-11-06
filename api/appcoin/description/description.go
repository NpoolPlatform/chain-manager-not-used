//nolint:nolintlint,dupl
package description

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/appcoin/description"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin/description"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/appcoin/description"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"

	"github.com/google/uuid"
)

func (s *Server) CreateCoinDescription(
	ctx context.Context,
	in *npool.CreateCoinDescriptionRequest,
) (
	*npool.CreateCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinDescription")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "error", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinDescriptionResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCoinDescriptions(
	ctx context.Context,
	in *npool.CreateCoinDescriptionsRequest,
) (
	*npool.CreateCoinDescriptionsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinDescriptions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "description", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateCoinDescriptions", "error", err)
		return &npool.CreateCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinDescriptionsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCoinDescription(
	ctx context.Context,
	in *npool.UpdateCoinDescriptionRequest,
) (
	*npool.UpdateCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCoinDescription")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if in.GetInfo().Title != nil && in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("UpdateCoinDescription", "Title", in.GetInfo().GetTitle())
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
	}
	if in.GetInfo().Message != nil && in.GetInfo().GetMessage() == "" {
		logger.Sugar().Errorw("UpdateCoinDescription", "Message", in.GetInfo().GetMessage())
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "Message is invalid")
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoinDescription", "error", err)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinDescription")
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
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetCoinDescription", "error", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinDescriptionOnly(
	ctx context.Context,
	in *npool.GetCoinDescriptionOnlyRequest,
) (
	*npool.GetCoinDescriptionOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinDescriptionOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := validateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("GetCoinDescriptionOnly", "error", err)
		return &npool.GetCoinDescriptionOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetCoinDescriptionOnly", "error", err)
		return &npool.GetCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCoinDescriptions(
	ctx context.Context,
	in *npool.GetCoinDescriptionsRequest,
) (
	*npool.GetCoinDescriptionsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinDescriptions")
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
		logger.Sugar().Errorw("GetCoinDescriptions", "error", err)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetCoinDescriptions", "error", err)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCoinDescription(
	ctx context.Context,
	in *npool.ExistCoinDescriptionRequest,
) (
	*npool.ExistCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCoinDescription")
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
		return &npool.ExistCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check description: %v", err)
		return &npool.ExistCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinDescriptionResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCoinDescriptionConds(
	ctx context.Context,
	in *npool.ExistCoinDescriptionCondsRequest,
) (
	*npool.ExistCoinDescriptionCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCoinDescriptionConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := validateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("ExistCoinDescriptionConds", "error", err)
		return &npool.ExistCoinDescriptionCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistCoinDescriptionConds", "error", err)
		return &npool.ExistCoinDescriptionCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinDescriptionCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCoinDescriptions(
	ctx context.Context,
	in *npool.CountCoinDescriptionsRequest,
) (
	*npool.CountCoinDescriptionsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCoinDescriptions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := validateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("CountCoinDescriptions", "error", err)
		return &npool.CountCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountCoinDescriptions", "error", err)
		return &npool.CountCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCoinDescriptionsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCoinDescription(
	ctx context.Context,
	in *npool.DeleteCoinDescriptionRequest,
) (
	*npool.DeleteCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCoinDescription")
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
		return &npool.DeleteCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "description", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteCoinDescription", "error", err)
		return &npool.DeleteCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinDescriptionResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
