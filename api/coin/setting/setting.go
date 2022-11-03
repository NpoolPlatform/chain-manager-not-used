//nolint:nolintlint,dupl
package setting

import (
	"context"

	converter "github.com/NpoolPlatform/chain-manager/pkg/converter/coin/setting"
	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/setting"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/setting"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	"github.com/google/uuid"
)

func (s *Server) CreateSetting(
	ctx context.Context,
	in *npool.CreateSettingRequest,
) (
	*npool.CreateSettingResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSetting")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "setting", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create setting: %v", err.Error())
		return &npool.CreateSettingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSettingResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSettings(
	ctx context.Context,
	in *npool.CreateSettingsRequest,
) (
	*npool.CreateSettingsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSettings")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateSettingsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	// TODO: verify infput

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "setting", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create settings: %v", err)
		return &npool.CreateSettingsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSettingsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSetting(ctx context.Context, in *npool.UpdateSettingRequest) (*npool.UpdateSettingResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSetting")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	// TODO: verify input

	span = commontracer.TraceInvoker(span, "setting", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create setting: %v", err.Error())
		return &npool.UpdateSettingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSettingResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSetting(ctx context.Context, in *npool.GetSettingRequest) (*npool.GetSettingResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSetting")
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
		return &npool.GetSettingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "setting", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get setting: %v", err)
		return &npool.GetSettingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSettingResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSettingOnly(
	ctx context.Context,
	in *npool.GetSettingOnlyRequest,
) (
	*npool.GetSettingOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSettingOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "setting", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get settings: %v", err)
		return &npool.GetSettingOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSettingOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSettings(ctx context.Context, in *npool.GetSettingsRequest) (*npool.GetSettingsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSettings")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "setting", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get settings: %v", err)
		return &npool.GetSettingsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSettingsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSetting(ctx context.Context, in *npool.ExistSettingRequest) (*npool.ExistSettingResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSetting")
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
		return &npool.ExistSettingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "setting", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check setting: %v", err)
		return &npool.ExistSettingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSettingResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSettingConds(ctx context.Context,
	in *npool.ExistSettingCondsRequest) (*npool.ExistSettingCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSettingConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "setting", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check setting: %v", err)
		return &npool.ExistSettingCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSettingCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSettings(ctx context.Context, in *npool.CountSettingsRequest) (*npool.CountSettingsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountSettings")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "setting", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count settings: %v", err)
		return &npool.CountSettingsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSettingsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSetting(ctx context.Context, in *npool.DeleteSettingRequest) (*npool.DeleteSettingResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteSetting")
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
		return &npool.DeleteSettingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "setting", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete setting: %v", err)
		return &npool.DeleteSettingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSettingResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
