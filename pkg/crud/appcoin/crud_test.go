package appcoin

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
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"
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

var entity = ent.AppCoin{
	ID:                       uuid.New(),
	AppID:                    uuid.New(),
	CoinTypeID:               uuid.New(),
	Name:                     uuid.NewString(),
	Logo:                     uuid.NewString(),
	ForPay:                   false,
	WithdrawAutoReviewAmount: decimal.RequireFromString("89.000"),
	DailyRewardAmount:        decimal.RequireFromString("89.001"),
	Disabled:                 false,
	Display:                  false,
	MaxAmountPerWithdraw:     decimal.RequireFromString("89.000"),
}

var (
	id                        = entity.ID.String()
	appID                     = entity.AppID.String()
	coinTypeID                = entity.CoinTypeID.String()
	name                      = entity.Name
	logo                      = entity.Logo
	forPay                    = entity.ForPay
	withdrawAutoReviewdAmount = entity.WithdrawAutoReviewAmount.String()
	dailyRewardAmount         = entity.DailyRewardAmount.String()
	maxAmountPerWithdraw      = entity.MaxAmountPerWithdraw.String()

	req = npool.AppCoinReq{
		ID:                       &id,
		AppID:                    &appID,
		CoinTypeID:               &coinTypeID,
		Name:                     &name,
		Logo:                     &logo,
		ForPay:                   &forPay,
		WithdrawAutoReviewAmount: &withdrawAutoReviewdAmount,
		DailyRewardAmount:        &dailyRewardAmount,
		Disabled:                 &entity.Disabled,
		Display:                  &entity.Display,
		MaxAmountPerWithdraw:     &maxAmountPerWithdraw,
	}
)

var info *ent.AppCoin

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
	entities := []*ent.AppCoin{
		{
			ID:                       uuid.New(),
			AppID:                    entity.AppID,
			CoinTypeID:               uuid.New(),
			Name:                     uuid.NewString(),
			Logo:                     uuid.NewString(),
			ForPay:                   false,
			WithdrawAutoReviewAmount: decimal.RequireFromString("89.000"),
			DailyRewardAmount:        decimal.RequireFromString("89.001"),
			Disabled:                 false,
			Display:                  false,
		},
		{
			ID:                       uuid.New(),
			AppID:                    entity.AppID,
			CoinTypeID:               uuid.New(),
			Name:                     uuid.NewString(),
			Logo:                     uuid.NewString(),
			ForPay:                   true,
			WithdrawAutoReviewAmount: decimal.RequireFromString("90.000"),
			DailyRewardAmount:        decimal.RequireFromString("90.001"),
			Disabled:                 false,
			Display:                  false,
		},
	}

	reqs := []*npool.AppCoinReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_name := _entity.Name
		_logo := _entity.Logo
		_forPay := _entity.ForPay
		_withdrawAutoReviewdAmount := _entity.WithdrawAutoReviewAmount.String()
		_dailyRewardAmount := _entity.DailyRewardAmount.String()

		reqs = append(reqs, &npool.AppCoinReq{
			ID:                       &_id,
			AppID:                    &_appID,
			CoinTypeID:               &_coinTypeID,
			Name:                     &_name,
			Logo:                     &_logo,
			ForPay:                   &_forPay,
			WithdrawAutoReviewAmount: &_withdrawAutoReviewdAmount,
			DailyRewardAmount:        &_dailyRewardAmount,
			Disabled:                 &_entity.Disabled,
			Display:                  &_entity.Display,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	name := uuid.NewString()
	withdrawAutoReviewdAmount := "81.0234"
	forPay := true

	req.Name = &name
	req.WithdrawAutoReviewAmount = &withdrawAutoReviewdAmount
	req.DailyRewardAmount = &dailyRewardAmount
	req.ForPay = &forPay

	entity.Name = name
	entity.WithdrawAutoReviewAmount = decimal.RequireFromString(withdrawAutoReviewdAmount)
	entity.DailyRewardAmount = decimal.RequireFromString(dailyRewardAmount)
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
