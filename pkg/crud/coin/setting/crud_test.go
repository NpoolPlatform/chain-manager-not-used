package setting

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
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"
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

var entity = ent.Setting{
	ID:                          uuid.New(),
	CoinTypeID:                  uuid.New(),
	FeeCoinTypeID:               uuid.New(),
	WithdrawFeeByStableUsd:      true,
	WithdrawFeeAmount:           decimal.RequireFromString("2.01"),
	CollectFeeAmount:            decimal.RequireFromString("2.01"),
	HotWalletFeeAmount:          decimal.RequireFromString("2.01"),
	LowFeeAmount:                decimal.RequireFromString("0.201"),
	HotWalletAccountAmount:      decimal.RequireFromString("0.201"),
	PaymentAccountCollectAmount: decimal.RequireFromString("0.201"),
	LeastTransferAmount:         decimal.RequireFromString("0.201"),
}

var (
	id                          = entity.ID.String()
	coinTypeID                  = entity.CoinTypeID.String()
	settingCoinTypeID           = entity.FeeCoinTypeID.String()
	withdrawFeeByStableUSD      = entity.WithdrawFeeByStableUsd
	withdrawFeeAmount           = entity.WithdrawFeeAmount.String()
	collectFeeAmount            = entity.CollectFeeAmount.String()
	hotWalletFeeAmount          = entity.HotWalletFeeAmount.String()
	lowFeeAmount                = entity.LowFeeAmount.String()
	hotWalletAccountAmount      = entity.HotWalletAccountAmount.String()
	paymentAccountCollectAmount = entity.PaymentAccountCollectAmount.String()
	leastTransferAmount         = entity.LeastTransferAmount.String()

	req = npool.SettingReq{
		ID:                          &id,
		CoinTypeID:                  &coinTypeID,
		FeeCoinTypeID:               &settingCoinTypeID,
		WithdrawFeeByStableUSD:      &withdrawFeeByStableUSD,
		WithdrawFeeAmount:           &withdrawFeeAmount,
		CollectFeeAmount:            &collectFeeAmount,
		HotWalletFeeAmount:          &hotWalletFeeAmount,
		LowFeeAmount:                &lowFeeAmount,
		HotWalletAccountAmount:      &hotWalletAccountAmount,
		PaymentAccountCollectAmount: &paymentAccountCollectAmount,
		LeastTransferAmount:         &leastTransferAmount,
	}
)

var info *ent.Setting

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
	entities := []*ent.Setting{
		{
			ID:                          uuid.New(),
			CoinTypeID:                  uuid.New(),
			FeeCoinTypeID:               uuid.New(),
			WithdrawFeeByStableUsd:      false,
			WithdrawFeeAmount:           decimal.RequireFromString("2.02"),
			CollectFeeAmount:            decimal.RequireFromString("2.02"),
			HotWalletFeeAmount:          decimal.RequireFromString("2.03"),
			LowFeeAmount:                decimal.RequireFromString("0.202"),
			HotWalletAccountAmount:      decimal.RequireFromString("0.201"),
			PaymentAccountCollectAmount: decimal.RequireFromString("0.201"),
			LeastTransferAmount:         decimal.RequireFromString("0.201"),
		},
		{
			ID:                          uuid.New(),
			CoinTypeID:                  uuid.New(),
			FeeCoinTypeID:               uuid.New(),
			WithdrawFeeByStableUsd:      true,
			WithdrawFeeAmount:           decimal.RequireFromString("2.03"),
			CollectFeeAmount:            decimal.RequireFromString("2.03"),
			HotWalletFeeAmount:          decimal.RequireFromString("2.02"),
			LowFeeAmount:                decimal.RequireFromString("0.201"),
			HotWalletAccountAmount:      decimal.RequireFromString("0.201"),
			PaymentAccountCollectAmount: decimal.RequireFromString("0.201"),
			LeastTransferAmount:         decimal.RequireFromString("0.201"),
		},
	}

	reqs := []*npool.SettingReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_settingCoinTypeID := _entity.FeeCoinTypeID.String()
		_withdrawFeeByStableUSD := _entity.WithdrawFeeByStableUsd
		_withdrawFeeAmount := _entity.WithdrawFeeAmount.String()
		_collectFeeAmount := _entity.CollectFeeAmount.String()
		_hotWalletFeeAmount := _entity.HotWalletFeeAmount.String()
		_lowFeeAmount := _entity.LowFeeAmount.String()
		_hotWalletAccountAmount := _entity.HotWalletAccountAmount.String()
		_paymentAccountCollectAmount := _entity.PaymentAccountCollectAmount.String()
		_leastTransferAmount := _entity.LeastTransferAmount.String()

		reqs = append(reqs, &npool.SettingReq{
			ID:                          &_id,
			CoinTypeID:                  &_coinTypeID,
			FeeCoinTypeID:               &_settingCoinTypeID,
			WithdrawFeeByStableUSD:      &_withdrawFeeByStableUSD,
			WithdrawFeeAmount:           &_withdrawFeeAmount,
			CollectFeeAmount:            &_collectFeeAmount,
			HotWalletFeeAmount:          &_hotWalletFeeAmount,
			LowFeeAmount:                &_lowFeeAmount,
			HotWalletAccountAmount:      &_hotWalletAccountAmount,
			PaymentAccountCollectAmount: &_paymentAccountCollectAmount,
			LeastTransferAmount:         &_leastTransferAmount,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	withdrawFeeByStableUSD := false
	withdrawFeeAmount := "22.01"
	collectFeeAmount := "23.01"
	hotWalletFeeAmount := "24.01"
	lowFeeAmount := "2.01"

	req.WithdrawFeeByStableUSD = &withdrawFeeByStableUSD
	req.WithdrawFeeAmount = &withdrawFeeAmount
	req.CollectFeeAmount = &collectFeeAmount
	req.HotWalletFeeAmount = &hotWalletFeeAmount
	req.LowFeeAmount = &lowFeeAmount

	entity.WithdrawFeeByStableUsd = withdrawFeeByStableUSD
	entity.WithdrawFeeAmount = decimal.RequireFromString(withdrawFeeAmount)
	entity.CollectFeeAmount = decimal.RequireFromString(collectFeeAmount)
	entity.HotWalletFeeAmount = decimal.RequireFromString(hotWalletFeeAmount)
	entity.LowFeeAmount = decimal.RequireFromString(lowFeeAmount)

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
