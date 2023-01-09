package appcoin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/chain-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appCoin = &npool.AppCoin{
	ID:                       uuid.NewString(),
	AppID:                    uuid.NewString(),
	CoinTypeID:               uuid.NewString(),
	Name:                     uuid.NewString(),
	Logo:                     uuid.NewString(),
	ForPay:                   false,
	WithdrawAutoReviewAmount: "89.123",
	DailyRewardAmount:        "89.122",
	Display:                  false,
	Disabled:                 false,
	MaxAmountPerWithdraw:     "89.123",
}

var appCoinReq = &npool.AppCoinReq{
	ID:                       &appCoin.ID,
	AppID:                    &appCoin.AppID,
	CoinTypeID:               &appCoin.CoinTypeID,
	Name:                     &appCoin.Name,
	Logo:                     &appCoin.Logo,
	ForPay:                   &appCoin.ForPay,
	WithdrawAutoReviewAmount: &appCoin.WithdrawAutoReviewAmount,
	DailyRewardAmount:        &appCoin.DailyRewardAmount,
	Display:                  &appCoin.Display,
	Disabled:                 &appCoin.Disabled,
	MaxAmountPerWithdraw:     &appCoin.MaxAmountPerWithdraw,
}

func createAppCoin(t *testing.T) {
	info, err := CreateAppCoin(context.Background(), appCoinReq)
	if assert.Nil(t, err) {
		appCoin.CreatedAt = info.CreatedAt
		appCoin.UpdatedAt = info.UpdatedAt
		assert.Equal(t, appCoin, info)
	}
}

func updateAppCoin(t *testing.T) {
	name := uuid.NewString()
	withdrawAutoReviewdAmount := "81.0234"
	forPay := true
	displayIndex := uint32(1)

	appCoinReq.Name = &name
	appCoinReq.WithdrawAutoReviewAmount = &withdrawAutoReviewdAmount
	appCoinReq.DailyRewardAmount = &withdrawAutoReviewdAmount
	appCoinReq.ForPay = &forPay
	appCoinReq.DisplayIndex = &displayIndex

	appCoin.Name = name
	appCoin.WithdrawAutoReviewAmount = withdrawAutoReviewdAmount
	appCoin.DailyRewardAmount = withdrawAutoReviewdAmount
	appCoin.ForPay = forPay
	appCoin.DisplayIndex = displayIndex

	info, err := UpdateAppCoin(context.Background(), appCoinReq)
	if assert.Nil(t, err) {
		appCoin.UpdatedAt = info.UpdatedAt
		assert.Equal(t, appCoin, info)
	}
}

func createAppCoins(t *testing.T) {
}

func getAppCoin(t *testing.T) {
	info, err := GetAppCoin(context.Background(), appCoin.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, appCoin, info)
	}
}

func deleteAppCoin(t *testing.T) {
	info, err := DeleteAppCoin(context.Background(), appCoin.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, appCoin, info)
	}

	_, err = GetAppCoin(context.Background(), info.ID)
	assert.NotNil(t, err)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAppCoin", createAppCoin)
	t.Run("updateAppCoin", updateAppCoin)
	t.Run("createAppCoins", createAppCoins)
	t.Run("getAppCoin", getAppCoin)
	t.Run("deleteAppCoin", deleteAppCoin)
}
