package tx

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/tx"

	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"

	"github.com/google/uuid"
)

func CreateSet(c *ent.TranCreate, in *npool.TxReq) *ent.TranCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.FromAccountID != nil {
		c.SetFromAccountID(uuid.MustParse(in.GetFromAccountID()))
	}
	if in.ToAccountID != nil {
		c.SetToAccountID(uuid.MustParse(in.GetToAccountID()))
	}
	if in.Amount != nil {
		c.SetAmount(decimal.RequireFromString(in.GetAmount()))
	}
	if in.FeeAmount != nil {
		c.SetFeeAmount(decimal.RequireFromString(in.GetFeeAmount()))
	}
	if in.ChainTxID != nil {
		c.SetChainTxID(in.GetChainTxID())
	}
	c.SetState(npool.TxState_StateCreated.String())
	if in.Extra != nil {
		c.SetExtra(in.GetExtra())
	}
	if in.Type != nil {
		c.SetType(in.GetType().String())
	}
	return c
}

func Create(ctx context.Context, in *npool.TxReq) (*ent.Tran, error) {
	var info *ent.Tran
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
		c := cli.Tran.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.TxReq) ([]*ent.Tran, error) {
	var err error
	rows := []*ent.Tran{}

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
		bulk := make([]*ent.TranCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Tran.Create(), info)
		}
		rows, err = tx.Tran.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.TxReq) (*ent.Tran, error) {
	var info *ent.Tran
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
		info, err = tx.Tran.Query().Where(tran.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query tx: %v", err)
		}

		stm := info.Update()

		if in.State != nil {
			switch info.State {
			case npool.TxState_StateCreated.String():
				switch in.GetState() {
				case npool.TxState_StateWait:
				default:
					return fmt.Errorf("state is invalid")
				}
			case npool.TxState_StateWait.String():
				switch in.GetState() {
				case npool.TxState_StateTransferring:
				default:
					return fmt.Errorf("state is invalid")
				}
			case npool.TxState_StateTransferring.String():
				switch in.GetState() {
				case npool.TxState_StateSuccessful:
				case npool.TxState_StateFail:
				default:
					return fmt.Errorf("state is invalid")
				}
			case npool.TxState_StateSuccessful.String():
				fallthrough //nolint
			case npool.TxState_StateFail.String():
				fallthrough //nolint
			default:
				return fmt.Errorf("state is invalid")
			}
			stm = stm.SetState(in.GetState().String())
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update tx: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update tx: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Tran, error) {
	var info *ent.Tran
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
		info, err = cli.Tran.Query().Where(tran.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.TranQuery, error) {
	stm := cli.Tran.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(tran.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.AccountID != nil {
		switch conds.GetAccountID().GetOp() {
		case cruder.EQ:
			stm.Where(
				tran.Or(
					tran.FromAccountID(uuid.MustParse(conds.GetAccountID().GetValue())),
					tran.ToAccountID(uuid.MustParse(conds.GetAccountID().GetValue())),
				),
			)
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.AccountIDs != nil {
		switch conds.GetAccountIDs().GetOp() {
		case cruder.IN:
			ids := []uuid.UUID{}
			for _, acc := range conds.GetAccountIDs().GetValue() {
				ids = append(ids, uuid.MustParse(acc))
			}

			stm.Where(
				tran.Or(
					tran.FromAccountIDIn(ids...),
					tran.ToAccountIDIn(ids...),
				),
			)
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(tran.State(npool.TxState(conds.GetState().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.Type != nil {
		switch conds.GetType().GetOp() {
		case cruder.EQ:
			stm.Where(tran.Type(npool.TxType(conds.GetType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Tran, int, error) {
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

	rows := []*ent.Tran{}
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
			Order(ent.Desc(tran.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Tran, error) {
	var info *ent.Tran
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
		exist, err = cli.Tran.Query().Where(tran.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Tran, error) {
	var info *ent.Tran
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
		info, err = cli.Tran.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
