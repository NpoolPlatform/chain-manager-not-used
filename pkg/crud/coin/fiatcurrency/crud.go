package fiatcurrency

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"

	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrency"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	"github.com/google/uuid"
)

func CreateSet(c *ent.FiatCurrencyCreate, in *npool.FiatCurrencyReq) *ent.FiatCurrencyCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.FiatCurrencyTypeID != nil {
		c.SetFiatCurrencyTypeID(uuid.MustParse(in.GetFiatCurrencyTypeID()))
	}
	if in.FeedType != nil {
		c.SetFeedType(in.GetFeedType().String())
	}
	if in.MarketValueHigh != nil {
		c.SetMarketValueHigh(decimal.RequireFromString(in.GetMarketValueHigh()))
	}
	if in.MarketValueLow != nil {
		c.SetMarketValueLow(decimal.RequireFromString(in.GetMarketValueLow()))
	}
	return c
}

func Create(ctx context.Context, in *npool.FiatCurrencyReq) (*ent.FiatCurrency, error) {
	var info *ent.FiatCurrency
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
		c := cli.FiatCurrency.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.FiatCurrencyReq) ([]*ent.FiatCurrency, error) {
	var err error
	rows := []*ent.FiatCurrency{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.FiatCurrencyCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.FiatCurrency.Create(), info)
		}
		rows, err = tx.FiatCurrency.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.FiatCurrency, in *npool.FiatCurrencyReq) *ent.FiatCurrencyUpdateOne {
	stm := info.Update()

	if in.MarketValueHigh != nil {
		stm = stm.SetMarketValueHigh(decimal.RequireFromString(in.GetMarketValueHigh()))
	}
	if in.MarketValueLow != nil {
		stm = stm.SetMarketValueLow(decimal.RequireFromString(in.GetMarketValueLow()))
	}

	return stm
}

func Update(ctx context.Context, in *npool.FiatCurrencyReq) (*ent.FiatCurrency, error) {
	var info *ent.FiatCurrency
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
		info, err = tx.FiatCurrency.Query().Where(fiatcurrency.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query fiatcurrency: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update fiatcurrency: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update fiatcurrency: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.FiatCurrency, error) {
	var info *ent.FiatCurrency
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
		info, err = cli.FiatCurrency.Query().Where(fiatcurrency.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.FiatCurrencyQuery, error) {
	stm := cli.FiatCurrency.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(fiatcurrency.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.FiatCurrencyTypeID != nil {
		switch conds.GetFiatCurrencyTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(fiatcurrency.FiatCurrencyTypeID(uuid.MustParse(conds.GetFiatCurrencyTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.StartAt != nil {
		switch conds.GetStartAt().GetOp() {
		case cruder.GTE:
			stm.Where(fiatcurrency.CreatedAtGTE(conds.GetStartAt().GetValue()))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.EndAt != nil {
		switch conds.GetEndAt().GetOp() {
		case cruder.GTE:
			stm.Where(fiatcurrency.CreatedAtGTE(conds.GetEndAt().GetValue()))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.FiatCurrency, int, error) {
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

	rows := []*ent.FiatCurrency{}
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
			Order(ent.Desc(fiatcurrency.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.FiatCurrency, error) {
	var info *ent.FiatCurrency
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
		exist, err = cli.FiatCurrency.Query().Where(fiatcurrency.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.FiatCurrency, error) {
	var info *ent.FiatCurrency
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
		info, err = cli.FiatCurrency.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
