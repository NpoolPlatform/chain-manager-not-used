package tx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/chain-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entity = ent.Tran{
	ID:            uuid.New(),
	FromAccountID: uuid.New(),
	ToAccountID:   uuid.New(),
	Amount:        decimal.RequireFromString("89.000"),
	FeeAmount:     decimal.RequireFromString("2.000"),
	ChainTxID:     uuid.NewString(),
	State:         npool.TxState_StateCreated.String(),
	Extra:         uuid.NewString(),
}

var (
	id            = entity.ID.String()
	fromAccountID = entity.FromAccountID.String()
	toAccountID   = entity.ToAccountID.String()
	amount        = entity.Amount.String()
	feeAmount     = entity.FeeAmount.String()
	chainTxID     = entity.ChainTxID
	state         = npool.TxState(npool.TxState_value[entity.State])
	extra         = entity.Extra

	req = npool.TxReq{
		ID:            &id,
		FromAccountID: &fromAccountID,
		ToAccountID:   &toAccountID,
		Amount:        &amount,
		FeeAmount:     &feeAmount,
		ChainTxID:     &chainTxID,
		State:         &state,
		Extra:         &extra,
	}
)

var info *ent.Tran

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Tran{
		{
			ID:            uuid.New(),
			FromAccountID: uuid.New(),
			ToAccountID:   uuid.New(),
			Amount:        decimal.RequireFromString("99.000"),
			FeeAmount:     decimal.RequireFromString("2.000"),
			ChainTxID:     uuid.NewString(),
			State:         npool.TxState_StateTransferring.String(),
			Extra:         uuid.NewString(),
		},
		{
			ID:            uuid.New(),
			FromAccountID: uuid.New(),
			ToAccountID:   uuid.New(),
			Amount:        decimal.RequireFromString("109.000"),
			FeeAmount:     decimal.RequireFromString("2.000"),
			ChainTxID:     uuid.NewString(),
			State:         npool.TxState_StateWait.String(),
			Extra:         uuid.NewString(),
		},
	}

	reqs := []*npool.TxReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_fromAccountID := _entity.FromAccountID.String()
		_toAccountID := _entity.ToAccountID.String()
		_amount := _entity.Amount.String()
		_feeAmount := _entity.FeeAmount.String()
		_chainTxID := _entity.ChainTxID
		_state := npool.TxState(npool.TxState_value[_entity.State])
		_extra := _entity.Extra

		reqs = append(reqs, &npool.TxReq{
			ID:            &_id,
			FromAccountID: &_fromAccountID,
			ToAccountID:   &_toAccountID,
			Amount:        &_amount,
			FeeAmount:     &_feeAmount,
			ChainTxID:     &_chainTxID,
			State:         &_state,
			Extra:         &_extra,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	state := npool.TxState_StateSuccessful

	req.State = &state

	entity.State = state.String()

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0].String(), entity.String())
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("add", add)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
