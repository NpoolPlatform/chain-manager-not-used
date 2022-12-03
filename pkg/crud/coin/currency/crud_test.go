package currency

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
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
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

var entity = ent.Currency{
	ID:              uuid.New(),
	CoinTypeID:      uuid.New(),
	FeedType:        npool.FeedType_CoinBase.String(),
	MarketValueHigh: decimal.RequireFromString("88.9123"),
	MarketValueLow:  decimal.RequireFromString("84.9123"),
}

var (
	id              = entity.ID.String()
	coinTypeID      = entity.CoinTypeID.String()
	feedType        = npool.FeedType(npool.FeedType_value[entity.FeedType])
	marketValueHigh = entity.MarketValueHigh.String()
	marketValueLow  = entity.MarketValueLow.String()

	req = npool.CurrencyReq{
		ID:              &id,
		CoinTypeID:      &coinTypeID,
		FeedType:        &feedType,
		MarketValueHigh: &marketValueHigh,
		MarketValueLow:  &marketValueLow,
	}
)

var info *ent.Currency

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
	entities := []*ent.Currency{
		{
			ID:              uuid.New(),
			CoinTypeID:      uuid.New(),
			FeedType:        npool.FeedType_CoinBase.String(),
			MarketValueHigh: decimal.RequireFromString("84.9123"),
			MarketValueLow:  decimal.RequireFromString("82.9123"),
		},
		{
			ID:              uuid.New(),
			CoinTypeID:      uuid.New(),
			FeedType:        npool.FeedType_CoinGecko.String(),
			MarketValueHigh: decimal.RequireFromString("88.4123"),
			MarketValueLow:  decimal.RequireFromString("84.2123"),
		},
	}

	reqs := []*npool.CurrencyReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_feedType := npool.FeedType(npool.FeedType_value[_entity.FeedType])
		_marketValueHigh := _entity.MarketValueHigh.String()
		_marketValueLow := _entity.MarketValueLow.String()

		reqs = append(reqs, &npool.CurrencyReq{
			ID:              &_id,
			CoinTypeID:      &_coinTypeID,
			FeedType:        &_feedType,
			MarketValueHigh: &_marketValueHigh,
			MarketValueLow:  &_marketValueLow,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	valueHigh := "11.22"
	valueLow := "11.12"

	req.MarketValueHigh = &valueHigh
	req.MarketValueLow = &valueLow

	entity.MarketValueHigh = decimal.RequireFromString(valueHigh)
	entity.MarketValueLow = decimal.RequireFromString(valueLow)

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
		entity.UpdatedAt = info.UpdatedAt
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
