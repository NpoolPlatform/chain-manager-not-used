package currencyfeed

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/chain-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
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

var entity = ent.CurrencyFeed{
	ID:         uuid.New(),
	CoinTypeID: uuid.New(),
	FeedSource: uuid.NewString(),
	FeedType:   npool.FeedType_CoinBase.String(),
}

var (
	id         = entity.ID.String()
	coinTypeID = entity.CoinTypeID.String()
	feedSource = entity.FeedSource
	feedType   = npool.FeedType_CoinBase

	req = npool.CurrencyFeedReq{
		ID:         &id,
		CoinTypeID: &coinTypeID,
		FeedSource: &feedSource,
		FeedType:   &feedType,
	}
)

var info *ent.CurrencyFeed

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
	entities := []*ent.CurrencyFeed{
		{
			ID:         uuid.New(),
			CoinTypeID: uuid.New(),
			FeedSource: uuid.NewString(),
			FeedType:   npool.FeedType_CoinBase.String(),
		},
		{
			ID:         uuid.New(),
			CoinTypeID: uuid.New(),
			FeedSource: uuid.NewString(),
			FeedType:   npool.FeedType_CoinGecko.String(),
		},
	}

	reqs := []*npool.CurrencyFeedReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_coinTypeID := entity.CoinTypeID.String()
		_feedSource := entity.FeedSource
		_feedType := npool.FeedType(npool.FeedType_value[_entity.FeedType])

		reqs = append(reqs, &npool.CurrencyFeedReq{
			ID:         &_id,
			CoinTypeID: &_coinTypeID,
			FeedSource: &_feedSource,
			FeedType:   &_feedType,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	feedSource := uuid.NewString()
	req.FeedSource = &feedSource
	entity.FeedSource = feedSource

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
