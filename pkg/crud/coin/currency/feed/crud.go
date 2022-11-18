package currencyfeed

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyfeed"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"

	"github.com/google/uuid"
)

func CreateSet(c *ent.CurrencyFeedCreate, in *npool.CurrencyFeedReq) *ent.CurrencyFeedCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.FeedSource != nil {
		c.SetFeedSource(in.GetFeedSource())
	}
	if in.FeedType != nil {
		c.SetFeedType(in.GetFeedType().String())
	}
	return c
}

func Create(ctx context.Context, in *npool.CurrencyFeedReq) (*ent.CurrencyFeed, error) {
	var info *ent.CurrencyFeed
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.CurrencyFeed.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.CurrencyFeedReq) ([]*ent.CurrencyFeed, error) {
	var err error
	rows := []*ent.CurrencyFeed{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.CurrencyFeedCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.CurrencyFeed.Create(), info)
		}
		rows, err = tx.CurrencyFeed.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.CurrencyFeed, in *npool.CurrencyFeedReq) *ent.CurrencyFeedUpdateOne {
	stm := info.Update()

	if in.FeedSource != nil {
		stm = stm.SetFeedSource(in.GetFeedSource())
	}
	if in.Disabled != nil {
		stm = stm.SetDisabled(in.GetDisabled())
	}

	return stm
}

func Update(ctx context.Context, in *npool.CurrencyFeedReq) (*ent.CurrencyFeed, error) {
	var info *ent.CurrencyFeed
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.CurrencyFeed.Query().Where(currencyfeed.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query currencyfeed: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update currencyfeed: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update currencyfeed: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.CurrencyFeed, error) {
	var info *ent.CurrencyFeed
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
		info, err = cli.CurrencyFeed.Query().Where(currencyfeed.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.CurrencyFeedQuery, error) {
	stm := cli.CurrencyFeed.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(currencyfeed.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(currencyfeed.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.Disabled != nil {
		switch conds.GetDisabled().GetOp() {
		case cruder.EQ:
			stm.Where(currencyfeed.Disabled(conds.GetDisabled().GetValue()))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.CurrencyFeed, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.CurrencyFeed{}
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
			Order(ent.Desc(currencyfeed.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.CurrencyFeed, error) {
	var info *ent.CurrencyFeed
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

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
		exist, err = cli.CurrencyFeed.Query().Where(currencyfeed.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.CurrencyFeed, error) {
	var info *ent.CurrencyFeed
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
		info, err = cli.CurrencyFeed.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
