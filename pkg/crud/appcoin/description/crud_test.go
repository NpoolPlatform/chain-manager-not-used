package description

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
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
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

var entity = ent.CoinDescription{
	ID:         uuid.New(),
	AppID:      uuid.New(),
	CoinTypeID: uuid.New(),
	UsedFor:    npool.UsedFor_ProductPage.String(),
	Title:      uuid.NewString(),
	Message:    uuid.NewString(),
}

var (
	id         = entity.ID.String()
	appID      = entity.AppID.String()
	coinTypeID = entity.CoinTypeID.String()
	usedFor    = npool.UsedFor(npool.UsedFor_value[entity.UsedFor])
	title      = entity.Title
	message    = entity.Message

	req = npool.CoinDescriptionReq{
		ID:         &id,
		AppID:      &appID,
		CoinTypeID: &coinTypeID,
		UsedFor:    &usedFor,
		Title:      &title,
		Message:    &message,
	}
)

var info *ent.CoinDescription

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
	entities := []*ent.CoinDescription{
		{
			ID:         uuid.New(),
			AppID:      entity.AppID,
			CoinTypeID: uuid.New(),
			UsedFor:    npool.UsedFor_ProductPage.String(),
			Title:      uuid.NewString(),
			Message:    uuid.NewString(),
		},
		{
			ID:         uuid.New(),
			AppID:      entity.AppID,
			CoinTypeID: uuid.New(),
			UsedFor:    npool.UsedFor_ProductPage.String(),
			Title:      uuid.NewString(),
			Message:    uuid.NewString(),
		},
	}

	reqs := []*npool.CoinDescriptionReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_usedFor := npool.UsedFor(npool.UsedFor_value[entity.UsedFor])
		_title := _entity.Title
		_message := _entity.Message

		reqs = append(reqs, &npool.CoinDescriptionReq{
			ID:         &_id,
			AppID:      &_appID,
			CoinTypeID: &_coinTypeID,
			UsedFor:    &_usedFor,
			Title:      &_title,
			Message:    &_message,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	_title := uuid.NewString()
	_message := uuid.NewString()

	req.Title = &_title
	req.Message = &_message

	entity.Title = _title
	entity.Message = _message

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
