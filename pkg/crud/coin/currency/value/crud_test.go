package coinbase

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
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"
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

var entity = ent.CoinBase{
	ID:             uuid.New(),
	Name:           uuid.NewString(),
	Logo:           uuid.NewString(),
	Presale:        false,
	Unit:           "BTC",
	Env:            "main",
	ReservedAmount: decimal.RequireFromString("89.000"),
	ForPay:         false,
}

var (
	id             = entity.ID.String()
	name           = entity.Name
	logo           = entity.Logo
	presale        = entity.Presale
	unit           = entity.Unit
	env            = entity.Env
	reservedAmount = entity.ReservedAmount.String()
	forPay         = entity.ForPay

	req = npool.CoinBaseReq{
		ID:             &id,
		Name:           &name,
		Logo:           &logo,
		Presale:        &presale,
		Unit:           &unit,
		ENV:            &env,
		ReservedAmount: &reservedAmount,
		ForPay:         &forPay,
	}
)

var info *ent.CoinBase

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
	entities := []*ent.CoinBase{
		{
			ID:             uuid.New(),
			Name:           uuid.NewString(),
			Logo:           uuid.NewString(),
			Presale:        false,
			Unit:           "BTC",
			Env:            "main",
			ReservedAmount: decimal.RequireFromString("89.000"),
			ForPay:         false,
		},
		{
			ID:             uuid.New(),
			Name:           uuid.NewString(),
			Logo:           uuid.NewString(),
			Presale:        true,
			Unit:           "BTT",
			Env:            "main",
			ReservedAmount: decimal.RequireFromString("88.000"),
			ForPay:         false,
		},
	}

	reqs := []*npool.CoinBaseReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_name := _entity.Name
		_logo := _entity.Logo
		_presale := _entity.Presale
		_unit := _entity.Unit
		_env := _entity.Env
		_reservedAmount := _entity.ReservedAmount.String()
		_forPay := _entity.ForPay

		reqs = append(reqs, &npool.CoinBaseReq{
			ID:             &_id,
			Name:           &_name,
			Logo:           &_logo,
			Presale:        &_presale,
			Unit:           &_unit,
			ENV:            &_env,
			ReservedAmount: &_reservedAmount,
			ForPay:         &_forPay,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	presale := false
	reservedAmount := "81.0234"
	forPay := true

	req.Presale = &presale
	req.ReservedAmount = &reservedAmount
	req.ForPay = &forPay

	entity.Presale = presale
	entity.ReservedAmount = decimal.RequireFromString(reservedAmount)
	entity.ForPay = forPay

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
