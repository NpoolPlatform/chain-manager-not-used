//nolint:nolintlint,dupl
package currencyfeed

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/currency/feed"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/feed"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"

	"github.com/google/uuid"
)

func (s *Server) CreateCurrencyFeed(ctx context.Context, in *npool.CreateCurrencyFeedRequest) (*npool.CreateCurrencyFeedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCurrencyFeed")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currencyfeed: %v", err.Error())
		return &npool.CreateCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyFeedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCurrencyFeeds(ctx context.Context, in *npool.CreateCurrencyFeedsRequest) (*npool.CreateCurrencyFeedsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCurrencyFeeds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCurrencyFeedsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		return &npool.CreateCurrencyFeedsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create currencyfeeds: %v", err)
		return &npool.CreateCurrencyFeedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyFeedsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCurrencyFeed(ctx context.Context, in *npool.UpdateCurrencyFeedRequest) (*npool.UpdateCurrencyFeedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCurrencyFeed")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCurrencyFeed", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().FeedSource != nil && in.GetInfo().GetFeedSource() == "" {
		logger.Sugar().Errorw("UpdateCurrencyFeed", "FeedSource", in.GetInfo().GetFeedSource())
		return &npool.UpdateCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currencyfeed: %v", err.Error())
		return &npool.UpdateCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCurrencyFeedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyFeed(ctx context.Context, in *npool.GetCurrencyFeedRequest) (*npool.GetCurrencyFeedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyFeed")
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
		return &npool.GetCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get currencyfeed: %v", err)
		return &npool.GetCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyFeedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyFeedOnly(ctx context.Context, in *npool.GetCurrencyFeedOnlyRequest) (*npool.GetCurrencyFeedOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyFeedOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCurrencyFeedOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get currencyfeeds: %v", err)
		return &npool.GetCurrencyFeedOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyFeedOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCurrencyFeeds(ctx context.Context, in *npool.GetCurrencyFeedsRequest) (*npool.GetCurrencyFeedsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCurrencyFeeds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetCurrencyFeedsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get currencyfeeds: %v", err)
		return &npool.GetCurrencyFeedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyFeedsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCurrencyFeed(ctx context.Context, in *npool.ExistCurrencyFeedRequest) (*npool.ExistCurrencyFeedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCurrencyFeed")
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
		return &npool.ExistCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check currencyfeed: %v", err)
		return &npool.ExistCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCurrencyFeedResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCurrencyFeedConds(ctx context.Context,
	in *npool.ExistCurrencyFeedCondsRequest) (*npool.ExistCurrencyFeedCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCurrencyFeedConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistCurrencyFeedCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check currencyfeed: %v", err)
		return &npool.ExistCurrencyFeedCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCurrencyFeedCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCurrencyFeeds(ctx context.Context, in *npool.CountCurrencyFeedsRequest) (*npool.CountCurrencyFeedsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCurrencyFeeds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountCurrencyFeedsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count currencyfeeds: %v", err)
		return &npool.CountCurrencyFeedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCurrencyFeedsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCurrencyFeed(ctx context.Context, in *npool.DeleteCurrencyFeedRequest) (*npool.DeleteCurrencyFeedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCurrencyFeed")
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
		return &npool.DeleteCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "currencyfeed", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete currencyfeed: %v", err)
		return &npool.DeleteCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCurrencyFeedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
