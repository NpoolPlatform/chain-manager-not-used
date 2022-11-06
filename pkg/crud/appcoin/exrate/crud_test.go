package exrate

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
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"
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

var entity = ent.ExchangeRate{
	ID:            uuid.New(),
	AppID:         uuid.New(),
	CoinTypeID:    uuid.New(),
	MarketValue:   decimal.RequireFromString("89.123"),
	SettlePercent: 80,
	Setter:        uuid.New(),
}

var (
	id            = entity.ID.String()
	appID         = entity.AppID.String()
	coinTypeID    = entity.CoinTypeID.String()
	marketValue   = entity.MarketValue.String()
	settlePercent = entity.SettlePercent
	setter        = entity.Setter.String()

	req = npool.ExchangeRateReq{
		ID:            &id,
		AppID:         &appID,
		CoinTypeID:    &coinTypeID,
		MarketValue:   &marketValue,
		SettlePercent: &settlePercent,
		Setter:        &setter,
	}
)

var info *ent.ExchangeRate

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		entity.SettleValue = info.MarketValue.
			Mul(decimal.NewFromInt(int64(info.SettlePercent))).
			Div(decimal.NewFromInt(100))
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.ExchangeRate{
		{
			ID:            uuid.New(),
			AppID:         entity.AppID,
			CoinTypeID:    uuid.New(),
			MarketValue:   decimal.RequireFromString("99.123"),
			SettlePercent: 80,
			Setter:        uuid.New(),
		},
		{
			ID:            uuid.New(),
			AppID:         entity.AppID,
			CoinTypeID:    uuid.New(),
			MarketValue:   decimal.RequireFromString("89.123"),
			SettlePercent: 80,
			Setter:        uuid.New(),
		},
	}

	reqs := []*npool.ExchangeRateReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_marketValue := _entity.MarketValue.String()
		_settlePercent := _entity.SettlePercent
		_setter := _entity.Setter.String()

		reqs = append(reqs, &npool.ExchangeRateReq{
			ID:            &_id,
			AppID:         &_appID,
			CoinTypeID:    &_coinTypeID,
			MarketValue:   &_marketValue,
			SettlePercent: &_settlePercent,
			Setter:        &_setter,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	/*
		marketValue := "1000.001"
		settlePercent := uint32(90)

		req.MarketValue = &marketValue
		req.SettlePercent = &settlePercent

		entity.MarketValue = decimal.RequireFromString(marketValue)
		entity.SettlePercent = settlePercent

		info, err := Update(context.Background(), &req)
		if assert.Nil(t, err) {
			entity.UpdatedAt = info.UpdatedAt
			entity.SettleValue = info.MarketValue.
				Mul(decimal.NewFromInt(int64(info.SettlePercent))).
				Div(decimal.NewFromInt(100))
			assert.Equal(t, info.String(), entity.String())
		}
	*/
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
