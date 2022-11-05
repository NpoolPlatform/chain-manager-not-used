package setting

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/setting"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	"github.com/google/uuid"
)

func CreateSet(c *ent.SettingCreate, in *npool.SettingReq) *ent.SettingCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.FeeCoinTypeID != nil {
		c.SetFeeCoinTypeID(uuid.MustParse(in.GetFeeCoinTypeID()))
	}
	if in.WithdrawFeeByStableUSD != nil {
		c.SetWithdrawFeeByStableUsd(in.GetWithdrawFeeByStableUSD())
	}
	if in.WithdrawFeeAmount != nil {
		c.SetWithdrawFeeAmount(decimal.RequireFromString(in.GetWithdrawFeeAmount()))
	}
	if in.CollectFeeAmount != nil {
		c.SetCollectFeeAmount(decimal.RequireFromString(in.GetCollectFeeAmount()))
	}
	if in.HotWalletFeeAmount != nil {
		c.SetHotWalletFeeAmount(decimal.RequireFromString(in.GetHotWalletFeeAmount()))
	}
	if in.LowFeeAmount != nil {
		c.SetLowFeeAmount(decimal.RequireFromString(in.GetLowFeeAmount()))
	}
	if in.HotWalletAccountAmount != nil {
		c.SetHotWalletAccountAmount(decimal.RequireFromString(in.GetHotWalletAccountAmount()))
	}
	if in.PaymentAccountCollectAmount != nil {
		c.SetPaymentAccountCollectAmount(decimal.RequireFromString(in.GetPaymentAccountCollectAmount()))
	}
	return c
}

func Create(ctx context.Context, in *npool.SettingReq) (*ent.Setting, error) {
	var info *ent.Setting
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.Setting.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.SettingReq) ([]*ent.Setting, error) {
	var err error
	rows := []*ent.Setting{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.SettingCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Setting.Create(), info)
		}
		rows, err = tx.Setting.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.SettingReq) (*ent.Setting, error) {
	var info *ent.Setting
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Setting.Query().Where(setting.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query setting: %v", err)
		}

		stm := info.Update()

		if in.WithdrawFeeByStableUSD != nil {
			stm.SetWithdrawFeeByStableUsd(in.GetWithdrawFeeByStableUSD())
		}
		if in.WithdrawFeeAmount != nil {
			stm.SetWithdrawFeeAmount(decimal.RequireFromString(in.GetWithdrawFeeAmount()))
		}
		if in.CollectFeeAmount != nil {
			stm.SetCollectFeeAmount(decimal.RequireFromString(in.GetCollectFeeAmount()))
		}
		if in.HotWalletFeeAmount != nil {
			stm.SetHotWalletFeeAmount(decimal.RequireFromString(in.GetHotWalletFeeAmount()))
		}
		if in.LowFeeAmount != nil {
			stm.SetLowFeeAmount(decimal.RequireFromString(in.GetLowFeeAmount()))
		}
		if in.HotWalletAccountAmount != nil {
			stm.SetHotWalletAccountAmount(decimal.RequireFromString(in.GetHotWalletAccountAmount()))
		}
		if in.PaymentAccountCollectAmount != nil {
			stm.SetPaymentAccountCollectAmount(decimal.RequireFromString(in.GetPaymentAccountCollectAmount()))
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update setting: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update setting: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Setting, error) {
	var info *ent.Setting
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Setting.Query().Where(setting.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.SettingQuery, error) {
	stm := cli.Setting.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(setting.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(setting.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	if conds.FeeCoinTypeID != nil {
		switch conds.GetFeeCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(setting.FeeCoinTypeID(uuid.MustParse(conds.GetFeeCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Setting, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Setting{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(setting.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Setting, error) {
	var info *ent.Setting
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Setting.Query().Where(setting.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Setting, error) {
	var info *ent.Setting
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Setting.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
