package appcoin

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/appcoin"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"

	"github.com/google/uuid"
)

func CreateSet(c *ent.AppCoinCreate, in *npool.AppCoinReq) *ent.AppCoinCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Logo != nil {
		c.SetLogo(in.GetLogo())
	}
	if in.ForPay != nil {
		c.SetForPay(in.GetForPay())
	}
	if in.WithdrawAutoReviewAmount != nil {
		c.SetWithdrawAutoReviewAmount(decimal.RequireFromString(in.GetWithdrawAutoReviewAmount()))
	}
	if in.ProductPage != nil {
		c.SetProductPage(in.GetProductPage())
	}
	if in.Disabled != nil {
		c.SetDisabled(in.GetDisabled())
	}
	if in.DailyRewardAmount != nil {
		c.SetDailyRewardAmount(decimal.RequireFromString(in.GetDailyRewardAmount()))
	}
	return c
}

func Create(ctx context.Context, in *npool.AppCoinReq) (*ent.AppCoin, error) {
	var info *ent.AppCoin
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
		c := cli.AppCoin.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppCoinReq) ([]*ent.AppCoin, error) {
	var err error
	rows := []*ent.AppCoin{}

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
		bulk := make([]*ent.AppCoinCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.AppCoin.Create(), info)
		}
		rows, err = tx.AppCoin.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.AppCoin, in *npool.AppCoinReq) *ent.AppCoinUpdateOne {
	stm := info.Update()

	if in.Name != nil {
		stm = stm.SetName(in.GetName())
	}
	if in.Logo != nil {
		stm = stm.SetLogo(in.GetLogo())
	}
	if in.ForPay != nil {
		stm = stm.SetForPay(in.GetForPay())
	}
	if in.WithdrawAutoReviewAmount != nil {
		stm = stm.SetWithdrawAutoReviewAmount(decimal.RequireFromString(in.GetWithdrawAutoReviewAmount()))
	}
	if in.DeletedAt != nil {
		stm = stm.SetDeletedAt(in.GetDeletedAt())
	}
	if in.ProductPage != nil {
		stm = stm.SetProductPage(in.GetProductPage())
	}
	if in.Disabled != nil {
		stm = stm.SetDisabled(in.GetDisabled())
	}
	if in.DailyRewardAmount != nil {
		stm = stm.SetDailyRewardAmount(decimal.RequireFromString(in.GetDailyRewardAmount()))
	}

	return stm
}

func Update(ctx context.Context, in *npool.AppCoinReq) (*ent.AppCoin, error) {
	var info *ent.AppCoin
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
		info, err = tx.AppCoin.Query().Where(appcoin.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query appcoin: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update appcoin: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update appcoin: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppCoin, error) {
	var info *ent.AppCoin
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
		info, err = cli.AppCoin.Query().Where(appcoin.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppCoinQuery, error) {
	stm := cli.AppCoin.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appcoin.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appcoin.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(appcoin.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.ForPay != nil {
		switch conds.GetForPay().GetOp() {
		case cruder.EQ:
			stm.Where(appcoin.ForPay(conds.GetForPay().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.Disabled != nil {
		switch conds.GetDisabled().GetOp() {
		case cruder.EQ:
			stm.Where(appcoin.Disabled(conds.GetDisabled().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.IDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetIDs().GetOp() {
		case cruder.IN:
			stm.Where(appcoin.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetCoinTypeIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetCoinTypeIDs().GetOp() {
		case cruder.IN:
			stm.Where(appcoin.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppCoin, int, error) {
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

	rows := []*ent.AppCoin{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(appcoin.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppCoin, error) {
	var info *ent.AppCoin
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
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
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
		stm, err := SetQueryConds(conds, cli)
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
		exist, err = cli.AppCoin.Query().Where(appcoin.ID(id)).Exist(_ctx)
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
		stm, err := SetQueryConds(conds, cli)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppCoin, error) {
	var info *ent.AppCoin
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
		info, err = cli.AppCoin.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
